package usecase

import (
	"main/dto"
	"main/models"
	"main/repository"
)

type ShoesUsecase interface {
	CreateShoes(payload dto.ShoesRequest) (models.ShoesDetailData, error)
	CreateShoesSize(payload dto.ShoesSizeRequest) (models.ShoesSize, error)
	GetAllShoes(gender string) ([]models.ShoesListData, error)
	GetDetailShoes(id int) (models.ShoesDetailData, error)
	UpdateShoes(id int, payload dto.ShoesRequest) error
	DeleteShoes(id int) error
	DeleteShoesSize(payload dto.ShoesSizeRequest) error
}

type shoesUsecase struct {
	shoesRepository repository.ShoesRepository
}

func NewShoesUsecase(shoesRepo repository.ShoesRepository) *shoesUsecase {
	return &shoesUsecase{shoesRepository: shoesRepo}
}

func (u *shoesUsecase) CreateShoes(payload dto.ShoesRequest) (models.ShoesDetailData, error) {
	shoesSize := make([]models.ShoesSize, 0)

	for _, v := range payload.Sizes {
		shoesSize = append(shoesSize, models.ShoesSize{
			Qty:  v.Qty,
			Size: v.Size,
		})
	}
	shoes := models.Shoes{
		Name:   payload.Name,
		Gender: payload.Gender,
		Images: payload.Images,
		Price:  payload.Price,
		ShoesDetail: models.ShoesDetail{
			Description: payload.Description,
			Brand:       payload.Brand,
		},
		Sizes: shoesSize,
	}

	data, err := u.shoesRepository.CreateShoes(shoes)

	if err != nil {
		return models.ShoesDetailData{}, err
	}

	sizes := make([]dto.ShoesSizeRequest, 0)

	for _, v := range data.Sizes {
		sizes = append(sizes, dto.ShoesSizeRequest{
			Size:    v.Size,
			Qty:     v.Qty,
			ShoesId: int(v.ShoesId),
		})
	}

	detailShoes := models.ShoesDetailData{
		ID:          int(data.ID),
		Name:        data.Name,
		Images:      data.Images,
		Price:       data.Price,
		Gender:      data.Gender,
		Description: data.ShoesDetail.Description,
		Brand:       data.ShoesDetail.Brand,
		Sizes:       sizes,
	}
	return detailShoes, nil
}

func (u *shoesUsecase) CreateShoesSize(payload dto.ShoesSizeRequest) (models.ShoesSize, error) {

	data, err := u.shoesRepository.CreateShoesSize(models.ShoesSize{
		ShoesId: uint(payload.Size),
		Qty:     payload.Qty,
		Size:    payload.Size,
	})

	if err != nil {
		return data, err
	}
	return data, err

}

func (u *shoesUsecase) GetAllShoes(gender string) ([]models.ShoesListData, error) {
	shoes, err := u.shoesRepository.GetAllShoes(gender)

	listShoes := make([]models.ShoesListData, 0)

	for _, data := range shoes {
		listShoes = append(listShoes, models.ShoesListData{
			ID:     int(data.ID),
			Name:   data.Name,
			Images: data.Images,
			Price:  data.Price,
			Gender: data.Gender,
		})

	}

	if err != nil {
		return listShoes, err
	}
	return listShoes, nil
}

func (u *shoesUsecase) GetDetailShoes(id int) (models.ShoesDetailData, error) {
	shoes, err := u.shoesRepository.GetDetailShoes(id)
	if err != nil {
		return models.ShoesDetailData{}, err
	}

	sizes := make([]dto.ShoesSizeRequest, 0)

	for _, v := range shoes.Sizes {
		sizes = append(sizes, dto.ShoesSizeRequest{
			Size:    v.Size,
			Qty:     v.Qty,
			ShoesId: int(v.ShoesId),
		})

	}

	detailShoes := models.ShoesDetailData{
		ID:          int(shoes.ID),
		Name:        shoes.Name,
		Images:      shoes.Images,
		Price:       shoes.Price,
		Gender:      shoes.Gender,
		Description: shoes.ShoesDetail.Description,
		Brand:       shoes.ShoesDetail.Brand,
		Sizes:       sizes,
	}

	return detailShoes, nil
}

func (u *shoesUsecase) UpdateShoes(id int, payload dto.ShoesRequest) error {

	shoes := models.Shoes{
		Name:   payload.Name,
		Gender: payload.Gender,
		Images: payload.Images,
		Price:  payload.Price,
		ShoesDetail: models.ShoesDetail{
			Description: payload.Description,
			Brand:       payload.Brand,
		},
	}
	if err := u.shoesRepository.UpdateShoes(id, shoes); err != nil {
		return err
	}

	for _, v := range payload.Sizes {
		size := models.ShoesSize{
			ShoesId: uint(id),
			Size:    v.Size,
			Qty:     v.Qty,
		}
		if err := u.shoesRepository.UpdateShoesSize(size); err != nil {
			return err
		}
	}

	return nil
}

func (u *shoesUsecase) DeleteShoes(id int) error {
	if err := u.shoesRepository.DeleteShoes(id); err != nil {
		return err
	}
	if err := u.shoesRepository.DeleteShoesSize(models.ShoesSize{ShoesId: uint(id)}, true); err != nil {
		return err
	}
	return nil
}

func (u *shoesUsecase) DeleteShoesSize(payload dto.ShoesSizeRequest) error {
	data := models.ShoesSize{
		ShoesId: uint(payload.ShoesId),
		Size:    payload.Size,
	}
	if err := u.shoesRepository.DeleteShoesSize(data, false); err != nil {
		return err
	}
	return nil
}
