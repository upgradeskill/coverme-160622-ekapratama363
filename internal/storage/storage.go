package storage

import (
	"main/internal/dto"
)

var User = make(map[string]dto.User)

func Seeder() {
	User["1"] = dto.User{ID: "1", Name: "Eka"}
}
