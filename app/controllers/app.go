package controllers

import (
	"MyFTask/app/models"
	"fmt"
	"github.com/revel/revel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	dsn := "host=localhost user=selectel password=selectel dbname=selectel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var results []models.Result

	//db.Table("shop").Select("name").Find(&results)

/*	errsql:= db.Table("shop").Debug().Raw("select name::json->>'en' as name from shop;").Scan(&results).Error
	if errsql != nil{
		fmt.Printf("Error: %s",errsql)
	}
	fmt.Printf("Массив названий магазинов: %+v",results)*/
	db.Raw("select name::json->>'en' as name, address::json->>'en' as address, phone, contact_name::json->>'en' as contact, email from shop where blocked='false';").Scan(&results)
	fmt.Printf("Массив названий магазинов: %+v",results)
	return c.Render(results)
}