package models

type CreateAttributeRequest struct{
	Name string `json:"name" validate:"required"`
	DataType string `json:"dataType" validate:"required"`
}

type CreateFormulaRequest struct{
	CategoryID int `json:"categoryId" validate:"required"`
	TargetAttribute int `json:"targetAttribute" validate:"required"`
	Formula string `json:"formula" validate:"required"`
}