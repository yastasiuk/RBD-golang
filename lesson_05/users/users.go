package users

import (
	"encoding/json"
	"errors"
	"fmt"
	store "lesson_05/document_store"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID   string `json:"id" document:"_id"`
	Name string `json:"name" document:"firstName"`
}

type Service struct {
	coll *store.Collection
}

func InitService() *Service {
	cfg := store.CollectionConfig{PrimaryKey: "_id"}

	service := Service{
		coll: store.NewCollection(&cfg),
	}
	return &service
}

func (s *Service) CreateUser(body string) (*User, error) {
	user := &User{}

	if err := json.Unmarshal([]byte(body), user); err != nil {
		return nil, err
	}

	userDoc, err := MarshalDocument(user)
	if err != nil {
		return nil, errors.Join(errors.New("failed to convert User to Document"), err)
	}

	if _, creationErr := s.coll.Put(*userDoc); creationErr != nil {
		return nil, fmt.Errorf("failed to create new user: %w", creationErr)
	}

	return user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	var errs []error
	docs := s.coll.List()
	users := make([]User, 0, len(docs))

	for _, doc := range docs {
		user := User{}
		err := UnmarshalDocument(&doc, &user)
		if err != nil {
			errs = append(errs, err)
		} else {
			users = append(users, user)
		}
	}

	return users, errors.Join(errs...)
}

func (s *Service) GetUser(userID string) (*User, error) {
	doc, err := s.coll.Get(userID)
	if err != nil {
		if errors.Is(err, store.ErrDocumentNotFound) {
			return nil, fmt.Errorf("%w. id: %s", ErrUserNotFound, userID)
		}
		return nil, err
	}

	user := User{}
	if err := UnmarshalDocument(doc, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) DeleteUser(userID string) error {
	if ok := s.coll.Delete(userID); !ok {
		return ErrUserNotFound
	}

	return nil
}
