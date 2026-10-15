package model

import "time"

// a central model file defines our pure Go structs. These are not teid to any database library or Http framewrok
//user represents an authenticated applicaiton user
type User stuct {
	ID int64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"-"` //The "-" tag ensures the hash is Never serialized to Jason
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at`
}

//Product is the core domain entity
type Product struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	Name string `json:"name"`
	Description sring `json:"description.omitempty"`
	Price float64 `json;"price"`
	Stock int `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type RegisterRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateProductRequest struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}

//UpdateProductRequest user pointer fields so absent json keys are distinguishable
//from zero values.only no-nil fileds are written to the database on PATCH

type updateProductRequest struct {
	Name *string `json:"name"`
	Description *string `json:"description"`
	Price *float64 `json:"price"`
	Stock *int `json:"stock"`
}

//response payloads
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	User *User `json:"user"`
}

//APIError is the standard error response format for the API
type APIError struct {
	Code int `json:"code"`
	Message string `json:"message`
}

//PaginatedResponse wraps a list of items with paginaiton meta data
type PaginatedResponse[T any] struct {
	Data []T `json:"data"`
	Total int64 `json:"total"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}

//Claims represents the data embedded inside our JWT
type Claims struct {
	UserID int64 `json:"user_id"`
	Email string `json:"email"`
	Role string `json:"role"`
}