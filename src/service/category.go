package service

import (
	"context"
	"fmt"
	"io"

	"grpcTreino/src/api/models"
	"grpcTreino/src/grpc"
)

type CategoryService struct {
	grpc.UnimplementedCategoryServiceServer
	Category models.Category
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (c *CategoryService) GetAllCategory(ctx context.Context, request *grpc.Blank) (*grpc.CategoryResponse, error) {
	categories := c.Category.GetAllCategories()
	var categoryResponse []*grpc.Category
	for _, category := range categories {
		categoryResponse = append(categoryResponse, &grpc.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &grpc.CategoryResponse{Category: categoryResponse}, nil

}

func (c *CategoryService) CreateProtoCategory(ctx context.Context, request *grpc.CategoryRequest) (*grpc.Category, error) {
	c.Category.Name = request.Name
	c.Category.Description = request.Description
	result := c.Category.CreateCategory()
	categoryResponse := &grpc.Category{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	}
	fmt.Println(categoryResponse)
	return categoryResponse, nil
}

func (c *CategoryService) DeleteProtoCategory(ctx context.Context, request *grpc.Category) (*grpc.Category, error) {
	c.Category.ID = request.ID
	result := c.Category.DeleteCategory()
	categoryResponse := &grpc.Category{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) CreateProtoCategoryStream(stream grpc.CategoryService_CreateProtoCategoryStreamServer) error {
	categories := &grpc.CategoryResponse{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		c.Category.Name = category.Name
		c.Category.Description = category.Description
		result := c.Category.CreateCategory()
		categories.Category = append(categories.Category, &grpc.Category{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		})
	}
}

func (c *CategoryService) DeleteProtoCategoryStream(stream grpc.CategoryService_DeleteProtoCategoryStreamServer) error {
	categories := &grpc.CategoryResponse{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		c.Category.ID = category.ID
		result := c.Category.DeleteCategory()
		categories.Category = append(categories.Category, &grpc.Category{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		})
	}
}

func (c *CategoryService) CreateProtoCategoryServerStream(request *grpc.CategoryRequest, stream grpc.CategoryService_CreateProtoCategoryServerStreamServer) error {
	categories := &grpc.CategoryResponse{}
	c.Category.Name = request.Name
	c.Category.Description = request.Description
	result := c.Category.CreateCategory()
	categories.Category = append(categories.Category, &grpc.Category{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	})
	return stream.Send(categories)
}

func (c *CategoryService) CreateProtoCategoryBidirectionalStream(stream grpc.CategoryService_CreateProtoCategoryBidirectionalStreamServer) error {
	categories := &grpc.CategoryResponse{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		c.Category.Name = category.Name
		c.Category.Description = category.Description
		result := c.Category.CreateCategory()
		categories.Category = append(categories.Category, &grpc.Category{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		})
		if err := stream.Send(categories); err != nil {
			return err
		}
	}
}
