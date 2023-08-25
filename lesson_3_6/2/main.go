package main // https://stepik.org/lesson/353243/step/9

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StructOKVED []struct {
	ID int `json:"global_id"`
}

func main() {
	url := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при загрузке JSON:", err)
		return
	}
	defer resp.Body.Close()

	var data StructOKVED

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	sum := 0
	for _, item := range data {
		sum += item.ID
	}

	fmt.Println("Сумма полей global_id:", sum)
}
