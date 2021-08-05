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
	curentLocale:=c.Request.Locale
	fmt.Printf("Текущая локаль: %s\n",curentLocale)
	cl:=curentLocale[:2]
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

	srtforsql:="select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact, email from shop where blocked='false' and name::json->>'"+cl+"' !=''"

	db.Raw(srtforsql).Scan(&results)
	fmt.Printf("Массив названий магазинов: %+v",results)
	return c.Render(results,cl)
}