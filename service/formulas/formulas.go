package formulas

import (
	"calculationengine/models"
	storage "calculationengine/store"
	// "context"
	// "gorm.io/gorm"
)

type Formula struct {
	categoryId int
	formula string
	attributes []string
}

func CreateFormula(request models.CreateFormulaRequest) (*storage.ApiResponse, error) {

	// formula := &Formula{formula:request.Formula, categoryId: request.CategoryID}
	// formula.setAttributes()
	// formula.validateAttributes()


	return &storage.ApiResponse{Message: "success", Data: []any{}}, nil

}

func (formula *Formula) setAttributes() {

}

func (formula *Formula) validateAttributes() {

}