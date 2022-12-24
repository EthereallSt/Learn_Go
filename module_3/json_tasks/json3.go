package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type dataStruct struct {
	GlobalId int64 `json:"global_id"`
}

type rawData struct {
	units []dataStruct
}

var r rawData

func main() {
	//Answer := dataStruct{}
	url := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	var netClient = http.Client{}
	res, err := netClient.Get(url)
	if err != nil {
		log.Println("ошбика подключения")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &r.units)
	if jsonErr != nil {
		log.Println(jsonErr, "ошибка в декодировании")
	}
	sum := 0
	for i := range r.units {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}
