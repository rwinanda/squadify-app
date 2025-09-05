package dto

type RegisterInput struct {
	Username   string  `json:"username" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Password   string  `json:"password" binding:"required"`
	FirstName  string  `json:"first_name" binding:"required"`
	MiddleName string  `json:"middle_name" binding:"required"`
	LastName   string  `json:"last_name" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	Address    string  `json:"address" binding:"required"`
	Gender     int64   `json:"gender" binding:"required"`
}
