package main

import "fmt"



type Valentin struct {
	a string
	b int
	c bool
}

func main() {
	maSuperbeVariable := Valentin {
		a: "Valentin",
		b: 24,
		c: true,
	}

	lePlusGrandDesMagiciens := Valentin {
		a:"Je suis le plus grand des Magiciens",
		b: 190,
		c: true,
	}

	bigConpyuteur := Valentin {
		a: "Où est le BIG CONPYUTEUR ?",
		b: 50,
		c:false,
	}

	jeCPoCommentNommerLaVariable := Valentin {"Alexis", 28, false}


	fmt.Println(maSuperbeVariable.a)
	fmt.Println(lePlusGrandDesMagiciens.b)
	fmt.Println(bigConpyuteur)
	fmt.Println(jeCPoCommentNommerLaVariable.a)

	myContact := newContact("Alex", 30)
	fmt.Println("Ma fonction Contact : ", myContact)
}