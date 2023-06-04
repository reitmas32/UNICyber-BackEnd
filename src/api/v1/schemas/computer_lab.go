package schemas

type ComputerLabSchema struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	JWT         string `json:"jwt" binding:"required"`
}

type ComputerLabCreateSchema struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ComputerLabFindSchema struct {
	Name string `json:"name" binding:"required"`
}

type ComputerLabUpdateSchema struct {
	// Metadata
	//IdComputer string `gorm:"primaryKey"`
	//IdRoom     string

	// Data
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
