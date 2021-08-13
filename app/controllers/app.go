package controllers

import (
	"MyFTask/app/models"
	"fmt"
	"github.com/revel/revel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var activepageint =1
/*var counttoshow  =10
var searching =""
var orderby =""*/

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	curentLocale:=c.Request.Locale //локаль
	cl:=curentLocale[:2]
	counttoshow:= c.Params.Get("datatable_length")//количетсво для показа
	if counttoshow==""{
		counttoshow="10"
	}
	searching:=c.Params.Get("search")//слово для поиска
	orderby:=c.Params.Get("sortbyname")//сортировка вверх/вниз
	orderbyon:=orderby
	if orderby == "on" {
		orderby = "order by name desc"
	} else{
		orderby="order by name asc"
	}

	//fmt.Printf("GET.Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s, локаль:%s\n",counttoshow,searching,orderby,cl)

	dsn := "host=localhost user=selectel password=selectel dbname=selectel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	srtforsql := "select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact, email from shop where blocked='false' and length(name::json->>'"+cl+"')>0 order by name asc"

	var results []models.Result
	if c.Request.Method=="POST" {
		//fmt.Printf("POST.Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s, локаль:%s\n",counttoshow,searching,orderby,cl)
		srtforsql="select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact," +
			" email from shop where blocked='false'and length(name::json->>'"+cl+"')>0 and( lower(name::json->>'"+cl+"') like lower('%"+searching+"%') or lower(address::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(phone) like lower('%"+searching+"%') or lower(contact_name::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(email) like lower('%"+searching+"%')) "+orderby
	}

 	counttoshowint,err:= strconv.Atoi(counttoshow)	//конвертируем в int для получения offset
	if err !=nil{
		counttoshowint=10
	}


	var count int//считаю количество записей по select в БД до оффсета и лимита
	db.Raw("select count(name) from ("+srtforsql+") as count").Scan(&count)

	pages:=1//количество страниц

	for i:=count-counttoshowint; i>0;i-=counttoshowint {
		pages++
	}

	pagesarr:= make([]int,pages)//массив страниц, чтобы в шаблоне пробежаться через range
	for i:=0;i<pages;i++ {
		pagesarr[i]=i+1
	}

	if 	prevpage :=c.Params.Get("prev"); prevpage !="" && activepageint>1 { //если  prevpage
		activepageint--
	}
	if 	page :=c.Params.Get("page");page!= ""{//для нажатия на номера страничек
		activepageint,err=strconv.Atoi(page)
		if err != nil {
			activepageint=1
		}
	}
	if nextpage :=c.Params.Get("next");nextpage != ""  && activepageint<pages {//если  nextpage
		activepageint++
	}

	offset:=strconv.Itoa((activepageint-1)*counttoshowint)//оффсет

	fmt.Printf("Активная страничка: %d, сколько показывать: %d, offset %s, количество записей в БД %d, количество страниц:%d\n",activepageint,counttoshowint,offset,count,pages)
	fmt.Printf("Массив страничек:%v\n",pagesarr)

	srtforsql+=" offset "+offset+" limit "+counttoshow//добавляю к запросу оффсет и лимит

	db.Raw(srtforsql).Scan(&results)
	return c.Render(results,counttoshowint,searching,orderbyon,activepageint,pagesarr,pages)
}