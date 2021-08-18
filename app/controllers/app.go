package controllers

import (
	"MyFTask/app/models"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var activepageint =1
var counttoshow  ="10"
var searching =""
var orderby =""

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Redirect("/shops")
}

func (c App) GetShops() revel.Result {//getter
	curentLocale:=c.Request.Locale//локализация
	cl:=curentLocale[:2]//локализация
	if countshow:= c.Params.Get("datatable_length");countshow!=""{//количетсво для показа
		counttoshow = countshow
	} else {
		counttoshow="10"
	}
	searching=c.Params.Get("search")//слово для поиска
	orderby=c.Params.Get("sortbyname")//сортировка вверх/вниз

	//пагинация
	if 	prevpage :=c.Params.Get("prev"); prevpage !="" && activepageint>1 { //если  prevpage
		activepageint--
	}
	var err error
	if 	page :=c.Params.Get("page");page!= ""{//для нажатия на номера страничек
		activepageint,err=strconv.Atoi(page)
		if err != nil {
			activepageint=1
		}
	}
	if nextpage :=c.Params.Get("next");nextpage != ""/*  && activepageint<pages*/ {//если  nextpage
		activepageint++
	}

	httpreqforAPI:="http://localhost:9999/?locale="+cl+"&counttoshow="+counttoshow+"&search="+searching+"&orderby="+orderby+"&activepageint="+strconv.Itoa(activepageint)//строка для запроса
	res,err:=http.Get(httpreqforAPI)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", body)

	var shops models.Shops
	err=json.Unmarshal(body,&shops)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Вывожу шопс %+v\n\n\n", shops)
	//results, counttoshowint, activepageint, pagesarr, pages := Processing(cl,counttoshow,searching,orderby,method,activepageint)//вызываю функцию обработки из API

	/*fmt.Printf("\nПосле API.counttoshowint %d , activepageint %d, pagesarr %v, pages %d\n", counttoshowint, activepageint, pagesarr, pages)*/
	results:=shops.Results
	counttoshowint:=shops.Counttoshowint
	activepageint=shops.Activepageint
	pagesarr:=shops.Pagesarr
	pages:=shops.Pages
	return c.Render(results,counttoshowint,searching,orderby,activepageint,pagesarr,pages)
}