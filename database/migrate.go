package database

import (
	"ktfs/config"
	"ktfs/handler"
	"ktfs/model"

	"fmt"
)

func Migrate() {
	tables, err := config.DB.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	for _, table := range tables {
		config.DB.Migrator().DropTable(table)
	}

	fmt.Println("Successfully drop all table")

	config.DB.AutoMigrate(&model.User{})
	config.DB.AutoMigrate(&model.Student{})
	config.DB.AutoMigrate(&model.Subject{})

	fmt.Println("Successfully migrate database")

	users := []model.User{
		{
			Name: "Admin",
			Email: "admin@mail.com",
			Password:	handler.GeneratePassword("password"),
		},
		{
			Name: "User",
			Email: "user@mail.com",
			Password:	handler.GeneratePassword("password"),
		},
	}
	result := config.DB.Create(users)
	if result.Error != nil {
		panic(result.Error)
	} else {
		fmt.Println("Successfully seed users")
	}
}