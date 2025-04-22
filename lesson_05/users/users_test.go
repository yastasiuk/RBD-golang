package users

import (
	"errors"
	"fmt"
	store "lesson_05/document_store"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	t.Run("Should create user", func(t *testing.T) {
		serv := InitService()

		user, err := serv.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`)
		if err != nil {
			t.Error(fmt.Errorf("should not return any errors on collection creation, got: %w", err))
		}

		if user.ID != "unique-id" {
			t.Error(fmt.Errorf("should return unique-id on user, got: %s", user.ID))
		}

		if user.Name != "Jon Doe" {
			t.Error(fmt.Errorf("should return Jon Doe, got: %s", user.Name))
		}
	})

	t.Run("Should create user with only ID", func(t *testing.T) {
		serv := InitService()

		user, err := serv.CreateUser(`{"id": "unique-id"}`)
		if err != nil {
			t.Error(fmt.Errorf("should not return any errors on collection creation, got: %w", err))
		}

		if user.ID != "unique-id" {
			t.Error(fmt.Errorf("should return unique-id on user, got: %s", user.ID))
		}

		if user.Name != "" {
			t.Error(fmt.Errorf("should return empty field, got: %s", user.Name))
		}
	})

	t.Run("Should not create duplicate user", func(t *testing.T) {
		serv := InitService()

		if _, err := serv.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`); err != nil {
			t.Error(fmt.Errorf("should not return any errors on collection creation, got: %w", err))
		}

		if _, err := serv.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`); !errors.Is(err, store.ErrUserUniquenessValidation) {
			t.Error(fmt.Errorf("should return error on duplicate user creation: %w", err))
		}
	})

	t.Run("Should not create user without required fields", func(t *testing.T) {
		serv := InitService()

		if _, err := serv.CreateUser(`{"name": "Jon Doe"}`); !errors.Is(err, store.ErrValidationFailed) {
			t.Error("Should validation error when required field is missing")
		}
	})
}

func TestUser_ListUser(t *testing.T) {
	t.Run("Should return list of users", func(t *testing.T) {
		serv := InitService()

		if _, err := serv.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`); err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}

		if _, err := serv.CreateUser(`{"name": "Santa", "id": "santa-id"}`); err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}

		users, err := serv.ListUsers()
		if err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}
		if len(users) != 2 {
			t.Error(fmt.Errorf("should return 2 users, got: %v", len(users)))
		}
	})
}

func TestUser_GetUser(t *testing.T) {
	t.Run("Should return user", func(t *testing.T) {
		serv := InitService()

		if _, err := serv.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`); err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}

		user, err := serv.GetUser("unique-id")
		if err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}

		if user.Name != "Jon Doe" {
			t.Error(fmt.Errorf("should return Jon Doe, got: %s", user.Name))
		}

		if user.ID != "unique-id" {
			t.Error(fmt.Errorf("should return correct ID, got: %s", user.ID))
		}
	})

	t.Run("Should return error if user not found", func(t *testing.T) {
		serv := InitService()

		user, err := serv.GetUser("not-exists")
		if !errors.Is(err, ErrUserNotFound) {
			t.Error(fmt.Errorf("error has incorrect type: %w", err))
		}

		if user != nil {
			t.Error(fmt.Errorf("should return nil user, got: %v", user))
		}
	})
}

func TestUser_DeleteUser(t *testing.T) {
	t.Run("Should delete user", func(t *testing.T) {
		serv := InitService()

		if _, err := serv.CreateUser(`{"name": "Jon Doe", "id": "unique-id"}`); err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}

		err := serv.DeleteUser("unique-id")
		if err != nil {
			t.Error(fmt.Errorf("should not fail: %w", err))
		}
	})

	t.Run("Should return error if user not found", func(t *testing.T) {
		serv := InitService()

		err := serv.DeleteUser("not-exists")
		if !errors.Is(err, ErrUserNotFound) {
			t.Error(fmt.Errorf("incorrect error type: %w", err))
		}
	})
}
