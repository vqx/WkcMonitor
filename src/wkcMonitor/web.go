package wkcMonitor

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"sync"
	"encoding/json"
)

type WebContext struct {
	MysqlPwd           string
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
	http.HandleFunc("/getIsOnline", GetIsOnline)
	http.HandleFunc("/getDataGroup", GetDataGroup)
	//go collectHistoryThread()
	fmt.Println("n7k8d8cc8m start http server in 127.0.01:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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

func GetIsOnline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("true"))
}

func GetDataGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(getQValue(r,"sub"))
	w.Write([]byte("successed"))
}

func getQValue(r *http.Request, key string) string {
	r.Form.Encode()
	return r.FormValue(key)
}