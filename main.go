package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status      `json:"status"`
	StatusWater string `json:"status water"`
	StatusWind  string `json:"status wind"`
}

func updateData() {
	for {

		var data = Data{}
		waterMin := 1
		waterMax := 30

		data.Status.Water = rand.Intn(waterMax - waterMin + 1)

		data.Status.Wind = rand.Intn(waterMax - waterMin + 1)

		if data.Status.Water <= 5 {
			data.StatusWater = "Aman"
		} else if data.Status.Water <= 8 && data.Status.Water >= 6 {
			data.StatusWater = "Siaga"
		} else if data.Status.Water > 8 {
			data.StatusWater = "Bahaya"
		}

		if data.Status.Wind <= 6 {
			data.StatusWind = "Aman"
		} else if data.Status.Wind <= 15 && data.Status.Wind >= 7 {
			data.StatusWind = "Siaga"
		} else if data.Status.Wind > 15 {
			data.StatusWind = "Bahaya"
		}

		b, err := json.MarshalIndent(data, "", " ")

		if err != nil {
			log.Fatalln("error while marshalling json data  =>", err.Error())
		}

		err = ioutil.WriteFile("data.json", b, 0644)

		if err != nil {
			log.Fatalln("error while writing value to data.json file  =>", err.Error())
		}
		fmt.Println("menggungu 15 detik")
		time.Sleep(time.Second * 15)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go updateData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")

		var data = Data{}

		b, err := ioutil.ReadFile("data.json")

		if err != nil {
			fmt.Fprint(w, "error")
			return
		}

		err = json.Unmarshal(b, &data)

		err = tpl.ExecuteTemplate(w, "index.html", data)

	})

	http.ListenAndServe(":8080", nil)
}
