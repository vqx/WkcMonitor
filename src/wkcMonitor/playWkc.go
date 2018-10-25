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

func GetRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, "https://wss.playwkc.com/sell/WKC", nil)
	if err != nil {
		fmt.Println("h2aau5jxw9")
		return nil
	}
	return req
}

func GetResponse(req *http.Request) *http.Response {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ut2cj5wmas")
		return nil
	}
	return resp
}

func GetPrice(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("n75u53w2a4")
		return -1
	}
	resp.Body.Close()
	content := string(body)
	index := strings.Index(content, "<div class=\"price\">")
	tmp := content[index + 19:index + 25]
	tmp = strings.Replace(tmp, " ", "", -1)
	price, err := strconv.ParseFloat(tmp, 10)
	if err != nil {
		fmt.Println("ebhnr5rmmt")
		return -1
	}
	return price
}

var dbObj *sql.DB
var dbOnce sync.Once

func init() {
	dbOnce.Do(func() {
		var err error
		//dbObj, err = sql.Open("mysql", "root:@/WKC?charset=utf8")
		dbObj, err = sql.Open("mysql", "root:***********@/WKC?charset=utf8")
		if err != nil {
			fmt.Println("swjcwkwmka", err)
			return
		}
		r := dbObj.QueryRow("select * from ExpectPrice where `Id` = 'now'")
		if r != nil {
			var id string
			var price float64
			r.Scan(&id, &price)
			priceLock.Lock()
			expectPrice = price
			priceLock.Unlock()
		}
		client = &http.Client{}
	})

}

func SaveToDb(price float64) {
	t := time.Now()
	_, err := dbObj.Exec("Insert into History(`Id`,`Price`) VALUES(?,?)", t.Format("2006-01-02 15:04:05"), price)
	if err != nil {
		fmt.Println(err)
		fmt.Println("7m28fkvnes")
		return
	}
}

func collectHistoryThread() {
	for time.Now().Second() % 30 != 0 {
		fmt.Println("858quu5hjs wait")
		time.Sleep(time.Second)
	}
	for {
		fmt.Println("329u5epfxr start one round")
		req := GetRequest()
		if req == nil {
			time.Sleep(time.Second * 30)
			continue
		}
		resp := GetResponse(req)
		if resp == nil {
			time.Sleep(time.Second * 30)
			continue
		}
		price := GetPrice(resp)
		if price == -1 {
			time.Sleep(time.Second * 30)
			continue
		}
		SaveToDb(price)
		if price >= expectPrice && expectPrice != 0 {
			SendEmail("it's time tp sale wkc", "wkc price is good to sale" + strconv.FormatFloat(expectPrice, 'f', 3, 64), "727998535@qq.com")
		}
		fmt.Println("329u5epfxr end one round ", price)
		time.Sleep(time.Second * 30)
	}
}