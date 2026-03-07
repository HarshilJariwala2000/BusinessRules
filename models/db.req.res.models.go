package models

type DataType string

const (
    Integer DataType = "integer"
    String  DataType = "string"
    Boolean DataType = "boolean"
	Float 	DataType = "float"
)

type AttributeDependenciesResult struct {
	TargetAttributeID int `json:"target_attribute_id"`
	DependentAttributeID int `json:"dependent_attribute_id"`
	CategoryID int `json:"category_id"`
}

type ProductDatasResult struct {
	CategoryID int `json:"category_id"`
	ID string `json:"id"`
	AttributeID int `json:"attribute_id"`
	Data string `json:"data"`
	DataType DataType `json:"data_type"`
	AttributeName string `json:"attribute_name"`
}

type ProductListResult struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CategoryID int `json:"category_id"`
	CategoryPath string `json:"category_path"`
}

type GetCategoriesResult struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type GetCategoryWiseCommonAttributesResult struct {
	ID int `json:"id"`
	Name string `json:"name"`
	DataType string `json:"dataType"`
	Assigned bool `json:"assigned"`
}

type GetAttributesResult struct {
	ID int `json:"id"`
	Name string `json:"name"`
	DataType string `json:"dataType"`
}

type FormulasResult struct {
	CategoryID int `json:"category_id"`
	TargetAttributeID int `json:"target_attribute_id"`
	Expression string `json:"expression"`
}

type SaveFormulaParams struct {
	CategoryID int
	TopologicallySortedAttributeIDs []int
	Formula string
	DependentAttributeIDs []int
	TargetAttributeID int
}

type CreateProductParams struct {
	ID string
	CategoryID uint
	AttributeID uint
	Data string
}

type TopologicalSortResult struct {
	CategoryID int `json:"category_id"`
	AttributeID int `json:"attribute_id"`
	TopologicalSortOrder int `json:"topological_sort_order"`
}


