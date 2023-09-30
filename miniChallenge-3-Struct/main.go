package main

import (
	"fmt"
	"os"
)

type Person struct {
	ID       string
	Name     string
	Adddress string
	Job      string
	Reason   string
}

var peoples = []Person{
	{"1", "Aqmarina", "Jakarta Selatan", "Junior-Staff", "Learn Fundamental"},
	{"2", "Fanny", "Jakarta Pusat", "Software-Engineer", "Improve Skill"},
	{"3", "FanFan", "Samarinda", "Student", "Learn Something new"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please input a name or index! example : go run main.go Fanny")
		return
	}

	targetName := os.Args[1]

	found := FoundPerson(targetName)

	if found != nil {
		fmt.Println("Information Found:")
		fmt.Println("Name:", found.Name)
		fmt.Println("Address:", found.Adddress)
		fmt.Println("Job:", found.Job)
		fmt.Println("Why choose Golang :", found.Reason)
	} else {
		fmt.Println("Information for : ", targetName, "Not Found.")
	}

}

func FoundPerson(targetName string) (result *Person) {
	var foundPerson *Person
	for _, person := range peoples {
		if person.Name == targetName {
			foundPerson = &person
			break
		} else if person.ID == targetName {
			foundPerson = &person
			break
		} else if person.Job == targetName {
			foundPerson = &person
		}
	}
	return foundPerson
}
