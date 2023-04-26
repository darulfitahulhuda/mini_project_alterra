package usecase

import (
	"main/dto"
	"main/models"
	"main/repository"
)

type ShoesUsecase interface {
	CreateShoes(payload dto.Shoes) error
	GetAllShoes() ([]models.ShoesListData, error)
	GetDetailShoes(id int) (models.ShoesDetailData, error)
	UpdateShoes(id int, payload dto.Shoes) error
	DeleteShoes(id int) error
}

type shoesUsecase struct {
	shoesRepository repository.ShoesRepository
}

func NewShoesUsecase(shoesRepo repository.ShoesRepository) *shoesUsecase {
	return &shoesUsecase{shoesRepository: shoesRepo}
}

func (u *shoesUsecase) CreateShoes(payload dto.Shoes) error {
	shoes := models.Shoes{
		Name:   payload.Name,
		Gender: payload.Gender,
		Images: payload.Images,
		Price:  payload.Price,
		ShoesDetail: models.ShoesDetail{
			Description: payload.Description,
			Category:    payload.Category,
			Color:       payload.Color,
			Size:        payload.Size,
			Qty:         payload.Qty,
			Brand:       payload.Brand,
		},
	}
	if err := u.shoesRepository.CreateShoes(shoes); err != nil {
		return err
	}
	return nil
}

func (u *shoesUsecase) GetAllShoes() ([]models.ShoesListData, error) {
	shoes, err := u.shoesRepository.GetAllShoes()

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

	detailShoes := models.ShoesDetailData{
		ID:          int(shoes.ID),
		Name:        shoes.Name,
		Images:      shoes.Images,
		Price:       shoes.Price,
		Gender:      shoes.Gender,
		Description: shoes.ShoesDetail.Description,
		Category:    shoes.ShoesDetail.Category,
		Color:       shoes.ShoesDetail.Color,
		Size:        shoes.ShoesDetail.Size,
		Qty:         shoes.ShoesDetail.Qty,
		Brand:       shoes.ShoesDetail.Brand,
	}

	if err != nil {
		return detailShoes, err
	}
	return detailShoes, nil
}

func (u *shoesUsecase) UpdateShoes(id int, payload dto.Shoes) error {
	shoes := models.Shoes{
		Name:   payload.Name,
		Gender: payload.Gender,
		Images: payload.Images,
		Price:  payload.Price,
		ShoesDetail: models.ShoesDetail{
			Description: payload.Description,
			Category:    payload.Category,
			Color:       payload.Color,
			Size:        payload.Size,
			Qty:         payload.Qty,
			Brand:       payload.Brand,
		},
	}
	if err := u.shoesRepository.UpdateShoes(id, shoes); err != nil {
		return err
	}

	return nil
}

func (u *shoesUsecase) DeleteShoes(id int) error {
	if err := u.shoesRepository.DeleteShoes(id); err != nil {
		return err
	}
	return nil
}
