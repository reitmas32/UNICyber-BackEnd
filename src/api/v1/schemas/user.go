package schemas

type UserSignInSchema struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserSignUpSchema struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
