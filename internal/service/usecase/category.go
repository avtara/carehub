package usecase

import (
	"context"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"sync"
)

type categoryUseCase struct {
	categoryRepository service.CategoryRepository
}

func NewCategoryUseCase(categoryRepository service.CategoryRepository) service.CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (u categoryUseCase) GetAllCategories(ctx context.Context, limit int) (response []models.Category, err error) {
	var (
		wg sync.WaitGroup
	)

	response, err = u.categoryRepository.GetAllCategories(ctx, limit)
	if err != nil {
		err = fmt.Errorf("[Usecase][GetAllCategories] failed while get all categories: %s", err.Error())
		return
	}

	for index, category := range response {
		wg.Add(1)
		go func(index int, categoryID int64) {
			defer wg.Done()
			var extraFields []models.ExtraFieldCategory
			extraFields, err = u.categoryRepository.GetExtraFieldByCategoryID(context.Background(), categoryID)
			if err != nil {
				err = fmt.Errorf("[Usecase][GetAllCategories] failed while get all categories: %s", err.Error())
				return
			}
			response[index].ExtraFieldCategories = extraFields
		}(index, category.ID)
	}
	wg.Wait()

	return
}

func (u categoryUseCase) GetCategoryByID(ctx context.Context, ID int64) (response models.Category, err error) {

	response, err = u.categoryRepository.GetCategoryByID(ctx, ID)
	if err != nil {
		err = fmt.Errorf("[Usecase][GetCategoryByID] failed while category by id: %s", err.Error())
		return
	}

	var extraFields []models.ExtraFieldCategory
	extraFields, err = u.categoryRepository.GetExtraFieldByCategoryID(ctx, response.ID)
	if err != nil {
		err = fmt.Errorf("[Usecase][GetCategoryByID] failed while get category extrafield by category id: %s", err.Error())
		return
	}
	response.ExtraFieldCategories = extraFields

	return
}
