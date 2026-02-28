package formulas

import (
	"calculationengine/models"
	"calculationengine/service/dag"
	"calculationengine/service/parser"
	"calculationengine/service/utils"
	storage "calculationengine/store"
	"context"
	// "errors"	
	"fmt"
	"slices"
	// "gorm.io/gorm"
)

type Formula struct {
	categoryId int
	formula string
	attributes []string
}

func CreateFormula(ctx context.Context, request models.CreateFormulaRequest) (*storage.ApiResponse, error) {
	attributesInFormula, err := getAttributesFromFormula(request.Formula)
	if err != nil {
		return &storage.ApiResponse{Message: err.Error(), Data: []any{}}, nil
	}
	attributesInFormula = utils.RemoveArrayDuplicates(attributesInFormula)

	var store = storage.NewStore(storage.DB)
	attributesAssignedToCategory := store.GetAttributesByNames(ctx, request.CategoryID, attributesInFormula)
	attributeNameAssignedToCategory := utils.Map(attributesAssignedToCategory, func(a storage.Attribute) string {
		return a.Name
	})

	attributesNotExisting := utils.ArrayDifference(attributesInFormula, attributeNameAssignedToCategory)
	
	if(len(attributesNotExisting)!=0){
		return &storage.ApiResponse{Message: fmt.Sprintf("%s Attributes does not exist", attributesNotExisting), Data: []any{}}, nil
		//!Handle error
	}
	
	attributeDependencies := store.GetAllFormulaDependencies(ctx, request.CategoryID, request.TargetAttribute)
	
	//^Remove duplicates from this array
	var attributeIdsInFormula []int
	for _, attribute := range attributesAssignedToCategory {
		if(slices.Contains(attributesInFormula, attribute.Name)){
			attributeIdsInFormula = append(attributeIdsInFormula, int(attribute.ID))
		}
	}

	for _, dependentAttributeId := range attributeIdsInFormula{
		isDependencyExisting := false
		for _, attributeDependency := range attributeDependencies {
			if(attributeDependency.TargetAttributeID==request.TargetAttribute && attributeDependency.DependentAttributeID==dependentAttributeId){
				isDependencyExisting = true
				break
			}
		}
		if(isDependencyExisting){
			continue
		}
		attributeDependencies = append(attributeDependencies, models.AttributeDependencies{TargetAttributeID: request.TargetAttribute, DependentAttributeID: dependentAttributeId})
	}

	graph := generateGraphFromDependencies(attributeDependencies)
	topologicalSortedAttributes, hasCycle := graph.TopologicalSort()
	if(hasCycle){
		fmt.Println("Has Cycle")
		return &storage.ApiResponse{Message: "Cycle Detected", Data: []any{}}, nil
	}

	err1 := store.SaveFormula(ctx, models.SaveFormulaParams{
		CategoryID: request.CategoryID,
		TopologicallySortedAttributeIDs:topologicalSortedAttributes,
		Formula:request.Formula,
		DependentAttributeIDs:attributeIdsInFormula,
		TargetAttributeID:request.TargetAttribute,
	})
	if(err1!=nil){
		return &storage.ApiResponse{Message: "SOmething Went wrong", Data: []any{}}, nil
	}
	return &storage.ApiResponse{Message: "success", Data: []any{}}, nil
}

func generateGraphFromDependencies(attributeDependencies []models.AttributeDependencies) *dag.GraphList{
	graph := dag.NewGraphList()
	for _, value := range attributeDependencies {
		graph.AddEdge(value.DependentAttributeID, value.TargetAttributeID)
	}
	return graph
}

func getAttributesFromFormula(formula string) ([]string, error) {
	var attributesInFormula []string
	lexer := parser.NewLexer(formula)
	illegalTokens := []string{}
	formulaTokenLoop:for {
		var token = lexer.NextToken()
		switch token.TokenType {
		case parser.EOF:
			break formulaTokenLoop
		case parser.ILLEGAL:
			illegalTokens = append(illegalTokens, token.TokenValue)
			break formulaTokenLoop
			//!Handle error
		case parser.IDENT:
			attributesInFormula = append(attributesInFormula, token.TokenValue)
		}
	}
	if(len(illegalTokens)!=0){
		return []string{}, fmt.Errorf("%s nt allowed in formula", illegalTokens)
	}
	return attributesInFormula, nil
}

func (formula *Formula) validateAttributes() {

}