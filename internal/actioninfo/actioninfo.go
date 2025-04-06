package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(dataString string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		err := dp.Parse(value)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}