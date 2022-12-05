package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ResponseAboutStudents struct {
	Students []struct {
		Rating []int `json:"Rating"`
	} `json:"Students"`
}

func main() {
	var unic ResponseAboutStudents
	var ch, summa int
	data, err := ioutil.ReadAll(os.Stdin) // здесь мы создали переменную (структуру unic) и записали то на входе
	if err != nil {
		err = json.Unmarshal(data, &unic)
		//fmt.Println((unic.Students.Rating))
	}
	for in := range unic.Students {
		in
		ch += 1
		summa += len(in)
	}
	answere := summa / ch

}
