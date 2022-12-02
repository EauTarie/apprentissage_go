// package main

// import "fmt"

// func main(){
// 	supermarketPrice := map[string]int {
// 		"C": 8,
// 		"eau": 2,
// 		"2": 5,
// 	}

// 	fmt.Println(supermarketPrice["C"])

// 	for key, value := range supermarketPrice {
// 		fmt.Println(key, value)
// 	}


// 	daysInAYear := 0

// 	monthDays := map[string]int {
// 		"Janvier": 31 ,
// 		"Février": 28 ,
// 		"Mars": 31 ,
// 		"Avril": 30 ,
// 		"Mai": 31 ,
// 		"Juin": 30 ,
// 		"Juillet": 31 ,
// 		"Août": 31 ,
// 		"Septembre": 30 ,
// 		"Octobre": 31 ,
// 		"Novembre": 30 ,
// 		"Décembre": 31 ,
// 	}

// 	fmt.Println(monthDays["Janvier"])

// 	for key, value := range monthDays {
// 		fmt.Println(key,value)
// 	}

// 	for _, value := range monthDays {
// 		daysInAYear = daysInAYear + value
// 	}
// 	fmt.Printf("le nombre de jour dans une année est de %d jours \n", daysInAYear)
// }


package main

func main() {
	w := func() {
		println("BONJOUR, je suis une fonction anonyme sans return")
	}

	w()

	z := func()  string {
		println("Je suis une fonction anonyme AVEC return")
		return "Alex"
	}()
	
	println(z)


	x, y := func() (int, int) {
		println("Aucun paramètre. Mais DEUX RETURNS")
		return 5, 7
	}()

	println(x,y)


	func(a, b int) {
		println("A*A + B*B = ", a*a+b*b)
	}(x, y)
}