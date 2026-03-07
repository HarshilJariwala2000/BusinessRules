package category

import (
	"context"
	"calculationengine/models"
		storage "calculationengine/store"
)

func CreateCategory(ctx context.Context, request models.CreateCategoryRequest)(*storage.ApiResponse, error){
	s := storage.NewStore(storage.DB)
	err := s.CreateCategory(ctx, request.Name)
	if err != nil {
		return &storage.ApiResponse{Message: "Something went wrong", Data: []any{}}, nil
	}
	return &storage.ApiResponse{Message: "success", Data: []any{}}, nil
}

func GetAllCategories(ctx context.Context) (*models.GetCategoriesResponse, error) {
	response :=  models.GetCategoriesResponse{}
	s := storage.NewStore(storage.DB)
	data, err := s.GetAllCategories(ctx)
	if err != nil {
		response.Message = "Something went wrong"
		return &response, nil
	}
	response.Message = "Success"
	response.Data = data
	return &response, nil
}