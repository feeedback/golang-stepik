package main // https://stepik.org/lesson/351892/step/13

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

func main() {
	zipFile, _ := zip.OpenReader("task.zip")
	defer zipFile.Close()

	for _, file := range zipFile.File {
		readFile, _ := file.Open()

		bytes, _ := io.ReadAll(readFile)
		str := string(bytes)

		if len(str) > 10 && strings.Contains(str, ",") {
			rows := strings.Split(str, "\n")
			fmt.Println(strings.Split(rows[4], ",")[2])
		}
		readFile.Close()
		// if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
		// 	fmt.Println(rows[4][2])
		// }
	}
}
