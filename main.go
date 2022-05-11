package main

import (
	"encoding/json"
	"fmt"
	"test-gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "root:Admin123@tcp(127.0.0.1:3306)/relasi?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

var (
	DB *gorm.DB
)

func main() {
	// users := []model.UserReligion{}
	// err := DB.Raw(`SELECT user.id, user.name, user.user, user.password, user.id_religion, religion.name as religion_name
	// FROM user JOIN religion ON user.id_religion = religion.id`).Scan(&users).Error
	// // err := DB.Find(&users).Error
	// if err != nil {
	// 	panic(err)
	// }
	// result := []model.UserReligionDTO{}
	// for _, row := range users {
	// 	result = append(result, model.UserReligionDTO{
	// 	Id:       row.Id,
	// 	Name:     row.Name,
	// 	User:     row.User,
	// 	Password: row.Password,
	// 	Religion: model.Religion{
	// 		Id:   row.IdReligion,
	// 		Name: row.ReligionName,
	// 	},
	// })
	// }

	result := []model.User{}
	err := DB.Model(&result).Debug().Preload("Language").Preload("Religion").Find(&result).Error
	if err != nil {
		panic(err)
	}
	asByteJson, err := json.Marshal(result)
	fmt.Println(string(asByteJson))
}
