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

func WebRun() {
	http.HandleFunc("/hirstory", History)
	http.HandleFunc("/setExpectPrice", SetExpectPrice)
	go collectHistoryThread()
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	SendEmail("wkc web stoped", "wkc web stoped", "727998535@qq.com")
}

var expectPrice float64
var priceLock sync.Mutex

func insertExpectPrice(price float64) {
	_, err := dbObj.Exec("Insert into ExpectPrice(`Id`,`Price`) VALUES(?,?)", "now", price)
	if err != nil {
		fmt.Println("wamg73wv72", err.Error())
		return
	}
}

func SetExpectPrice(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	price := r.Form.Get("price")
	tmp, err := strconv.ParseFloat(price, 10)
	if err != nil {
		fmt.Println("324trdmraj ", err.Error())
		w.Write([]byte("324trdmraj " + err.Error()))
		return
	}
	priceLock.Lock()
	expectPrice = tmp
	priceLock.Unlock()
	insertExpectPrice(tmp)
	w.Write([]byte("h2z5qj4pcd success"))
}

func History(w http.ResponseWriter, r *http.Request) {
	startTime := time.Date(2018, 10, 23, 0, 0, 0, 0, time.Local)
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
	w.Write(data)
}

func GetData(startTime time.Time, endTime time.Time) []byte {
	rows, err := dbObj.Query("select * from History where `Id` >= ? and `Id` <= ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("wysujxdywu", err)
		return []byte{}
	}
	type tmp struct {
		Time  string
		Price float64
	}
	result := []tmp{}
	for rows.Next() {
		var id string
		var price float64
		rows.Scan(&id, &price)
		if price != 0 {
			result = append(result, tmp{id, price})
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
	AuthCode = "00000000000000"
	Host = "smtp.163.com:25"
)

func SendEmail(title string, content string, to string) {
	auth := smtp.PlainAuth("", DefaultMail, AuthCode, strings.Split(Host, ":")[0])
	msg := []byte("To: " + to + "\r\nFrom: " + AuthCode + "<" + AuthCode + ">\r\nSubject: " + title + "\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n" + content)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(Host, auth, DefaultMail, send_to, msg)
	if err != nil {
		fmt.Println("rbxerx2hv4 " + err.Error())
	}
}