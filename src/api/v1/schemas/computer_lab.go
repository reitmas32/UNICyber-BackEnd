package schemas

type ComputerLabCreateSchema struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ComputerLabFindSchema struct {
	Name string `json:"name" binding:"required"`
}

type ComputerLabUpdateSchema struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
