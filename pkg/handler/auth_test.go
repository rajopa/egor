package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egor/watcher/pkg/handler"
	domain "github.com/egor/watcher/pkg/model"
	"github.com/egor/watcher/pkg/service"
	service_mocks "github.com/egor/watcher/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock" // Используем новую библиотеку
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockAuthorization, user domain.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "password": "qwerty"}`,
			inputUser: domain.User{
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user domain.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Wrong Input",
			inputBody: `{"username": "username"}`,
			inputUser: domain.User{},
			mockBehavior: func(r *service_mocks.MockAuthorization, user domain.User) {
				// Никаких вызовов не ожидается
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			h := handler.NewHandler(services)

			r := gin.New()

			r.POST("/sign-up", h.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
