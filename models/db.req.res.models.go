package models

type AttributeDependencies struct {
	TargetAttributeID int `json:"target_attribute_id"`
	DependentAttributeID int `json:"dependent_attribute_id"`
}

type SaveFormulaParams struct {
	CategoryID int
	TopologicallySortedAttributeIDs []int
	Formula string
	DependentAttributeIDs []int
	TargetAttributeID int
}


