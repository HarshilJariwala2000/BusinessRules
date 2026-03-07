package product

import (
	"calculationengine/models"
	"calculationengine/service/formulas"
	"calculationengine/service/utils"
	storage "calculationengine/store"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func UpsertProduct(ctx context.Context, request models.CreateProductRequest)(*storage.ApiResponse, error){
	s := storage.NewStore(storage.DB)
	attributes, err := s.GetAttributesIdDataMap(ctx, []int{request.CategryID})
	var id string
	create := false
	if request.ProductID == "" {
		create = true
		id = uuid.NewString()
	}else {
		id = request.ProductID
	}
	if err != nil {
		return &storage.ApiResponse{Message: "Something went wrong", Data: []any{}}, nil
	}
	var createProductParams []models.CreateProductParams
	defaultAttributePresent := false
	for _, data := range request.ProductData {
		value := strings.TrimSpace(data.Value)
		if value == "" {
			continue
		} 
		if data.AttributeID==0 {
			defaultAttributePresent = true
		}
		attribute, ok := attributes[data.AttributeID]; 
		if !ok {
			return &storage.ApiResponse{Message: fmt.Sprintf("Attribute %d does not exist in this category", data.AttributeID), Data: []any{}}, nil
		}
		validationErr := validateData(value, attribute)
		if validationErr != nil {
			return &storage.ApiResponse{Message: validationErr.Error(), Data: []any{}}, nil
		}
		createProductParams = append(createProductParams, models.CreateProductParams{
			ID:id,
			CategoryID:uint(request.CategryID),
			AttributeID:uint(data.AttributeID),
			Data:data.Value,
		})
	}

	if create && !defaultAttributePresent {
		return &storage.ApiResponse{Message: "Please enter Product Name to create a new product", Data: []any{}}, nil
	}

	s.UpsertProduct(ctx, createProductParams)
	evaluateFormulaRequest := models.EvaluateFormulaRequest{
		ProductID: []string{id},
	}
	formulas.EvaluateFormula(ctx, evaluateFormulaRequest)
	return &storage.ApiResponse{Message: "success", Data: []any{}}, nil
}

func GetProductList(ctx context.Context) (*models.GetProductListResponse, error) {
	response := models.GetProductListResponse{}
	s := storage.NewStore(storage.DB)
	data, err := s.GetProductList(ctx)
	if err != nil {
		response.Message = "Something went wrong"
		return &response, nil
	}
	response.Message = "Success"
	response.Data = data
	return &response, nil
}

func GetSingleProductData(ctx context.Context, request models.GetProductDataRequest) (*models.GetProductDataResponse, error) {
	response := models.GetProductDataResponse{}
	s := storage.NewStore(storage.DB)
	data, err := s.GetProductData(ctx, []string{request.ProductID})
	if err != nil {
		response.Message = "Something went wrong"
		return &response, nil
	}
	response.Message = "Success"
	response.Data = data
	return &response, nil

}

func validateData(value string, attribute storage.Attribute) error {
	switch attribute.DataType{
	case "integer":
		_, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
	case "float":
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
	case "boolean":
		_, err := utils.StringToBoolean(value)
		if err != nil {
			return err
		}
	}
	return nil
}