package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
	"net/http"
)

func (u *HTTPHandler) TransactionsHandler(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	accountNo := c.Param("account_no")
	transactions, err := u.Repository.GetTransactions(accountNo)
	if err != nil {
		util.Response(c, "error", 500, nil, []string{"unable to get transactions"})
		return
	}

	util.Response(c, "success", 200, transactions, nil)
}
