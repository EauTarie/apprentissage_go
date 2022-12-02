// package main

// import (
// 	"errors"
// 	"fmt"
// ) // Comment
// /* Comment */



// func sayHello(name string) {
// 	fmt.Printf("Bonjour à tous,ici l'Apave, il va falloir être plus réactif %v", name)
// }

// func calculateOhm(u, i int) {
// 	z:= u/i;
// 	fmt.Printf("Pour calculer la résistance d'un circuit, c'est U/I, donc pour 230V et 10 A, le résultat est %v Ohm", z)
// }

// func sayHelloPrime(name, catchPhrase string) {
// 	fmt.Printf("\nBonjour, je suis %v, et je suis %v", name, catchPhrase)
// }

// func grosRelou(name, chieur string) {
// 	fmt.Printf("\nBonjour ! Je m'appelle %v, et %v", name, chieur)
// }


// func calculatePerimRect(lo,la float64) float64 {
// 	return 2* (lo + la)
// }


// func sayBye(name string) (string, error) {
// 	if name == "" {
// 		return "", errors.New("\bErreur: Vous avez oublié de spécifier un NOM ! ")
// 	}
// 	byeMessage := // It's printing the string and the variable name.
// 	fmt.Sprintf("%v s'en vas ! Bonne soirée ! ", name)
// 	return byeMessage, nil
// }

// func main() {

// 	sayHello("Brice \n")
// 	calculateOhm(230,10)
// 	sayHelloPrime("Valentin","tué de mort !")
// 	grosRelou("Alexis", "je suis un gros relou !")

// 	rl := calculatePerimRect(5.5, 2.4)
// 	fmt.Printf("\nLe périmètre de mon rectangle est de : %v. \n", rl)
// 	fmt.Println(sayBye(""))
// }