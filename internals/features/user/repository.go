package user

import "gorm.io/gorm"

type RepositoryInterface interface {
	Query(offset, limit int, q string) ([]User, error)
	Get(id uint) (User, error)
	Create(req *User) error
	Update(id uint, update *User) error
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
// select * from users where name like %q% or email like %q% limit 20 offset 0
func (repository *repository) Query(offset, limit int, q string) ([]User, error) {
	var dataList []User
	err := repository.db.Debug().Model(&User{}).
		Where("name like ? or email like ?", "%"+q+"%", "%"+q+"%").
		Limit(limit).Offset(offset).
		Find(&dataList).Error
	return dataList, err
}

func (repository *repository) Get(id uint) (User, error) {
	user := User{}
	err := repository.db.Debug().Model(&User{}).First(&user, id).Error
	return user, err
}

func (repository *repository) Create(req *User) error {
	return repository.db.Debug().Create(&req).Error
}

func (repository *repository) Update(id uint, update *User) error {
	err := repository.db.Debug().
		Where("id = ?", id).
		Updates(&update).Error
	return err
}

func (repository *repository) Delete(id uint) error {
	err := repository.db.Debug().Delete(&User{}, id).Error
	return err
}
