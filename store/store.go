package storage

import (
	"calculationengine/models"
	"context"
	// "fmt"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *Store{
	return &Store{DB:db}
}

func (s *Store) GetAttributesByNames(ctx context.Context, categoryId int, names []string) []Attribute {
	var attributes []Attribute
	s.DB.Model(&CategoryAttributeAssignment{}).
		Select("attributes.*").
		Joins("JOIN attributes ON attributes.id = category_attribute_assignments.attribute_id AND category_attribute_assignments.category_id = ? AND attributes.name in ?", categoryId, names).
		Scan(&attributes)
	return attributes
}

func (s *Store) GetAllFormulaDependencies(ctx context.Context, categoryId int, targetAttributeId int) []models.AttributeDependencies {
	var formulaDependencies []models.AttributeDependencies
	s.DB.Model(&FormulaDependencies{}).
		Select("target_attribute_id, dependent_attribute_id").
		Where(&FormulaDependencies{CategoryID: uint(categoryId), TargetAttributeID: uint(targetAttributeId)}).
		Scan(&formulaDependencies)
	return formulaDependencies
}

// func (s *Store) SaveTopologicalSorting(ctx context.Context, categoryId int, topologicallySortedAttributes []int, formula string, dependentAttributeIds []int){
// 	err := s.DB.Transaction(func(tx *gorm.DB) error {
// 		gorm.G[CategoryAttributeAssignment](tx).Where("category_id = ?", categoryId).Update(ctx, "topological_sort_order", nil)
// 		for index, value := range topologicallySortedAttributes{
// 			gorm.G[CategoryAttributeAssignment](tx).Where("category_id = ? AND attribute_id = ?", categoryId, value).Update(ctx, "topological_sort_order", index)
// 		}
// 		return nil
// 	})
// 	if err !=nil {

// 	}
// }

func (s *Store) SaveFormula(ctx context.Context, params models.SaveFormulaParams) error {
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		gorm.G[CategoryAttributeAssignment](tx).Where("category_id = ?", params.CategoryID).Update(ctx, "topological_sort_order", nil)
		for index, value := range params.TopologicallySortedAttributeIDs{
			gorm.G[CategoryAttributeAssignment](tx).Where("category_id = ? AND attribute_id = ?", params.CategoryID, value).Update(ctx, "topological_sort_order", index)
		}
		tx.Save(&Formulas{
			CategoryID: uint(params.CategoryID),
			Expression: params.Formula,
			TargetAttributeID: uint(params.TargetAttributeID),
		})
		var formulaDependencies []FormulaDependencies
		for _, dependentAttributeId := range params.DependentAttributeIDs{
			formulaDependency := FormulaDependencies{
				CategoryID: uint(params.CategoryID),
				TargetAttributeID: uint(params.TargetAttributeID),
				DependentAttributeID: uint(dependentAttributeId),
			}
			formulaDependencies = append(formulaDependencies, formulaDependency)	
		}
		tx.Where(&FormulaDependencies{
			CategoryID: uint(params.CategoryID),
			TargetAttributeID: uint(params.TargetAttributeID),
		}).Delete(&FormulaDependencies{})

		result := tx.Create(&formulaDependencies)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

