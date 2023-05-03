package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
	"log"
	"net/http"
)

func (u *HTTPHandler) WithdrawHandler(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	withdraw := &models.Money{}
	err = c.ShouldBindJSON(&withdraw)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//validate amount
	if withdraw.Amount <= 0 {
		util.Response(c, "error", 400, nil, []string{"invalid amount"})
		return
	}

	//check if account number exists
	user, err := u.Repository.FindUserByAccountNos(withdraw.AccountNo)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"account number does not exist"})
		return
	}

	//check if user has enough money
	if util.InsufficientBalance(user.Balance.Available, withdraw.Amount) {
		util.Response(c, "error", 400, nil, []string{"insufficient balance"})
		return
	}

	//withdraw user's money
	transaction, withdrawErr := u.Repository.Withdraw(withdraw, user)
	if withdrawErr != nil {
		log.Println(withdrawErr)
		util.Response(c, "unable to withdraw from user account", 500, nil, []string{"withdraw error"})
		return
	}

	util.Response(c, "amount debited successfully", 200, transaction, nil)

}
