package repository

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateShop(shop *models.Shop) error {
	err := r.pg.Session(&gorm.Session{FullSaveAssociations: true}).Create(shop).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllShops() (*[]models.Shop, error) {
	var shops []models.Shop
	err := r.pg.Preload("ShopOwner.User").Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return &shops, nil
}

func (r *Repository) GetShopById(id int32) (*models.Shop, error) {
	var shop models.Shop
	err := r.pg.Preload("ShopOwner.User").Where("id = ?", id).Take(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *Repository) GetOwnedShops(uid string) ([]models.Shop, error) {
	var shopOwners []models.ShopOwner
	err := r.pg.Preload("Shop").Where("owner_id = ?", uid).Find(&shopOwners).Error
	if err != nil {
		return nil, err
	}
	var shops []models.Shop
	for _, shopOwner := range shopOwners {
		shops = append(shops, shopOwner.Shop)
	}
	return shops, nil
}

func (r *Repository) EditShop(shop *models.Shop) (*models.Shop, error) {
	err := r.pg.Session(&gorm.Session{FullSaveAssociations: true}).Model(&shop).Clauses(clause.Returning{}).Updates(shop).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to update shop")
	}
	return shop, nil
}

func (r *Repository) EditShopOwners(shopId int32, shopOwners []models.ShopOwner) (*models.Shop, error) {
	// Clear the old shop owners
	for _, shopOwner := range shopOwners {
		if err := r.pg.Delete(shopOwner).Error; err != nil {
			return nil, errors.Wrap(err, "failed to delete shop owner")
		}
	}

	// Add new shop owners
	if err := r.pg.Model(models.ShopOwner{}).Clauses(clause.Returning{}).Create(shopOwners).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create shop owner")
	}

	// Get the updated shop
	shop, err := r.GetShopById(shopId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shop")
	}

	return shop, nil
}

func (r *Repository) DeleteShop(id int32) (int32, error) {
	result := r.pg.Select(clause.Associations).Delete(&models.Shop{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int32(result.RowsAffected), nil
}
