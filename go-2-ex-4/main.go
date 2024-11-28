package main

import "fmt"

func main() {

	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}

	student1 := Student{"Eneas", "Infanger"}
	student2 := Student{"Aaaron", "Ettlin"}
	student3 := Student{"Tim", "Hüsler"}

	student4 := Student{"Arun", "Änishänslin"}
	student5 := Student{"Leonnn", "BCX"}
	student6 := Student{"Aaron", "Hüsler"}

	classA := Class{
		Name:     "Class A",
		Students: []Student{student1, student2, student3},
	}

	classB := Class{
		Name:     "Class B",
		Students: []Student{student4, student5, student6},
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		104: {classA, classB},
		117: {classA},
		346: {classB},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println("Klassenverwaltung:")

	fmt.Println("Klassen:")
	fmt.Printf("%s: %+v\n", classA.Name, classA.Students)
	fmt.Printf("%s: %+v\n", classB.Name, classB.Students)

	fmt.Println("\nModule:")
	for module, classes := range modules {
		fmt.Printf("Modul %d wird besucht von: ", module)
		for _, class := range classes {
			fmt.Printf("%s ", class.Name)
		}
		fmt.Println()
	}
}
