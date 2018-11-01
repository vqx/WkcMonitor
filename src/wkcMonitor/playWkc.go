package wkcMonitor

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var client *http.Client

const TimeFormat = "2006-01-02 15:04:05"

type ResultItem struct {
	Time      string
	BuyPrice  float64
	SellPrice float64
}

func (ri *ResultItem) SaveToDb() {
	_, err := dbObj.Exec("Insert into History(`Id`,`SellPrice`,`BuyPrice`) VALUES(?,?,?)", ri.Time, ri.SellPrice, ri.BuyPrice)
	if err != nil {
		fmt.Println("7m28fkvnes ", err)
		return
	}
	fmt.Println("save  price item", ri.Time, " buy: ", ri.BuyPrice, " sell: ", ri.SellPrice)
}

func (ri *ResultItem) CheckEmail() {
	if webContext.IsMonitorSellPrice {
		if webContext.ExpectSellPrice != 0 && ri.SellPrice >= webContext.ExpectSellPrice {
			SendEmail("ad sell wkc price was good", "ad sell wkc price was good" + strconv.FormatFloat(ri.SellPrice, 'f', 3, 64), "727998535@qq.com")
		}
	}
	if webContext.IsMonitorBuyPrice {
		if webContext.ExpectBuyPrice != 0 && ri.BuyPrice >= webContext.ExpectBuyPrice {
			SendEmail("ad buy wkc price was good", "ad buy wkc price was good" + strconv.FormatFloat(ri.SellPrice, 'f', 3, 64), "727998535@qq.com")
		}
	}
}

func GetRequest() *http.Request {
	url := "https://wss.playwkc.com"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("h2aau5jxw9")
		return nil
	}
	return req
}

func GetResponse(req *http.Request) *http.Response {
	client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ut2cj5wmas", err)
		return nil
	}
	return resp
}

func GetPriceItem(resp *http.Response) *ResultItem {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("n75u53w2a4")
		return nil
	}
	resp.Body.Close()
	content := string(body)
	tmp := content[strings.Index(content, "Price:") + 7:]
	buyStr := tmp[:strings.Index(tmp, " ")]
	tmp1 := tmp[strings.Index(tmp, "Sell Wankecoin"):strings.Index(tmp, "Sell Wankecoin") + 50]
	tmp2 := tmp1[strings.Index(tmp1, "Price:") + 7:]
	sellStr := tmp2[:strings.Index(tmp2, " ")]
	buy, err := strconv.ParseFloat(buyStr, 64)
	if err != nil {
		fmt.Println("rchnn9g9pt ParseFloat buy error")
		return nil
	}
	sell, err := strconv.ParseFloat(sellStr, 64)
	if err != nil {
		fmt.Println("g5f25djf9x ParseFloat sell error")
		return nil
	}
	return &ResultItem{
		Time:time.Now().Format(TimeFormat),
		SellPrice:sell,
		BuyPrice:buy,
	}
}

var dbObj *sql.DB
var dbOnce sync.Once

func initAll() {
	dbOnce.Do(func() {
		var err error
		dbObj, err = sql.Open("mysql", "root:" + webContext.MysqlPwd + "@/WKC?charset=utf8")
		if err != nil {
			fmt.Println("swjcwkwmka", err)
			return
		}
		client = &http.Client{}
		initExpectPrice()
	})
}

func collectHistoryThread() {
	for time.Now().Second() % 30 != 0 {
		fmt.Println("858quu5hjs wait")
		time.Sleep(time.Second)
	}
	for {
		collectHistory()
		time.Sleep(time.Second * 30)
	}
}

func collectHistory() {
	req := GetRequest()
	if req == nil {
		return
	}
	resp := GetResponse(req)
	if resp == nil {
		return
	}
	priceItem := GetPriceItem(resp)
	if priceItem == nil {
		return
	}
	priceItem.SaveToDb()
	priceItem.CheckEmail()
}
