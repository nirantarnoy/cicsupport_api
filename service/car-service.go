package service

import (
	"log"

	"github.com/mashingan/smapping"
	"gitlab.camelit.com/walofz/cicsupport-api/dto"
	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gitlab.camelit.com/walofz/cicsupport-api/repository"
)

type CarService interface {
	CreateCar(carDto dto.CarCreateDto) entity.CarCreate
	ListCarByEmpId(empId uint64) []entity.CarListEmp
}
type carService struct {
	carRepository repository.CarRepository
}

// ListCarByEmpId implements CarService
func (db *carService) ListCarByEmpId(empId uint64) []entity.CarListEmp {
	return db.carRepository.ListCarByEmpId(empId)
}

// CreateCar implements CarService
func (db *carService) CreateCar(carDto dto.CarCreateDto) entity.CarCreate {
	car := entity.CarCreate{}
	err := smapping.FillStruct(&car, smapping.MapFields(&carDto))
	if err != nil {
		log.Fatalf("Fail to mapping %v", err)
	}
	res := db.carRepository.CreateCar(car)
	return res
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{carRepository: repo}
}
