package request

type CreateProfileRequest struct {
	FirstName  string  `json:"first_name" binding:"required"`
	MiddleName string  `json:"middle_name" binding:"required"`
	LastName   string  `json:"last_name" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	Address    string  `json:"address" binding:"required"`
	GenderID   int64   `json:"gender_id" binding:"required"`
}
