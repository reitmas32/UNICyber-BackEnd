package schemas

type UserSignInSchema struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserSignInSchema) IsValid() bool {
	if u.UserName == "" || u.Password == "" {
		return false
	}
	return true
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

func (u *UserSignUpSchema) IsValid() bool {
	if u.UserName == "" || u.Password == "" || u.Email == "" || u.LastName == "" || u.Name == "" || u.DateOfBirth == "" || u.Role == "" || u.PhoneNumber == "" {
		return false
	}
	return true
}

type LinkAccountRequisitionSchema struct {
	UserName      string `json:"user_name" binding:"required"`
	IdComputerLab string `json:"idComputerLab" binding:"required"`
}

func (l *LinkAccountRequisitionSchema) IsValid() bool {
	if l.UserName == "" || l.IdComputerLab == "" {
		return false
	}
	return true
}

type LinkAccountConfirmationSchema struct {
	Code     string `json:"code" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
}

func (l *LinkAccountConfirmationSchema) IsValid() bool {
	if l.UserName == "" || l.Code == "" {
		return false
	}
	return true
}
