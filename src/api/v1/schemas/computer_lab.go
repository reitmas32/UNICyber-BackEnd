package schemas

type ComputerLabCreateSchema struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (c *ComputerLabCreateSchema) isValid() bool {
	if c.Name == "" || c.Description == "" {
		return false
	}
	return true
}

type ComputerLabFindSchema struct {
	Name string `json:"name" binding:"required"`
}

func (c *ComputerLabFindSchema) isValid() bool {
	if c.Name == "" {
		return false
	}
	return true
}

type ComputerLabUpdateSchema struct {
	// Metadata
	//IdComputer string `gorm:"primaryKey"`
	//IdRoom     string

	// Data
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (c *ComputerLabUpdateSchema) isValid() bool {
	if c.Name == "" || c.Description == "" {
		return false
	}
	return true
}
