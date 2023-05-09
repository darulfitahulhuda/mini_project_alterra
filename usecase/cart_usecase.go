package usecase

import (
	"main/models"
	"main/repository"
	"strconv"
)

type CartUsecase interface {
	CreateCart(userId int, payload models.Carts) (models.Carts, error)
	GetAllCarts(userId int) ([]models.Carts, error)
	UpdateCart(id, userId int, payload models.Carts) (models.Carts, error)
	DeleteCartItem(id int) error
}

type cartUsecase struct {
	cartRepo  repository.CartRepository
	shoesRepo repository.ShoesRepository
}

func NewCartUsecase(cartRepo repository.CartRepository, shoesRepo repository.ShoesRepository) *cartUsecase {
	return &cartUsecase{cartRepo: cartRepo, shoesRepo: shoesRepo}
}

func (u *cartUsecase) CreateCart(userId int, payload models.Carts) (models.Carts, error) {
	payload.UserId = uint(userId)

	cart, err := u.cartRepo.CreateCart(payload)
	if err != nil {
		return models.Carts{}, err
	}

	shoes, err := u.shoesRepo.GetDetailShoes(int(cart.ShoesId))

	if err != nil {
		return models.Carts{}, err
	}
	cart.Shoes = shoes

	return cart, nil
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
			v.Status = "Available"
		} else if data.Qty <= 0 {
			v.Status = "Not available"
		} else {
			v.Status = "Availabe only " + strconv.Itoa(data.Qty)
		}

		carts = append(carts, v)
	}

	return carts, nil

}
func (u *cartUsecase) UpdateCart(id, userId int, payload models.Carts) (models.Carts, error) {
	cartByID, _ := u.cartRepo.GetCartById(id)

	shoes, err := u.shoesRepo.GetDetailShoes(int(cartByID.ShoesId))
	status := "Available"

	if err != nil {
		return models.Carts{}, err
	}

	for _, v := range shoes.Sizes {
		if v.Size == payload.Size {
			if v.Qty >= payload.Qty {
				status = "Available"
			} else if v.Qty <= 0 {
				status = "Not available"
			} else {
				status = "Availabe only " + strconv.Itoa(v.Qty)
			}

		}

	}

	data := models.Carts{
		UserId:  uint(userId),
		ShoesId: uint(payload.ShoesId),
		Qty:     payload.Qty,
		Size:    payload.Size,
		Status:  status,
	}

	result, err := u.cartRepo.UpdateCart(id, data)

	if err != nil {
		return result, err
	}

	result.Shoes = shoes

	return result, nil

}

func (u *cartUsecase) DeleteCartItem(id int) error {
	err := u.cartRepo.DeleteCartItem(id)
	if err != nil {
		return err
	}
	return nil
}
