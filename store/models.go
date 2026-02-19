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
	ID uint `gorm:"primaryKey; autoIncrement"`
	CategoryID uint `gorm:"not null"`
	Expression string
	Catgeory Category `gorm:"foreignKey:CategoryID; references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TargetAttributeID uint `gorm:"not null"`
	Attribute Attribute `gorm:"foreignKey:TargetAttributeID; references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type FormulaDependencies struct{
	FormulaID uint `gorm:"primaryKey; not null;"`
	Formula Formulas `gorm:"foreignKey:FormulaID; references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	DependentAttributeID uint `gorm:"not null"`
	Attribute Attribute `gorm:"foreignKey:DependentAttributeID; references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
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
}

type Product struct{
	ID uint `grom:"primaryKey; autoIncrement" json:",omitempty"`
	Data JSONB `gorm:"type:jsonb; default:'{}'"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:",omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:",omitempty"`
	DeletedAt *time.Time `gorm:"index" json:",omitempty"`
}

type ApiResponse struct{
	Message string `json:"message"`
	Data []any `json:"data"`
}


func AutoMigrate(){
	DB.AutoMigrate(&Attribute{}, &Category{}, &Product{}, &FormulaDependencies{}, &Formulas{})
	// DB.AutoMigrate(&Product{})

}

