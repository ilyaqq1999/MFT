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

func (c App) Index(results []models.Result) revel.Result {
	//fmt.Printf("GET.Массив названий магазинов в начале: %+v\n",results)

	/*
	if len(results) > 0 { //пока не уверен, что это то, что нужно
		fmt.Printf("GET.results > 0:")
		return c.Render(results)
	}*/

	curentLocale:=c.Request.Locale
	fmt.Printf("Текущая локаль get: %s\n",curentLocale)
	cl:=curentLocale[:2]
	dsn := "host=localhost user=selectel password=selectel dbname=selectel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	srtforsql := "select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact, email from shop where blocked='false' and name::json->>'"+cl+"' !='' limit 10"
	if c.Request.Method=="POST" {
		counttoshow:= c.Params.Get("datatable_length")
		searching:=c.Params.Get("search")
		fmt.Printf("Сколько показывать:%s, слово для поиска: %s\n",counttoshow,searching)
		srtforsql="select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact," +
			" email from shop where blocked='false' and lower(name::json->>'"+cl+"') like lower('%"+searching+"%') or lower(address::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(phone::json->>'"+cl+"') like lower('%"+searching+"%') or lower(contact_name::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(email) like lower('%"+searching+"%') limit "+counttoshow+""
		db.Raw(srtforsql).Scan(&results)
		//fmt.Printf("GET.Массив названий магазинов перед рендером: %+v",results)
		return c.Render(results,cl)
	}
	db.Raw(srtforsql).Scan(&results)
	//fmt.Printf("GET.Массив названий магазинов перед рендером: %+v",results)
	return c.Render(results,cl)
}

/*
func (c App) IndexPost() revel.Result {
	//Как-то спарсить всё это засунуть и переделать запрос в БД достать тут сколько показывать, что искать и какая страница
	counttoshow:= c.Params.Get("datatable_length")
	searching:=c.Params.Get("search")
	fmt.Printf("Сколько показывать:%s, слово для поиска: %s\n",counttoshow,searching)
	//var counttoshow = "25" // сколько показывать на страничке
	//var searching = "test" // слово для поиска
	curentLocale:=c.Request.Locale
	fmt.Printf("Текущая локаль post: %s\n",curentLocale)
	cl:=curentLocale[:2]

	dsn := "host=localhost user=selectel password=selectel dbname=selectel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var results []models.Result

	//errsql:= db.Table("shop").Debug().Raw("select name::json->>'en' as name from shop;").Scan(&results).Error
	//if errsql != nil{
	//	fmt.Printf("Error: %s",errsql)
	//}
	//fmt.Printf("POST.Массив названий магазинов: %+v",results)

	srtforsql:="select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact," +
		" email from shop where blocked='false' and lower(name::json->>'"+cl+"') like lower('%"+searching+"%') or lower(address::json->>'"+cl+"') like lower('%"+searching+"%')" +
		" or lower(phone::json->>'"+cl+"') like lower('%"+searching+"%') or lower(contact_name::json->>'"+cl+"') like lower('%"+searching+"%')" +
		" or lower(email) like lower('%"+searching+"%') limit "+counttoshow+""

	db.Raw(srtforsql).Scan(&results)
	fmt.Printf("POST.Массив названий магазинов: %+v",results)
	//return c.Redirect(routes.App.Index(results))
	return c.Redirect((*App).Index,results)
}

*/