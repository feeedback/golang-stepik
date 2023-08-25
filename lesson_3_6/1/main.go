package main // https://stepik.org/lesson/353243/step/6

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int // required
}
type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student // required
}

func main() {
	group := new(Group)

	bytes, _ := io.ReadAll(os.Stdin)
	_ = json.Unmarshal(bytes, group)

	totalStudents := len(group.Students)
	totalRatingsCount := 0

	for _, student := range group.Students {
		totalRatingsCount += len(student.Rating)
	}

	average := float64(totalRatingsCount) / float64(totalStudents)

	result := map[string]float64{"Average": average}

	output, _ := json.MarshalIndent(result, "", "    ")

	fmt.Println(string(output))
}
