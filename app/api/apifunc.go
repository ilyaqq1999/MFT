package api

import (
	"MyFTask/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func Processing(cl, counttoshow, searching, orderby,method string, activepage int) ([]models.Result, int, int, []int, int) {

	//fmt.Printf("GET.Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s, локаль:%s\n",counttoshow,searching,orderby,cl)

	/*fmt.Printf("Получил в API. Локаль: %s, Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s, method %s, активная страничка %d\n",cl, counttoshow, searching, orderby, method, activepage)*/

	dsn := "host=localhost user=selectel password=selectel dbname=selectel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	srtforsql := "select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact, email from shop where blocked='false' and length(name::json->>'"+cl+"')>0 order by name asc"

	var results []models.Result
	if method=="POST" {
		//fmt.Printf("POST.Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s, локаль:%s\n",counttoshow,searching,orderby,cl)
		srtforsql="select name::json->>'"+cl+"' as name, address::json->>'"+cl+"' as address, phone, contact_name::json->>'"+cl+"' as contact," +
			" email from shop where blocked='false'and length(name::json->>'"+cl+"')>0 and( lower(name::json->>'"+cl+"') like lower('%"+searching+"%') or lower(address::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(phone) like lower('%"+searching+"%') or lower(contact_name::json->>'"+cl+"') like lower('%"+searching+"%')" +
			" or lower(email) like lower('%"+searching+"%')) "+orderby
	}

	counttoshowint,err:= strconv.Atoi(counttoshow)//конвертируем в int для получения offset
	if err !=nil{
		counttoshowint=10
	}

	var count int//считаю количество записей по select в БД до оффсета и лимита
	db.Raw("select count(name) from ("+srtforsql+") as count").Scan(&count)

	pages:=1//количество страниц

	for i:=count-counttoshowint; i>0;i-=counttoshowint {//количество страниц
		pages++
	}

	pagesarr:= make([]int,pages)//массив страниц, чтобы в шаблоне отобразить через range
	for i:=0;i<pages;i++ {
		pagesarr[i]=i+1
	}

	if pages<activepage {//если активная страничка была больше, чем количество страниц после запроса
		activepage=1
	}

	offset:=strconv.Itoa((activepage-1)*counttoshowint)//оффсет

	/*	fmt.Printf("Активная страничка: %d, сколько показывать: %d, offset %s, количество записей в БД %d, количество страниц:%d\n",activepageint,counttoshowint,offset,count,pages)
		fmt.Printf("Массив страничек:%v\n",pagesarr)*/

	srtforsql+=" offset "+offset+" limit "+counttoshow//добавляю к запросу оффсет и лимит

	db.Raw(srtforsql).Scan(&results)

/*	fmt.Printf("После обработки в API. Локаль: %s, Сколько показывать:%s, слово для поиска: %s, сортировка по имени: %s, method %s, активная страничка %d\n",cl, counttoshow, searching, orderby, method, activepage)
	fmt.Printf("После обработки в API.сколько показывать инт %d, активная страничка инт %d, массив страничек:%v, сколько страниц %d", counttoshowint, activepage, pagesarr, pages)*/
	return results, counttoshowint, activepage, pagesarr, pages
}