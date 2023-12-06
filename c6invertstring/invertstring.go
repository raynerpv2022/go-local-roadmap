package main

import "fmt"

/*Crea un programa que invierta el orden de una cadena de texto sin usar funciones propias del lenguaje que lo hagan de forma automática.
 * - Si le pasamos "Hola mundo" nos retornaría "odnum aloH"*/

func invertString(s string) string {
	sInv := ""
	for i := 0; i < len(s); i++ {

		sInv = string(s[i]) + sInv
	}
	return sInv
}

func recursiveWay(s string) string {
	/*do it*/
}

func main() {

	fmt.Println(invertString("hola cabeza ueca"))

}
