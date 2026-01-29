package router

type CreateAttributeRequest struct{
	Name string `validate:"required"`
	DataType string `validate:"required"`
}

type CreateFormulaRequest struct{
	TargetAttribute string
	FormulaType string
	Formula string `validate:"oneof=calculation"`
}