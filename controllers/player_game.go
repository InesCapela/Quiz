package controllers

import (
	"Project_2021_PSRS/model"
	"Project_2021_PSRS/services"
)

func AddPlayerGame(user *model.Users, game *model.Game) {
	var exists model.Game
	services.Db.Model(&game).Association("Game").Find(&user)

	if exists.ID != 0 {
		services.Db.Model(&game).Association("Game").Append(&user)
		services.Db.Save(&game)
	}
}

func RemovePlayerGame(user *model.Users, game *model.Game) {
	var exists model.Game
	services.Db.Model(&exists).Association("Game").Find(&user)

	if exists.ID != 0 {
		services.Db.Model(&game).Association("Game").Delete(&user)
		services.Db.Save(&game)
	}
}