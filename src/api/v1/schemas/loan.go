package schemas

type LoanCreateSchema struct {
	IdStudent  uint `json:"id_student" binding:"required"`
	IdComputer uint `json:"id_computer" binding:"required"`
}

type LoanCreateByAccountNumberSchema struct {
	AccountNumber string `json:"account_number" binding:"required"`
	IdComputer    uint   `json:"id_computer" binding:"required"`
}

type LoanUpdateSchema struct {
	// Metadata

	// Data
	IdStudent  uint `json:"id_student"`
	IdComputer uint `json:"id_computer"`
}
