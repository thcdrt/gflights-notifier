package main

import "fmt"

var dates = []string{"2024-06-05", "2024-06-06"}

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

// azti dgxf dndb mkcg

func main() {
	availabilities, err := getAvailabilities()
	if err != nil {
		panic(err)
	}

	if len(availabilities) > 0 {
		if err := sendAvailabilitiesByMail(availabilities); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No availabilities match requirements")
	}
}
