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
	"math/rand"
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

func GetRequest() *http.Request {
	url := "https://www.playwkc.com"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("h2aau5jxw9")
		return nil
	}
	userAgentSlice := []string{
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1",
		"Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5",
		"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
	}
	ra := rand.Intn(len(userAgentSlice))
	req.Header.Set("user-agent", userAgentSlice[ra])
	return req
}

func GetResponse(req *http.Request) *http.Response {
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
	defer resp.Body.Close()
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
		client = &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: true,
			},
		}
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
}