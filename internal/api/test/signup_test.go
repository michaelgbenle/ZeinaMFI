package test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/michaelgbenle/ZeinaMFI/cmd/server"
	"github.com/michaelgbenle/ZeinaMFI/internal/api"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)

	user := &models.User{
		Email:    "wenddy@ajah.com",
		Password: "Wenddy@123",
		UserType: "admin",
	}

	bodyJSON, err := json.Marshal(user)
	if err != nil {
		t.Fail()
	}
	t.Run("Successful Request", func(t *testing.T) {
		mockRepo.EXPECT().FindUserByEmail(user.Email).Return(nil, errors.New("user not found"))
		mockRepo.EXPECT().CreateUser(user).Return(nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(bodyJSON)))
		router.ServeHTTP(rw, req)

		assert.Equal(t, 201, rw.Code)
		assert.Contains(t, rw.Body.String(), "sign up successful")
	})
}
