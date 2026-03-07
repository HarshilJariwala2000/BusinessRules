package storage

import (
	"time"
)
type JSONB  map[string]any

type Attribute struct{
	ID  uint   `gorm:"primaryKey; autoIncrement" json:",omitempty"`
	Name string `gorm:"uniqueIndex;not null" validate:"required"`
	DataType string `gorm:"not null" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type Formulas struct{
	CategoryID uint `gorm:"primaryKey"`
	Category CategoryAttributeAssignment `gorm:"foreignKey:CategoryID; references:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TargetAttributeID uint `gorm:"primaryKey"`
	Attribute CategoryAttributeAssignment `gorm:"foreignKey:TargetAttributeID; references:AttributeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Expression string
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type FormulaDependencies struct{
	CategoryID uint `gorm:"primaryKey"`
	Category Formulas `gorm:"foreignKey:CategoryID; references:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TargetAttributeID uint `gorm:"primaryKey"`
	TargetAttribute Formulas `gorm:"foreignKey:TargetAttributeID; references:TargetAttributeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DependentAttributeID uint `gorm:"primaryKey"`
	DependentAttribute Attribute `gorm:"foreignKey:DependentAttributeID; references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type Category struct{
	ID uint `grom:"primaryKey;autoIncrement" json:",omitempty"`
	Path string `gorm:"uniqueIndex; not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
	Attributes []Attribute `gorm:"many2many:category_attribute_assignments;"`
}

type CategoryAttributeAssignment struct {
	CategoryID uint `gorm:"primaryKey"`
	Category Category `gorm:"foreignKey:CategoryID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	AttributeID uint `gorm:"primaryKey"`
	Attribute Attribute `gorm:"foreignKey:AttributeID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	TopologicalSortOrder uint
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type Product struct{
	ID string `grom:"primaryKey" json:",omitempty"`
	CategoryID uint `gorm:"primaryKey"`
	Category CategoryAttributeAssignment `gorm:"foreignKey:CategoryID; references:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AttributeID uint `gorm:"primaryKey"`
	Attribute CategoryAttributeAssignment `gorm:"foreignKey:AttributeID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	Data string
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type ApiResponse struct{
	Message string `json:"message"`
	Data []any `json:"data"`
}


func AutoMigrate(){
	// DB.AutoMigrate(&Attribute{}, &Category{}, &Formulas{}, &CategoryAttributeAssignment{}, &Product{}, &FormulaDependencies{})
	DB.AutoMigrate(&Product{})
}

