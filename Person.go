package main

import "fmt"

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile number"`
	Address string `json:"address"`
	DOB     string `json:"DOB"`
}

var people = make(map[int]Person)
var currentId = 1

func CreatePerson(p Person) Person {
	p.ID = currentId
	people[currentId] = p
	currentId++
	return p
}

func GetPerson(id int) Person {
	person, exists := people[id]
	if !exists {
		fmt.Printf("person with id %T does not exist", id)
	}
	return person
}

func UpdatePerson(id int, update Person) (Person, bool) {
	if person, exists := people[id]; exists {
		person.Name = update.Name
		person.DOB = update.DOB
		person.Mobile = update.Mobile
		person.Address = update.Address
		people[id] = person
		return person, true
	}
	return Person{}, false
}

// DeletePerson deletes a person from the in-memory map.
func DeletePerson(id int) bool {
	if _, exists := people[id]; exists {
		delete(people, id)
		return true
	}
	return false
}

func GetAllPeople() []Person {
	personList := []Person{}
	for _, person := range people {
		personList = append(personList, person)
	}
	return personList
}
