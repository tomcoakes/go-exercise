package main

import (
	"fmt"
	"strings"
)

func main() {
	p1 := person{name: "Tom"}
	p2 := person{name: "Dan"}
	p3 := person{name: "Chris"}

	p1, p2 = makeFriends(p1, p2)
	p1, p3 = makeFriends(p1, p3)
	fmt.Println(listFriends(p1))
}

type person struct {
	name    string
	friends []string
}

func makeFriends(firstPerson person, secondPerson person) (person, person) {
	firstPerson.friends = append(firstPerson.friends, secondPerson.name)
	secondPerson.friends = append(secondPerson.friends, firstPerson.name)
	return firstPerson, secondPerson
}

func listFriends(person person) string {
	listOfFriends := strings.Join(person.friends, ", ")
	return fmt.Sprintf("%s's friends are %v", person.name, listOfFriends)
}
