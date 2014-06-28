package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ziutek/rrd"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	step      = 300
	heartbeat = 2 * step
)

type JSONData struct {
	Devicename string `json:"devicename"`
	Dbfile     string `json:"dbfile"`
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 3, 64)
}

func (q *JSONData) FromJSON(file string) error {
	J, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var data = &q
	return json.Unmarshal(J, data)
}

func main() {
	JSONStruct := &JSONData{}
	err := JSONStruct.FromJSON("/etc/conf.json")
	if err != nil {
		panic(err)
	}

	inf, err := rrd.Info(JSONStruct.Dbfile)
	if err != nil {
		log.Fatal(err)
	}

	end := time.Unix(int64(inf["last_update"].(uint)), 0)
	start := end.Add(-288 * step * time.Second)
	fetchRes, err := rrd.Fetch(JSONStruct.Dbfile, "AVERAGE", start, end, step*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer fetchRes.FreeValues()

	num := start.Format("2006-01-02_15:04:05")
	fileName := JSONStruct.Devicename + "_" + num + ".csv"

	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)

	for i := 0; i <= 288; i++ {
		v1 := fetchRes.ValueAt(0, i) * 8 / 1024 / 1024
		v2 := fetchRes.ValueAt(1, i) * 8 / 1024 / 1024
		t1 := FloatToString(v1) + "Mbps"
		t2 := FloatToString(v2) + "Mbps"
		s := make([]string, 3)
		num := (i + 1) * step
		t22 := start.Add(time.Duration(num) * time.Second)
		t11 := (t22.String())
		s[0] = t11
		s[1] = t1
		s[2] = t2
		r2.Write(s)
		r2.Flush()
	}

	fout, err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		fmt.Println(fileName, err)
		return
	}

	fout.WriteString(buf.String())
	fmt.Printf("\n")

}
