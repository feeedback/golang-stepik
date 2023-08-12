package main // https://stepik.org/lesson/345543/step/6

func main() {
	// mock data
	groupCity := map[int][]string{
		10:   {"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  {"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: {"Миллионик"},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Село":      100,
		"Миллионик": 500,
	}

	// code for stepik
	for _, city := range groupCity[1000] {
		delete(cityPopulation, city)
	}
	for _, city := range groupCity[10] {
		delete(cityPopulation, city)
	}
}
