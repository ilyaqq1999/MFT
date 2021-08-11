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
	fmt.Printf("GET.Текущая локаль: %s\n",curentLocale)
	cl:=curentLocale[:2]
	dsn := "host=localhost user=selectel password=selectel dbname=selectel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	srtforsql := "select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact, email from shop where blocked='false' and length(name::json->>'"+cl+"')>0 order by name asc limit 10"

	var results []models.Result

	if c.Request.Method=="POST" {
		counttoshow:= c.Params.Get("datatable_length")
		searching:=c.Params.Get("search")
		orderby:=c.Params.Get("sortbyname")
		if orderby == "on" {
			orderby = "order by name desc"
		} else{
			orderby="order by name asc"
		}
		fmt.Printf("POST.Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s,локаль:%s\n",counttoshow,searching,orderby,cl)
		srtforsql="select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact," +
			" email from shop where blocked='false'and length(name::json->>'"+cl+"')>0 and( lower(name::json->>'"+cl+"') like lower('%"+searching+"%') or lower(address::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(phone) like lower('%"+searching+"%') or lower(contact_name::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(email) like lower('%"+searching+"%')) "+orderby+" limit "+counttoshow+""

		db.Raw(srtforsql).Scan(&results)
		//fmt.Printf("POST.Массив магазинов: %+v",results)

		return c.Render(results,cl,counttoshow,searching)
	}
	db.Raw(srtforsql).Scan(&results)
	//fmt.Printf("GET.Массив названий магазинов перед рендером: %+v",results)
	return c.Render(results,cl)
}