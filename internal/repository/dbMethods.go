package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
	"log"
	"strconv"
	"time"
)

// TokenInBlacklist checks if token is already in the blacklist collection
func (p *Postgres) TokenInBlacklist(token *string) bool {
	tok := &models.Blacklist{}
	if err := p.DB.Where("token = ?", token).First(&tok).Error; err != nil {
		return false
	}
	return true
}

func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) FindUserById(Id string) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("id = ?", Id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) FindUserByAccountNos(account string) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("account_no = ?", account).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) CreateUser(user *models.User) error {
	pass, pErr := util.HashPassword(user.Password)
	if pErr != nil {
		log.Println("error in hashing password")
		return pErr
	}
	user.ID = uuid.New().String()
	user.Password = pass
	user.AccountNo = strconv.Itoa(util.GenerateAccountNumber())
	user.Balance.Available = 0
	user.Balance.Locked = 0
	user.CreatedAt = time.Now()

	err := p.DB.Create(&user).Error
	if err != nil {
		log.Println("error in creating user")
		return err
	}
	return nil
}

func (p *Postgres) GetAllUsers() (*[]models.User, error) {
	users := &[]models.User{}
	if err := p.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (p *Postgres) GetTransactions(user *models.User) (*[]models.Transaction, error) {
	transactions := &[]models.Transaction{}
	if err := p.DB.Where("account_nos= ?", user.AccountNo).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p *Postgres) Deposit(money *models.Money, user *models.User) (*models.Transaction, error) {

	//update user account
	err := p.DB.Model(&user).Update("available", user.Balance.Available+money.Amount).Error
	if err != nil {
		log.Println(err)
		transaction := &models.Transaction{
			TransactionID: uuid.NewString(),
			UserEmail:     user.Email,
			AccountNo:     user.AccountNo,
			Balance: models.Balance{
				Available: user.Balance.Available,
				Locked:    user.Balance.Locked,
			},
			TransactionType: "Deposit",
			Success:         false,
			CreatedAt:       time.Now(),
		}
		err = p.DB.Create(&transaction).Error
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return transaction, err
	}
	transaction := &models.Transaction{
		TransactionID: uuid.NewString(),
		UserEmail:     user.Email,
		AccountNo:     user.AccountNo,
		Balance: models.Balance{
			Available: user.Balance.Available,
			Locked:    user.Balance.Locked,
		},
		TransactionType: "Deposit",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	err = p.DB.Create(&transaction).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return transaction, nil
}

func (p *Postgres) Withdraw(money *models.Money, user *models.User) (*models.Transaction, error) {

	//update user account
	err := p.DB.Model(&user).Update("available", user.Balance.Available-money.Amount).Error
	if err != nil {
		log.Println(err)
		transaction := &models.Transaction{
			TransactionID: uuid.NewString(),
			UserEmail:     user.Email,
			AccountNo:     user.AccountNo,
			Balance: models.Balance{
				Available: user.Balance.Available,
				Locked:    user.Balance.Locked,
			},
			TransactionType: "Withdraw",
			Success:         false,
			CreatedAt:       time.Now(),
		}
		err = p.DB.Create(&transaction).Error
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return transaction, err
	}
	transaction := &models.Transaction{
		TransactionID: uuid.NewString(),
		UserEmail:     user.Email,
		AccountNo:     user.AccountNo,
		Balance: models.Balance{
			Available: user.Balance.Available,
			Locked:    user.Balance.Locked,
		},
		TransactionType: "Withdraw",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	err = p.DB.Create(&transaction).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return transaction, nil
}

func (p *Postgres) LockSavings(money *models.Money, user *models.User) (*models.Transaction, error) {

	//Begin transaction to lock savings
	err := p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user).Update("available", user.Balance.Available-money.Amount).Error; err != nil {
			log.Println(err)
			return err
		}

		if err := tx.Model(&user).Update("locked", user.Balance.Locked+money.Amount).Error; err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
	if err != nil {
		log.Println(err)
		transaction := &models.Transaction{
			TransactionID: uuid.NewString(),
			UserEmail:     user.Email,
			AccountNo:     user.AccountNo,
			Balance: models.Balance{
				Available: user.Balance.Available,
				Locked:    user.Balance.Locked,
			},
			TransactionType: "Lock Savings",
			Success:         false,
			CreatedAt:       time.Now(),
		}
		err = p.DB.Create(&transaction).Error
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return transaction, err
	}

	transaction := &models.Transaction{
		TransactionID: uuid.NewString(),
		UserEmail:     user.Email,
		AccountNo:     user.AccountNo,
		Balance: models.Balance{
			Available: user.Balance.Available,
			Locked:    user.Balance.Locked,
		},
		TransactionType: "Lock Savings",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	err = p.DB.Create(&transaction).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return transaction, err
}
