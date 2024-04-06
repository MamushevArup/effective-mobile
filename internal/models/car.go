package models

import "github.com/google/uuid"

type Car struct {
	RegNum string `json:"regNum,omitempty" db:"registration_number"`
	Mark   string `json:"mark,omitempty" form:"mark" db:"mark"`
	Model  string `json:"model,omitempty" form:"model" db:"model"`
	Year   int    `json:"year,omitempty" form:"year" db:"year"`
	Owner  Owner  `json:"owner,omitempty"`
}

type Pagination struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type RegNumsReq struct {
	RegNums []string `json:"reg_nums"`
}

type Owner struct {
	Id         *uuid.UUID `json:"-,omitempty" db:"id"`
	Name       string     `json:"name,omitempty" db:"name"`
	Surname    string     `json:"surname,omitempty" db:"surname"`
	Patronymic string     `json:"patronymic,omitempty" db:"patronymic"`
}
