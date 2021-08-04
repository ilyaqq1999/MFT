package controllers

import (
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
	//var shopname = db.Select("name::json->>'en' from shop limit 1")
	//db.Select("name::json->>'en' from shop")
	//var shopnames = db.Select("name::json->>'en' from shop")
	type Result struct {
		Name string
	}
	var results []Result
	db.Raw("select name from shop").Scan(&results)
	//( "SELECT id, DATA, DATA::json->>'username' AS name FROM SAMPLE  where DATA::json->>'blocked'='true' ;" );
	return c.Render(results)
}
