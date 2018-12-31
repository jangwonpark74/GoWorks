package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	speciality string
}

func main() {

	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	jane.Skills = []string{"anatomy"}
	jane.Skills = append(jane.Skills, "Physics", "golang")
	jane.int = 3

	fmt.Println(jane)
}
