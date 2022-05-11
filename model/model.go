package model

type User struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	User       string     `json:"user"`
	Password   string     `json:"password"`
	IdReligion int        `json:"id_religion"`
	Religion   Religion   `json:"religion" gorm:"foreignKey:IdReligion;references:Id"`
	Language   []Language `json:"language" gorm:"foreignKey:IdUser;references:Id"`
}

func (User) TableName() string {
	return "user"
}

type Religion struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Religion) TableName() string {
	return "religion"
}

type Language struct {
	Id     int
	Name   string
	IdUser int
}

func (Language) TableName() string {
	return "language"
}

// ---custom-------------
type UserReligion struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	User         string `json:"user"`
	Password     string `json:"password"`
	IdReligion   int    `json:"id_religion"`
	ReligionName string `json:"religion_name"`
}

type UserReligionDTO struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	User     string   `json:"user"`
	Password string   `json:"password"`
	Religion Religion `json:"religion"`
}
