// package main

// import "fmt"

// func update(pointerOfNumber *int, value int) {
// 	*pointerOfNumber = value
// }

// func main() {
// 	number := 10
// 	fmt.Println(number)
// 	fmt.Println("Adresse mémoire de number : ", &number)

// 	myPointer := &number
// 	fmt.Println(myPointer)
// 	fmt.Printf("La valeur de l'adresse mémoire %v est %d.\n", myPointer, *myPointer)

// 	update(myPointer, 20)
// 	fmt.Printf("La valeur de l'adresse mémoire %v a changé -> %d", myPointer, number)
// }