package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/forgoty/go-todo/pkg/web"
	"github.com/stretchr/testify/assert"
)

const (
	validUserSignIn    = `{"username":"Jon Snow","password":"jonshow1"}`
	noUsernameField    = `{"user":"Jon Snow","password":"jonshow1"}`
	noPasswordField    = `{"username":"Jon Snow","pass":"jonshow1"}`
	noRequiredFields   = `{"user":"Jon Snow","pass":"jonshow1"}`
	emptyPayload       = `{}`
	usernameIsTooShort = `{"username":"Jon","password":"jonshow1"}`
	passwordIsTooShort = `{"username":"Jon Show","password":"jon"}`
)

type userControllerTest struct {
	input              string
	expectedStatusCode int
	happy              bool
}

func preparePostContext(endpoint, payload string, rec *httptest.ResponseRecorder) web.Context {
	e := web.New()
	req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	return e.NewContext(req, rec)
}

func TestUserSignUp(t *testing.T) {
	// Setup
	c, _ := provideUserController("123", "123", 12*time.Hour)
	tests := []userControllerTest{
		userControllerTest{validUserSignIn, http.StatusCreated, true},
		userControllerTest{noUsernameField, http.StatusBadRequest, false},
		userControllerTest{noPasswordField, http.StatusBadRequest, false},
		userControllerTest{noRequiredFields, http.StatusBadRequest, false},
		userControllerTest{emptyPayload, http.StatusBadRequest, false},
		userControllerTest{usernameIsTooShort, http.StatusBadRequest, false},
		userControllerTest{passwordIsTooShort, http.StatusBadRequest, false},
	}
	for _, test := range tests {
		rec := httptest.NewRecorder()
		// When
		ctx := preparePostContext("/signup", test.input, rec)
		err := c.signup(ctx)
		// Then
		if test.happy {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.expectedStatusCode, rec.Code)
	}
}

func TestUserSignIn(t *testing.T) {
	c, _ := provideUserController("123", "123", 12*time.Hour)
	tests := []userControllerTest{
		userControllerTest{validUserSignIn, http.StatusCreated, true},
		userControllerTest{noUsernameField, http.StatusBadRequest, false},
		userControllerTest{noPasswordField, http.StatusBadRequest, false},
		userControllerTest{noRequiredFields, http.StatusBadRequest, false},
		userControllerTest{emptyPayload, http.StatusBadRequest, false},
		userControllerTest{usernameIsTooShort, http.StatusBadRequest, false},
		userControllerTest{passwordIsTooShort, http.StatusBadRequest, false},
	}
	for _, test := range tests {
		rec := httptest.NewRecorder()
		// When
		ctx := preparePostContext("/signin", test.input, rec)
		err := c.signin(ctx)
		// Then
		if test.happy {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.expectedStatusCode, rec.Code)
	}
}
