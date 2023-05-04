package test

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/ZeinaMFI/cmd/server"
	"github.com/michaelgbenle/ZeinaMFI/internal/api"
	"github.com/michaelgbenle/ZeinaMFI/internal/middleware"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("./../../../.env"); err != nil {
		log.Println(err.Error())
	}
	os.Exit(m.Run())
}

func TestDepositHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)
	adminEmail := "kukus@yahoo.com"
	accClaim, _ := middleware.GenerateClaims(adminEmail)

	secret := os.Getenv("JWT_SECRET")
	acc, err := middleware.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	if err != nil {
		t.Fail()
	}

	testUser := &models.User{
		Email:     "love@aol.com",
		UserType:  "customer",
		AccountNo: "1234567890",
	}

	deposit := &models.Money{
		AccountNo: "1234567890",
		Amount:    7000,
	}
	transaction := &models.Transaction{
		AccountNo:       "1234567890",
		UserEmail:       testUser.Email,
		TransactionType: "deposit",
	}
	bodyJSON, err := json.Marshal(deposit)
	if err != nil {
		t.Fail()
	}

	t.Run("Successful Request", func(t *testing.T) {
		mockRepo.EXPECT().FindUserByEmail(adminEmail).Return(testUser, nil)
		mockRepo.EXPECT().FindUserByAccountNos(deposit.AccountNo).Return(testUser, nil)
		mockRepo.EXPECT().Deposit(deposit, testUser).Return(transaction, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/admin/deposit", strings.NewReader(string(bodyJSON)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
		router.ServeHTTP(rw, req)

		assert.Equal(t, 200, rw.Code)
		assert.Contains(t, rw.Body.String(), "account credited successfully")
	})

}
