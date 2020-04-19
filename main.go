package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)


func main() {
	e := getSpreadsheet()
	_, month, day := time.Now().Date()
	msg := "Parabéns aos aniversáriantes do dia:\n"

	for _, v := range e {
		if v.Birthday != "NULL" {
			t, err := time.Parse(layoutISO, v.Birthday)
			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println("element is:", v, "date:", t.Format(layoutUS))
			if month == t.Month() && day == t.Day() {
				msg = msg + v.Name + "\n"
				fmt.Println("############### Aniversário!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			}
		}
	}

	fmt.Println("Msg:", msg)

	SlackMsg(msg)
}
