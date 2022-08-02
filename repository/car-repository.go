package repository

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	"gitlab.camelit.com/walofz/cicsupport-api/entity"
	"gorm.io/gorm"
)

type CarRepository interface {
	CreateCar(car entity.CarCreate) entity.CarCreate
	ListCarByEmpId(empId uint64) entity.CarCreate
}

type carRepository struct {
	connect *gorm.DB
}

// ListCarByEmpId implements CarRepository
func (db *carRepository) ListCarByEmpId(empId uint64) entity.CarCreate {
	var carList entity.CarCreate
	db.connect.Table("car_inform").Where("emp_id = ?", empId).Scan(&carList)
	return carList
}

// CreateCar implements CarRepository
func (db *carRepository) CreateCar(car entity.CarCreate) entity.CarCreate {
	//	 var photo []
	//var id int
	var photo = car.CarPhoto
	res := db.connect.Table("car_inform").Create(&car)
	if res != nil {
		CreateCarImage(db, int(car.Id), photo)
	}
	return car
}

func CreateCarImage(db *carRepository, id int, photo []string) {
	var z = 0
	for _, s := range photo {
		//fmt.Println(i, s)
		z += 1
		y := fmt.Sprintf("%v", z)

		var b64 = s
		dc, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			panic(err)
		}
		var new_file = strconv.FormatInt(time.Now().Unix(), 20) + y + ".jpg"
		//f, err := os.Create("http://172.16.0.29/cicsupport/backend/web/uploads/myfilename.jpg")
		//	f, err := os.OpenFile("uploads/5s/car/"+new_file, os.O_WRONLY|os.O_CREATE, 0777)//administrator@172.16.0.240/uploads
		f, err := os.OpenFile("/Volumes/uploads/"+new_file, os.O_WRONLY|os.O_CREATE, 0777) //administrator@172.16.0.240/uploads
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err := f.Write(dc); err != nil {
			panic(err)
		}

		db.connect.Table("car_inform_image").Create([]map[string]interface{}{
			{"car_inform_id": id, "photo": new_file},
		})

		// if err := f.Sync(); err != nil {
		// 	panic(err)
		// }
		// f.Seek(0, 0)

		// io.Copy(os.Stdout, f)

	}

}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{connect: db}
}
