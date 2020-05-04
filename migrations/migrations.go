package migrations

import (
	"fridgeAssistant/database"
	"fridgeAssistant/models"

)

func Migrate() {
	database.DBCon.AutoMigrate(models.Grocery{}, models.Fridge{})

	// Set up foreign keys.
	database.DBCon.Model(&models.Grocery{}).AddForeignKey("fridge_id", "fridges(ID)", "CASCADE", "CASCADE")
}
