// 🎯 Juego propuesto: Adivina el número
// 🔹 Fase 1 — SRP (Single Responsibility Principle)

// Regla del juego:
// El sistema genera un número aleatorio del 1 al 100. El jugador tiene que adivinarlo. El sistema le dice si el número es mayor o menor.

// Responsabilidades separadas:

//     Game: Lógica general del juego (iniciar, jugar, terminar).

//     RandomNumberGenerator: Genera el número secreto.

//     PlayerInput: Lee el número ingresado por el jugador.

// Feedback: Dice si el número es mayor o menor.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type InputDataInterface interface {
	LevelGame() (int, error)
}

type NumberComparationaInterface interface {
	Comparation(ranNumber int, playerInput UserDataInterface, errorManager ErrorInterface, f FeedbackInterface) int
}

type UserDataInterface interface {
	ReadNumber() (int, error)
}

type ErrorInterface interface {
	DisplayError()
	GetError() error
	SetError(err error)
}
type RandomNumberGeneratorInterface interface {
	Generator(intervalo int) int
}

type FeedbackInterface interface {
	Max(randNumber int, userNumber int) string
	Result(text string)
}

type ErrorManager struct {
	err error
}

func (e *ErrorManager) DisplayError() {
	if e.GetError() != nil {
		fmt.Println("Hubo un error, saliendo del juego ...")
	}

}

func (e *ErrorManager) SetError(err error) {
	e.err = err
}

func (e *ErrorManager) GetError() error {
	return e.err
}

type InputData struct{}

func (lg *InputData) LevelGame() (int, error) {
	var level string
	fmt.Println("Dificltad del Juego")
	fmt.Scan(&level)
	number, err := strconv.Atoi(level) // convierte a int

	if err != nil {

		return 0, err
	}
	return number, nil

}

type Feedback struct{}

func (fb *Feedback) Result(text string) {
	fmt.Println(text)
}

func (ng *Feedback) Max(randNumber int, userNumber int) string {

	if randNumber > userNumber {
		return "Numero del Usuario Menor"
	}
	return "Numero del Usuario Mayor"

}

type NumberComparation struct{}

func (nc *NumberComparation) Comparation(ranNumber int, playerInput UserDataInterface, errorManager ErrorInterface, f FeedbackInterface) int {

	count := 0
	var usernumber int
	var err error
	var result string
	for {
		usernumber, err = playerInput.ReadNumber()

		if err != nil {
			errorManager.SetError(err)
			return 0
		}
		if ranNumber != usernumber {
			result = f.Max(ranNumber, usernumber)
		} else {
			return count
		}

		f.Result(result)
		count++

	}

}

type RandomNumberGenerator struct {
}

func (rng *RandomNumberGenerator) Generator(intervalo int) int {
	return rand.Intn(intervalo)

}

type PlayerInput struct{}

func (pi *PlayerInput) ReadNumber() (int, error) {
	var n int
	fmt.Println("Ingrese el numero: ")
	_, error := fmt.Scan(&n)
	if error != nil {
		return 0, error
	}
	return n, nil
}

type Game struct {
	randonNumber int
	dificultad   int
	counter      int
}

func (game *Game) Start(randonN RandomNumberGeneratorInterface, InputLevel InputDataInterface, errorManager ErrorInterface) {
	var err error

	game.dificultad, err = InputLevel.LevelGame()
	if err != nil {
		errorManager.SetError(err)
		return
	}

	game.randonNumber = randonN.Generator(game.dificultad)

}

func (game *Game) Play(playerInput UserDataInterface, f FeedbackInterface, nc NumberComparationaInterface, err ErrorInterface) {

	game.counter = nc.Comparation(game.randonNumber, playerInput, err, f)

}

func (game *Game) End(f FeedbackInterface) {
	f.Result(fmt.Sprintf("Has encontrado el numero %v en %v intentos", game.randonNumber, game.counter))
}

func main() {
	// ejercicio 3
	game := Game{}
	errorManager := &ErrorManager{}
	user := &PlayerInput{}
	randon := &RandomNumberGenerator{}
	feedback := &Feedback{}
	inputData := &InputData{}
	numberComparation := &NumberComparation{}

	game.Start(randon, inputData, errorManager)
	if errorManager.GetError() != nil {
		errorManager.DisplayError()
		return
	}
	game.Play(user, feedback, numberComparation, errorManager)
	if errorManager.GetError() != nil {
		errorManager.DisplayError()
		return
	} else {
		game.End(feedback)
	}

}
