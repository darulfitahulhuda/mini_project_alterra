package usecase

import (
	"main/dto"
	"main/models"
	"main/repository"
	"strconv"
)

type CartUsecase interface {
	CreateCart(userId int, payload dto.Cart) error
	GetAllCarts(userId int) ([]models.Carts, error)
	UpdateCart(id, userId int, payload dto.Cart) (models.Carts, error)
	DeleteCartItem(id int) error
}

type cartUsecase struct {
	cartRepo  repository.CartRepository
	shoesRepo repository.ShoesRepository
}

func NewCartUsecase(cartRepo repository.CartRepository, shoesRepo repository.ShoesRepository) *cartUsecase {
	return &cartUsecase{cartRepo: cartRepo, shoesRepo: shoesRepo}
}

func (u *cartUsecase) CreateCart(userId int, payload dto.Cart) error {
	data := models.Carts{
		UserId:  uint(userId),
		ShoesId: uint(payload.ShoesId),
		Size:    payload.Size,
		Status:  payload.Status,
		Qty:     payload.Qty,
	}

	if err := u.cartRepo.CreateCart(data); err != nil {
		return err
	}

	return nil
}
func (u *cartUsecase) GetAllCarts(userId int) ([]models.Carts, error) {
	carts := make([]models.Carts, 0)
	result, err := u.cartRepo.GetAllCarts(userId)

	if err != nil {
		return carts, err
	}

	for _, v := range result {
		shoesSize := models.ShoesSize{
			ShoesId: v.ShoesId,
			Size:    v.Size,
		}
		data, err := u.shoesRepo.GetShoesSize(shoesSize)

		if err != nil {
			return carts, err
		}

		if data.Qty >= v.Qty {
			v.Status = "available"
		} else if data.Qty <= 0 {
			v.Status = "not available"
		} else {
			v.Status = "availabe " + strconv.Itoa(data.Qty)
		}

		carts = append(carts, v)
	}

	return carts, nil

}
func (u *cartUsecase) UpdateCart(id, userId int, payload dto.Cart) (models.Carts, error) {
	data := models.Carts{
		UserId:  uint(userId),
		ShoesId: uint(payload.ShoesId),
		Qty:     payload.Qty,
		Size:    payload.Size,
		Status:  payload.Status,
	}

	result, err := u.cartRepo.UpdateCart(id, data)

	if err != nil {
		return result, err
	}

	return result, nil

}
func (u *cartUsecase) DeleteCartItem(id int) error {
	err := u.cartRepo.DeleteCartItem(id)
	if err != nil {
		return err
	}
	return nil
}
