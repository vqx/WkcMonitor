package wkcMonitor

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"strconv"
	"sync"
	"net/smtp"
	"strings"
	"encoding/json"
)

type WebContext struct {
	ExpectBuyPrice     float64
	ExpectSellPrice    float64
	MysqlPwd           string
	EmailCode          string
	IsMonitorBuyPrice  bool
	IsMonitorSellPrice bool
	lock               sync.Mutex
}

var webContext *WebContext

func WebRun(web *WebContext) {
	webContext = web
	initAll()
	http.HandleFunc("/history", History)
	http.HandleFunc("/getLastData", GetLastData)
	http.HandleFunc("/setExpectBuyPrice", SetExpectBuyPrice)
	http.HandleFunc("/setExpectSellPrice", SetExpectSellPrice)
	go collectHistoryThread()
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	SendEmail("wkc web stoped", "wkc web stoped", "727998535@qq.com")
}

func insertExpectPrice(id string, price float64) {
	_, err := dbObj.Exec("REPLACE into ExpectPrice(`Id`,`Price`) VALUES(?, " + strconv.FormatFloat(price, 'f', 3, 64) + ")", id)
	if err != nil {
		fmt.Println("wamg73wv72", err.Error())
		return
	}
}

func getExpectPrice(id string) float64 {
	r := dbObj.QueryRow("select * from ExpectPrice where `Id` = ?", id)
	if r != nil {
		var id1 string
		var price float64
		r.Scan(&id1, &price)
		return price
	}
	return 0
}

func initExpectPrice() {
	webContext.lock.Lock()
	webContext.ExpectBuyPrice = getExpectPrice("buy")
	webContext.ExpectSellPrice = getExpectPrice("sell")
	if webContext.ExpectBuyPrice == 0 {
		webContext.ExpectBuyPrice = 1
		insertExpectPrice("buy", 1.1)
	}
	if webContext.ExpectSellPrice == 0 {
		webContext.ExpectSellPrice = 1
		insertExpectPrice("sell", 1.1)
	}
	webContext.lock.Unlock()
}

func SetExpectBuyPrice(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	price := r.Form.Get("price")
	tmp, err := strconv.ParseFloat(price, 10)
	if err != nil {
		fmt.Println("324trdmraj ", err.Error())
		w.Write([]byte("324trdmraj " + err.Error()))
		return
	}
	webContext.lock.Lock()
	webContext.ExpectBuyPrice = tmp
	webContext.lock.Unlock()
	insertExpectPrice("buy", tmp)
	w.Write([]byte("h2z5qj4pcd success"))
}

func SetExpectSellPrice(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	price := r.Form.Get("price")
	tmp, err := strconv.ParseFloat(price, 10)
	if err != nil {
		fmt.Println("324trdmraj ", err.Error())
		w.Write([]byte("324trdmraj " + err.Error()))
		return
	}
	webContext.lock.Lock()
	webContext.ExpectSellPrice = tmp
	webContext.lock.Unlock()
	insertExpectPrice("sell", tmp)
	w.Write([]byte("h2z5qj4pcd success"))
}

func History(w http.ResponseWriter, r *http.Request) {
	startTime := time.Date(2018, 10, 28, 0, 0, 0, 0, time.Local)
	endTime := time.Now().In(time.Local)
	r.ParseForm()  //解析参数，默认是不会解析的
	startTimeStr := r.Form.Get("startTime")
	endTimeStr := r.Form.Get("endTime")
	if startTimeStr != "" {
		tmp, err := time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err == nil {
			startTime = tmp.In(time.Local)
		}
	}
	if endTimeStr != "" {
		tmp, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err == nil {
			endTime = tmp.In(time.Local)
		}
	}
	data := GetData(startTime, endTime)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func GetLastData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	row := dbObj.QueryRow("select * from History order by id DESC limit 1")
	if row == nil {
		fmt.Println("e7rdxz22br row == nil")
		w.Write([]byte("e7rdxz22br nodata"))
		return
	}
	var id string
	var sellPrice float64
	var buyPrice float64
	row.Scan(&id, &sellPrice, &buyPrice)
	byteSlice, err := json.Marshal(ResultItem{id, buyPrice, sellPrice})
	if err != nil {
		fmt.Println("bb8b46z2mt json.Marshal error, ", err.Error())
		w.Write([]byte("bb8b46z2mt json.Marshal error"))
		return
	}
	w.Write(byteSlice)
}

func GetData(startTime time.Time, endTime time.Time) []byte {
	rows, err := dbObj.Query("select * from History where `Id` >= ? and `Id` <= ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("wysujxdywu", err)
		return []byte{}
	}
	result := []ResultItem{}
	for rows.Next() {
		var id string
		var sellPrice float64
		var buyPrice float64
		rows.Scan(&id, &sellPrice, &buyPrice)
		if sellPrice != 0 && buyPrice != 0 {
			result = append(result, ResultItem{id, buyPrice, sellPrice})
		}
	}
	r, err := json.Marshal(result)
	if err != nil {
		fmt.Println("28cdksup6p", err)
		return []byte{}
	}
	return r
}

const (
	DefaultMail = "xpoony@163.com"
	Host = "smtp.163.com:25"
)

func SendEmail(title string, content string, to string) {
	body := `
    <html>
    <body>
    <pre>
    ` + content + `
    </pre>
    </body>
    </html>
    `
	//执行逻辑函数
	err := sendMail(DefaultMail, webContext.EmailCode, Host, to, title, body, "html")
	if err != nil {
		fmt.Println("发送邮件失败!")
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!")
	}
}


//发送邮件的逻辑函数
func sendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

