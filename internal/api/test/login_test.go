package test

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/michaelgbenle/ZeinaMFI/cmd/server"
	"github.com/michaelgbenle/ZeinaMFI/internal/api"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/repository/mocks"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)

	loginRequest := &models.LoginRequest{
		Email:    "wenddy@ajah.com",
		Password: "Wenddy@123",
	}
	hashedPass, _ := util.HashPassword(loginRequest.Password)
	user := &models.User{
		Email:    "wenddy@ajah.com",
		Password: hashedPass,
		UserType: models.Admin,
	}
	bodyJSON, err := json.Marshal(loginRequest)
	if err != nil {
		t.Fail()
	}
	mockRepo.EXPECT().FindUserByEmail(loginRequest.Email).Return(user, nil)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(bodyJSON)))
	router.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)
	assert.Contains(t, rw.Body.String(), "login successful")
}
