package service

import (
	"calculationengine/store"
	"context"
	"gorm.io/gorm"
)

func CreateAttribute(request storage.Attribute) (*storage.ApiResponse, error) {
	result := gorm.WithResult()
	ctx := context.Background()
	err := gorm.G[storage.Attribute](storage.DB, result).Create(ctx, &request)
	if err!=nil{
		return &storage.ApiResponse{Message:"Something Went Wrong", Data:[]any{}}, err
	}
	return &storage.ApiResponse{Message:"success", Data:[]any{}}, nil
}