package models

import "time"

type UserType string

const (
	Admin    UserType = "admin"
	Customer UserType = "customer"
)

type User struct {
	ID               string    `json:"id"`
	Email            string    `json:"email" binding:"required"`
	Password         string    `json:"password" binding:"required"`
	UserType         UserType  `json:"user_type" binding:"required"`
	AccountNo        string    `json:"account_no"`
	AvailableBalance float64   `json:"available_balance"`
	LockedBalance    float64   `json:"locked_balance"`
	CreatedAt        time.Time `json:"created_at"`
}

type Balance struct {
	Available float64 `json:"available"`
	Locked    float64 `json:"locked"`
}

type Transaction struct {
	TransactionID    string    `json:"transaction_id"`
	UserEmail        string    `json:"user_email"`
	AccountNo        string    `json:"account_no"`
	AvailableBalance float64   `json:"available_balance"`
	LockedBalance    float64   `json:"locked_balance"`
	TransactionType  string    `json:"transaction_type"`
	Success          bool      `json:"success"`
	CreatedAt        time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Money struct {
	AccountNo  string  `json:"account_no" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	LockPeriod int     `json:"lock_period"`
}

type Blacklist struct {
}
