package main

import "fmt"

func main() {
	fmt.Println("Hello")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// there is no while in the golang
	x := 10
	for x < 10 {
		x++
		fmt.Println(x)
	}
	// ASCII REPRESENT UINT8 BYTE

	// IN CODING WE IN GOLANG IF THE SINGLE CHAR WHICH REPRESENTED BY THE SINGLE QUOTE ARE IN THE ASCII AND IF THE STRING REPRESEND BY THE DOUBLE QUOTE WHICH ARE IN THE UTF-8
	STR := "hello"
	fmt.Println(STR[0])
	// 104

	fmt.Println(string(STR[0]))
	// h
	// slice of bytes byte=uint8 which is byte ascii code for the character h
	fmt.Printf("%T", STR[0])
	// uint8
	// ascii basic keys on the keyboard
	// 256 represented
	// utf8 4 bytes 32 bits >> ascii
	// 2**32>> 2**8
	// 60 miilion more char in the utf8 which comes after the ascii as 256 char are leess in uint8 oly the positive one
	// utf8 is the superset of ascii which are the keys in the keyboard
	// single char rep the ascii
	// special char  encoding with the multiple bytes using utf8
	// utf8 is everywhere in coding
	//  single for ascii and string for the utf8
	fmt.Println("")
	for i := 0; i < len(STR); i++ {
		fmt.Printf("%c\n", STR[i])
	}
	for i, char := range STR {
		fmt.Println(i, char)

		fmt.Println(i, string(char))
	}
	for _, char := range STR {
		fmt.Printf("%T", char)
		// int32
		// strings are the byte slice or byte array
		// int32 synonymous call to the rume
		// represented the utf8 charcters
		// ascii uses the one byte and utf8 uses the 4 bytes
		// we get the int32 because for neample if the special character in the string and we loop throught the string and
		// throught the each indivdual character index i wouldnt get the whole character rather get the all of the bytes that represent the character
		// he---llo
		//- rather give the single individual
		// --- combination of all of the byte that make the utf 8 character
		//  utf8 cannot be representd as the single byte
		//this range will us the whole character
		// 4 byte that make up the whole cahracter
		//we will get each indiviual byt no the each individual character
		//int32 is the entire character
		// we have the multiple byte to represent the special character
		// loop through string on each index we get the each individual byte not the each each individual character that why return the int32 not the int8
		// byte which is synonymous to the rume
		//because we might have the special character which represent the multiple bytes that why loop through the each index
		// int32 which is entire character
		//
		break
	}
}
