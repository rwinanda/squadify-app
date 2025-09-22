package response

type GetProfileResponse struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Address    string  `json:"address"`
	Gender     int64   `json:"gender"`
}
