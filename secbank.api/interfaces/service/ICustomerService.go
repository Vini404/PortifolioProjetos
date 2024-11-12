package interfaces

import (
	"mime/multipart"
	dto "secbank.api/dto/customer"
	"secbank.api/models"
)

type ICustomerService interface {
	S_List() (*[]models.Customer, error)
	S_Create(customer models.Customer, file multipart.File) error
	S_Delete(id int) error
	S_Update(customer models.Customer) error
	S_Get(id int) (*models.Customer, error)
	S_Auth(request dto.AuthRequest) (*dto.AuthResponse, error)
}
