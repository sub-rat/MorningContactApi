package contact

import "github.com/sub-rat/MorningContactApi/internals/models"

type ServiceInterface interface {
	Query(offset, limit int, q string, id uint) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(req *Contact) (Contact, error)
	Update(id uint, update *Contact) (Contact, error)
	Delete(id uint) error
}

type service struct {
	repo RepositoryInterface
}

type Contact struct {
	models.Contact
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return &service{repo}
}

func (service *service) Query(offset, limit int, q string, id uint) ([]Contact, error) {
	dataList, err := service.repo.Query(offset, limit, q, id)
	return dataList, err
}
func (service *service) Get(id uint) (Contact, error) {
	return service.repo.Get(id)
}
func (service *service) Create(req *Contact) (Contact, error) {
	err := service.repo.Create(req)
	return *req, err
}
func (service *service) Update(id uint, update *Contact) (Contact, error) {
	err := service.repo.Update(id, update)
	return *update, err
}
func (service *service) Delete(id uint) error {
	err := service.repo.Delete(id)
	return err

}
