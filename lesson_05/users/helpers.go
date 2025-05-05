package users

import (
	"errors"
	"fmt"
	. "lesson_05/document_store"
	"reflect"
)

var ErrorInvalidInputType = errors.New("invalid input type")
var ErrorInvalidOutputType = errors.New("invalid output type")
var ErrorUnmarshalError = errors.New("unmarshal error")

func MarshalDocument(input any) (*Document, error) {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if k := t.Kind(); k != reflect.Struct {
		return nil, fmt.Errorf("%w. Passed: %v", ErrorInvalidInputType, k)
	}

	documentFields := map[string]DocumentField{}
	for i := 0; i < t.NumField(); i++ {
		if docKey := t.Field(i).Tag.Get("document"); docKey != "" {
			value := v.Field(i)
			documentFields[docKey] = DocumentField{
				Type:  GetType(value.Kind()),
				Value: value.Interface(),
			}
		}
	}

	doc := Document{
		Fields: documentFields,
	}

	return &doc, nil
}

func UnmarshalDocument(doc *Document, output any) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("recovered from panic: %v", e)
		}
	}()

	t := reflect.TypeOf(output)
	v := reflect.ValueOf(output)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if k := t.Kind(); k != reflect.Struct {
		return fmt.Errorf("%w. Passed: %v", ErrorInvalidOutputType, k)
	}

	var errs []error
	for i := 0; i < t.NumField(); i++ {
		if docKey := t.Field(i).Tag.Get("document"); docKey != "" {
			if !v.FieldByName(t.Field(i).Name).CanSet() {
				errs = append(errs, fmt.Errorf("%w: field %v not settable", ErrorUnmarshalError, t.Field(i).Name))
			} else if reflect.TypeOf(doc.GetField(docKey)).Kind() != v.FieldByName(t.Field(i).Name).Kind() {
				errs = append(errs, fmt.Errorf("%w: doc %s key and and output %s have different types", ErrorUnmarshalError, docKey, t.Field(i).Name))
			}
		}
	}

	for i := 0; i < t.NumField(); i++ {
		if docKey := t.Field(i).Tag.Get("document"); docKey != "" {
			if v.FieldByName(t.Field(i).Name).CanSet() {
				if reflect.TypeOf(doc.GetField(docKey)).Kind() == v.FieldByName(t.Field(i).Name).Kind() {
					v.FieldByName(t.Field(i).Name).Set(reflect.ValueOf(doc.GetField(docKey)))
				}
			}
		}
	}

	if len(errs) > 0 {
		err = errors.Join(errs...)
	}

	return err
}
