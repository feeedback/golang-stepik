package main // https://stepik.org/lesson/266654/step/8

// import "fmt"

type Vehicle struct {
	On    bool
	Ammo  int
	Power int
}

func (v *Vehicle) Shoot() bool {
	if !v.On {
		return false
	}
	if v.Ammo > 0 {
		v.Ammo--
		return true
	}
	return false
}
func (v *Vehicle) RideBike() bool {
	if !v.On {
		return false
	}
	if v.Power > 0 {
		v.Power--
		return true
	}
	return false
}

// func main() {
// 	testStruct := new(Vehicle)
// 	/*
// 	 * Экземпляр созданной вами структуры необходимо передать в качестве
// 	 * аргумента функции testStruct, которая выполнит проверку соблюдения
// 	 * всех условий задания/
// 	 */

// 	// }
