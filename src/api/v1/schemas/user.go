package schemas

type UserSignInSchema struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignUpSchema struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	UserName    string `json:"user_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
}

type LinkAccountRequisitionSchema struct {
	UserName      string `json:"user_name" binding:"required"`
	IdComputerLab string `json:"idComputerLab" binding:"required"`
}

type LinkAccountConfirmationSchema struct {
	Code string `json:"code" binding:"required"`
}
