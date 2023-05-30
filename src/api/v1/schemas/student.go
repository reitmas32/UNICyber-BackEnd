package schemas

type StudentCreateSchema struct {
	//Info Personal
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"email"`

	//Info Academic
	UniversityProgram string `json:"university_program" binding:"required"`
	AccountNumber     string `json:"account_number" binding:"required,len=9,numeric"`
	Semester          uint8  `json:"semester" binding:"required"`
}
