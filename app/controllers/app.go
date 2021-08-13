package controllers

import (
	"MyFTask/app/api"
	"github.com/revel/revel"
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
	curentLocale:=c.Request.Locale//локализация
	cl:=curentLocale[:2]//локализация
	if countshow:= c.Params.Get("datatable_length");countshow!=""{//количетсво для показа
		counttoshow = countshow
	} else {
		counttoshow="10"
	}
	searching=c.Params.Get("search")//слово для поиска
	orderby=c.Params.Get("sortbyname")//сортировка вверх/вниз
	method:=c.Request.Method//post or get
	orderbyon:=orderby
	if orderby == "on" {
		orderby = "order by name desc"
	} else {
		orderby="order by name asc"
	}

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

	/*fmt.Printf("\n\nДо API. Локаль %s, counttoshowstr %s , searching %s,orderby %s, method %s, activepageint %d\n",cl,counttoshow,searching,orderby,method,activepageint)*/

	results, counttoshowint, activepageint, pagesarr, pages := api.Processing(cl,counttoshow,searching,orderby,method,activepageint)//вызываю функцию обработки из API

	/*fmt.Printf("\nПосле API.counttoshowint %d , activepageint %d, pagesarr %v, pages %d\n", counttoshowint, activepageint, pagesarr, pages)*/

	return c.Render(results,counttoshowint,searching,orderbyon,activepageint,pagesarr,pages)
}