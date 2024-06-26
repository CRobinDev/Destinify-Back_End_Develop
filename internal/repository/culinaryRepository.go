package repository

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/model"
	"gorm.io/gorm"
)

type ICulinaryRepository interface {
	CreateData(culinary entity.Culinary) (entity.Culinary, error)
	GetCulinaryByID(param model.CulinaryParam) (entity.Culinary, error)
	GetAllCulinary(param model.CulinaryParam) ([]entity.Culinary, error)
	SearchCulinary(param model.SearchCulinary) ([]entity.Culinary, error)
	GetCulinaryByCityID(param model.CulinaryParam) ([]entity.Culinary, error)
}

type CulinaryRepository struct {
	db *gorm.DB
}

func NewCulinaryRepository(db *gorm.DB) ICulinaryRepository {
	return &CulinaryRepository{db: db}
}

func (cr *CulinaryRepository) CreateData(culinary entity.Culinary) (entity.Culinary, error) {
	err := cr.db.Debug().Create(&culinary).Error
	if err != nil {
		return culinary, err
	}

	return culinary, nil
}

func (cr *CulinaryRepository) GetCulinaryByID(param model.CulinaryParam) (entity.Culinary, error) {
	culinary := entity.Culinary{}
	err := cr.db.Debug().Where(&param).First(&culinary).Error
	if err != nil {
		return culinary, err
	}

	return culinary, nil
}

func (cr *CulinaryRepository) GetAllCulinary(param model.CulinaryParam) ([]entity.Culinary, error) {
	var culinary []entity.Culinary
	err := cr.db.Debug().Where(&param).Find(&culinary).Error
	if err != nil {
		return culinary, err
	}

	return culinary, nil
}

func (cr *CulinaryRepository) SearchCulinary(param model.SearchCulinary) ([]entity.Culinary, error) {
	var culinary []entity.Culinary
	if err := cr.db.Where("LOWER(name) LIKE LOWER(?)", "%"+param.Name+"%").Find(&culinary).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Culinary{}, nil
		}
		return nil, err
	}
	return culinary, nil
}

func (cr *CulinaryRepository) GetCulinaryByCityID(param model.CulinaryParam) ([]entity.Culinary, error) {
	var culinary []entity.Culinary
	err := cr.db.Debug().Where(&param).Find(&culinary).Error
	if err != nil {
		return culinary, err
	}

	return culinary, nil
}