package schemas

type StudentCreateSchema struct {
	//Info Personal
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"email"`

	//Info Academic
	IdUniversityProgram uint   `json:"id_university_program" binding:"required"`
	AccountNumber       string `json:"account_number" binding:"required,len=9,numeric"`
	Semester            uint8  `json:"semester" binding:"required"`
}

type StudentUpdateSchema struct {
	//Info Personal
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email" binding:"email"`

	//Info Academic
	IdUniversityProgram uint   `json:"id_university_program"`
	UniversityProgram   string `json:"university_program"`
	AccountNumber       string `json:"account_number" binding:",len=9,numeric"`
	Semester            uint8  `json:"semester"`
}
