package attribute

import (
	"calculationengine/store"
	"calculationengine/models"
	"context"
	"gorm.io/gorm"
)

func CreateAttribute(request models.CreateAttributeRequest) (*storage.ApiResponse, error) {
	result := gorm.WithResult()
	ctx := context.Background()

	createObject := &storage.Attribute{Name:request.Name, DataType: request.DataType}

	err := gorm.G[storage.Attribute](storage.DB, result).Create(ctx, createObject)
	if err!=nil{
		return &storage.ApiResponse{Message:"Something Went Wrong", Data:[]any{}}, err
	}
	return &storage.ApiResponse{Message:"success", Data:[]any{}}, nil
}

