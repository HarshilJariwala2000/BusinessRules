package models

type CreateAttributeRequest struct{
	Name string `validate:"required"`
	DataType string `validate:"required"`
}

type CreateFormulaRequest struct{
	CategoryID int
	TargetAttribute int
	Formula string `validate:"oneof=calculation"`
}