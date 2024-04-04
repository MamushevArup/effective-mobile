package models

import "github.com/google/uuid"

type Car struct {
	RegNum string `json:"regNums" db:"registration_number"`
	Mark   string `json:"mark" db:"mark"`
	Model  string `json:"model" db:"model"`
	Owner  Owner
}

type RegNumsReq struct {
	RegNums []string `json:"reg_nums"`
}

type Owner struct {
	Id         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Surname    string    `json:"surname" db:"surname"`
	Patronymic string    `json:"patronymic" db:"patronymic"`
}
