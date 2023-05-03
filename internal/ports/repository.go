package ports

import "github.com/michaelgbenle/ZeinaMFI/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	TokenInBlacklist(token *string) bool
	FindUserById(Id string) (*models.User, error)
	FindUserByAccountNos(account string) (*models.User, error)
	CreateUser(user *models.User) error
	Deposit(money *models.Money, creditor *models.User) (*models.Transaction, error)
	Withdraw(money *models.Money, user *models.User) (*models.Transaction, error)
	GetAllUsers() (*[]models.User, error)
	GetTransactions(accountNo string) (*[]models.Transaction, error)
	LockSavings(money *models.Money, user *models.User) (*models.Transaction, error)
}
