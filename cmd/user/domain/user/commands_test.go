package user

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
)

func TestUnmarshalChangeEmailAddress(t *testing.T) {
	testJSON := []byte(`{"id":"4dded431-acee-4078-86c6-9dffa5efba1e","email":"test@test.com"}`)

	testUnmarshalCommand(t, testJSON, &ChangeEmailAddress{})
}

func TestUnmarshalRegisterWithEmail(t *testing.T) {
	testJSON := []byte(`{"email":"test@test.com"}`)

	testUnmarshalCommand(t, testJSON, &RegisterWithEmail{})
}

func TestUnmarshalRegisterWithFacebook(t *testing.T) {
	testJSON := []byte(`{"email":"test@test.com","facebookId":""}`)

	testUnmarshalCommand(t, testJSON, &RegisterWithFacebook{})
}

func TestUnmarshalRegisterWithGoogle(t *testing.T) {
	testJSON := []byte(`{"email":"test@test.com","googleId":""}`)

	testUnmarshalCommand(t, testJSON, &RegisterWithGoogle{})
}

func testUnmarshalCommand(t *testing.T, testJSON []byte, c interface{}) {
	err := json.Unmarshal(testJSON, c)
	if err != nil {
		t.Fatal(err)
	}

	j, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	cmp := bytes.Compare(j, testJSON)
	if cmp != 0 {
		t.Errorf("Serialize command did not match expected result: %s | %d", string(j), cmp)
	}
}

// func TestOnChangeEmailAddress(t *testing.T) {
// 	handler := OnChangeEmailAddress(new(repositoryMock))
// 	cmd := &ChangeEmailAddress{}

// 	testCommandHandler(t, func(ctx context.Context, out chan<- error) {
// 		if f, ok := handler.(func(context.Context, *ChangeEmailAddress, chan<- error)); ok {
// 			f(ctx, cmd, out)
// 		} else {
// 			t.Fatal("Could not call handler")
// 		}
// 	})
// }

// func TestOnRegisterWithEmail(t *testing.T) {
// 	handler := OnRegisterWithEmail(new(repositoryMock))
// 	cmd := &RegisterWithEmail{}

// 	testCommandHandler(t, func(ctx context.Context, out chan<- error) {
// 		if f, ok := handler.(func(context.Context, *RegisterWithEmail, chan<- error)); ok {
// 			f(ctx, cmd, out)
// 		} else {
// 			t.Fatal("Could not call handler")
// 		}
// 	})
// }

// func TestOnRegisterWithFacebook(t *testing.T) {
// 	handler := OnRegisterWithFacebook(new(repositoryMock))
// 	cmd := &RegisterWithFacebook{}

// 	testCommandHandler(t, func(ctx context.Context, out chan<- error) {
// 		if f, ok := handler.(func(context.Context, *RegisterWithFacebook, chan<- error)); ok {
// 			f(ctx, cmd, out)
// 		} else {
// 			t.Fatal("Could not call handler")
// 		}
// 	})
// }

// func TestOnRegisterWithGoogle(t *testing.T) {
// 	handler := OnRegisterWithGoogle(new(repositoryMock))
// 	cmd := &RegisterWithGoogle{}

// 	testCommandHandler(t, func(ctx context.Context, out chan<- error) {
// 		if f, ok := handler.(func(context.Context, *RegisterWithGoogle, chan<- error)); ok {
// 			f(ctx, cmd, out)
// 		} else {
// 			t.Fatal("Could not call handler")
// 		}
// 	})
// }

type repositoryMock int

func (r *repositoryMock) Save(ctx context.Context, u *User) error {
	return nil
}

func (r *repositoryMock) Get(id uuid.UUID) *User {
	return New()
}

func testCommandHandler(t *testing.T, fn func(context.Context, chan<- error)) {
	ctx := context.Background()
	out := make(chan error)
	defer close(out)

	go fn(ctx, out)

	select {
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	case err := <-out:
		if err != nil {
			t.Error(err)
		}
	}
}
