package models

import "github.com/google/uuid"

type Car struct {
	RegNum string `json:"regNum" db:"registration_number"`
	Mark   string `json:"mark,omitempty" db:"mark"`
	Model  string `json:"model,omitempty" db:"model"`
	Year   int    `json:"year,omitempty" db:"year"`
	Owner  Owner  `json:"owner,omitempty"`
}

type RegNumsReq struct {
	RegNums []string `json:"reg_nums"`
}

type Owner struct {
	Id         uuid.UUID `json:"-,omitempty" db:"id"`
	Name       string    `json:"name,omitempty" db:"name"`
	Surname    string    `json:"surname,omitempty" db:"surname"`
	Patronymic string    `json:"patronymic,omitempty" db:"patronymic"`
}
