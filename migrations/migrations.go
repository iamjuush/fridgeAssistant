package migrations

import (
	"fridgeAssistant/database"
	"fridgeAssistant/models"
)

func Migrate() {
	database.DBCon.AutoMigrate(models.Grocery{}, models.Fridge{})

	// Set up foreign keys. More info on delete and update actions: https://www.sqlite.org/foreignkeys.html
	// This method is bugged currently and does not work: https://github.com/jinzhu/gorm/issues/635
	//database.DBCon.Model(&models.Grocery{}).AddForeignKey("FridgeID", "Fridge(ID)", "CASCADE", "RESTRICT")
}
