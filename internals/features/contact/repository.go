package contact

import "gorm.io/gorm"

type RepositoryInterface interface {
	Query(offset, limit int, q string, userId uint) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(req *Contact) error
	Update(id uint, update *Contact) error
	Delete(id uint) error
}

type repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) RepositoryInterface {
	return &repository{db}
}

// 100
// limit = 20
// page 1 = 1-1 =0; page 2 = 2-1=1
// page no 1 = page*limit= offset 0; page no 2; page*limit; page no 3; page*limit
// select * from contacts where name like %q% or email like %q% limit 20 offset 0
func (repository *repository) Query(offset, limit int, q string, userId uint) ([]Contact, error) {
	var dataList []Contact
	err := repository.db.Debug().Model(&Contact{}).
		Where("first_name like ? or last_name like ?", "%"+q+"%", "%"+q+"%").
		Where("user_id = ?", userId).
		Limit(limit).Offset(offset).
		Find(&dataList).Error
	return dataList, err
}

func (repository *repository) Get(id uint) (Contact, error) {
	contact := Contact{}
	err := repository.db.Debug().Model(&Contact{}).First(&contact, id).Error
	return contact, err
}

func (repository *repository) Create(req *Contact) error {
	return repository.db.Debug().Create(&req).Error
}

func (repository *repository) Update(id uint, update *Contact) error {
	err := repository.db.Debug().
		Where("id = ?", id).
		Updates(&update).Error
	return err
}

func (repository *repository) Delete(id uint) error {
	err := repository.db.Debug().Delete(&Contact{}, id).Error
	return err
}
