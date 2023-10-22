package usecase

import (
	"context"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"sync"
)

type complainUseCase struct {
	complainRepository service.ComplainRepository
}

func NewComplainUseCase(complainRepository service.ComplainRepository) service.ComplainUseCase {
	return &complainUseCase{
		complainRepository: complainRepository,
	}
}

func (c *complainUseCase) GetAllComplain(ctx context.Context, limit int) (responses []models.Complain, err error) {
	var wg sync.WaitGroup

	responses, err = c.complainRepository.GetAllComplain(ctx, limit)
	if err != nil {
		err = fmt.Errorf("[Usecase][GetAllComplain] while get all complain data: %s", err.Error())
		return
	}

	for index, response := range responses {
		wg.Add(1)
		go func(index int, complainID int64) {
			defer wg.Done()
			var resolution []models.Resolution
			resolution, err = c.complainRepository.GetResolutionByComplainID(ctx, complainID)
			if err != nil {
				err = fmt.Errorf("[Usecase][GetAllComplain] while get all resolution complain data: %s", err.Error())
				return
			}
			responses[index].Resolution = resolution
		}(index, response.ID)
	}
	wg.Wait()
	return
}

func (c *complainUseCase) GetComplainByID(ctx context.Context, ID int64) (response models.Complain, err error) {
	response, err = c.complainRepository.GetComplainByID(ctx, ID)
	if err != nil {
		err = fmt.Errorf("[Usecase][GetComplainByID] while get all complain by id: %s", err.Error())
		return
	}

	var resolution []models.Resolution
	resolution, err = c.complainRepository.GetResolutionByComplainID(ctx, response.ID)
	if err != nil {
		err = fmt.Errorf("[Usecase][GetComplainByID] while get resolution complain data: %s", err.Error())
		return
	}
	response.Resolution = resolution

	return
}

func (c *complainUseCase) InsertComplain(ctx context.Context, args models.Complain, userID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *complainUseCase) InsertResolution(ctx context.Context, args models.Resolution, complainID, adminID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *complainUseCase) UpdateStatus(ctx context.Context, status string) (err error) {
	//TODO implement me
	panic("implement me")
}
