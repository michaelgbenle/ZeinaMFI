package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
	"log"
	"net/http"
)

//if i had more time, i would have implemented the time duration period for locking savings
func (u *HTTPHandler) LockSavingsHandler(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	lock := &models.Money{}
	err = c.ShouldBindJSON(&lock)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//validate amount
	if lock.Amount <= 0 {
		util.Response(c, "error", 400, nil, []string{"invalid amount"})
		return
	}

	//check if account number exists
	user, err := u.Repository.FindUserByAccountNos(lock.AccountNo)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"account number does not exist"})
		return
	}

	//check if user has enough money
	if util.InsufficientBalance(user.AvailableBalance, lock.Amount) {
		util.Response(c, "error", 400, nil, []string{"insufficient balance"})
		return
	}

	//lock user's money
	transaction, lockErr := u.Repository.LockSavings(lock, user)
	if lockErr != nil {
		log.Println(lockErr)
		util.Response(c, "unable to lock user savings", 500, nil, []string{"lock error"})
		return
	}

	util.Response(c, "amount locked successfully", 200, transaction, nil)

}
