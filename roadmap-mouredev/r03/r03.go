/*
 * EJERCICIO:
 * - Muestra ejemplos de creación de todas las estructuras soportadas por defecto en tu lenguaje.
 * - Utiliza operaciones de inserción, borrado, actualización y ordenación.
 *

 */

package main

import (
	"fmt"
	"sort"
)

func teoria() {
	// array
	cryptoPlus := [4]int{60000, 55, 2000, 100}
	// slice
	crypto := []string{"BITCOIN", "CRONOS", "ETHERIUM", "BINANCE", "USDT"}
	fmt.Println(cryptoPlus)
	fmt.Println(crypto)
	crypto = append(crypto, "PETRO") // insertar al final
	fmt.Println(crypto)
	crypto = append([]string{"CUBACOIN"}, crypto...) // insertar al inicio
	fmt.Println(crypto)
	crypto = append(crypto[:3], append([]string{"AURICHCOIN"}, crypto[3:]...)...) // insertar en posicion 3
	fmt.Println(crypto)
	crypto = append(crypto[:0], crypto[1:]...) // eliminar 1er item
	fmt.Println(crypto)
	crypto = append(crypto[:2], crypto[3:]...) // eliminar 3er elemento
	fmt.Println(crypto)
	crypto = append(crypto[:len(crypto)-1]) // eliminar ultimo elemento
	fmt.Println(crypto)
	crypto = crypto[:len(crypto)-1] //eliminar ultimo elemento
	fmt.Println(crypto)
	// mapas
	cryptoMap := map[string]int{"BITCOIN": 20, "CUBACOIN": 10, "AURICHCOIN": 90, "CASARABONELACOIN": 100}
	fmt.Println((cryptoMap))
	cryptoMap["ESPCOIN"] = 20 // agregar
	fmt.Println((cryptoMap))
	delete(cryptoMap, "BITCOIN") // eliminar
	fmt.Println((cryptoMap))
	cryptoMap["CUBACOIN"] = 1 // actualizar
	fmt.Println((cryptoMap))

	// busqueda de un elemento

	value, ok := cryptoMap["CUBACOIN"]
	if ok {
		fmt.Println("ENCONTRADO", value)
	} else {
		fmt.Println("NO ENCONTRADO CUBACOiN")
	}

	keys := []string{}
	for k := range cryptoMap {
		keys = append(keys, k)
	}
	fmt.Println(cryptoMap)
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	fmt.Println(cryptoMap)
}

type PhoneBook struct {
	name  string
	phone int
}

func TelephonBook() {

}
func main() {
	// teoria()

}

/*
* DIFICULTAD EXTRA (opcional):
 * Crea una agenda de contactos por terminal.
 * - Debes implementar funcionalidades de búsqueda, inserción, actualización y eliminación de contactos.
 * - Cada contacto debe tener un nombre y un número de teléfono.
 * - El programa solicita en primer lugar cuál es la operación que se quiere realizar, y a continuación
 *   los datos necesarios para llevarla a cabo.
 * - El programa no puede dejar introducir números de teléfono no númericos y con más de 11 dígitos.
 *   (o el número de dígitos que quieras)
 * - También se debe proponer una operación de finalización del programa.
*/
