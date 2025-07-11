package main

import (
	"fmt"
	"math/rand"
)

// * EJERCICIO:
//  * Explora el "Principio SOLID de Responsabilidad Única (Single Responsibility
//  * Principle, SRP)" y crea un ejemplo simple donde se muestre su funcionamiento
//  * de forma correcta e incorrecta.
//  *
//  * DIFICULTAD EXTRA (opcional):
//  * Desarrolla un sistema de gestión para una biblioteca. El sistema necesita
//  * manejar diferentes aspectos como el registro de libros, la gestión de usuarios
//  * y el procesamiento de préstamos de libros.
//  * Requisitos:
//  * 1. Registrar libros: El sistema debe permitir agregar nuevos libros con
//  * información básica como título, autor y número de copias disponibles.
//  * 2. Registrar usuarios: El sistema debe permitir agregar nuevos usuarios con
//  * información básica como nombre, número de identificación y correo electrónico.
//  * 3. Procesar préstamos de libros: El sistema debe permitir a los usuarios
//  * tomar prestados y devolver libros.
//  * Instrucciones:
//  * 1. Diseña una clase que no cumple el SRP: Crea una clase Library que maneje
//  * los tres aspectos mencionados anteriormente (registro de libros, registro de
//  * usuarios y procesamiento de préstamos).
//  * 2. Refactoriza el código: Separa las responsabilidades en diferentes clases
//  * siguiendo el Principio de Responsabilidad Única.

type Book struct {
	title  string
	author string
	copie  int
}

// *********************************

type User struct {
	name  string
	id    int
	email string
}

// ***********************

type Loan struct {
	loans map[string][]string
}

func (l *Loan) CreateLoan(u string, b string) {
	l.loans[b] = append(l.loans[b], u)
}

func (l *Loan) DeleteLoan(b string) {
	if len((l.loans[b])) == 1 {
		delete(l.loans, b)

	} else {
		l.loans[b] = l.loans[b][1:]
	}

}

// ***********************
type Library struct {
	books []Book
	users []User
	loans Loan
}

func (l *Library) AddLoan(b *Book, u User) {
	if b.copie == 0 {
		fmt.Printf("Libro %v No hay copias disponibles para Prestar : copias %v\n", b.title, b.copie)
		return
	} else {
		b.copie -= 1
	}

	if l.loans.loans == nil {
		l.loans.loans = map[string][]string{}
	}

	l.loans.CreateLoan(u.name, b.title)
}

func (l *Library) DeleteLoan(b *Book) {

	l.loans.DeleteLoan(b.title)
	b.copie += 1
}

func (l *Library) AddUser(u User) {

	l.users = append(l.users, u)
}

func (l *Library) AddBook(b Book) {

	l.books = append(l.books, b)
}

// **********************************

type PrintService struct {
}

func (p PrintService) PrintBook(b Book) {
	fmt.Printf("Libro ** Titulo : %v Autor %v, Copias %v\n", b.title, b.author, b.copie)
}

func (p PrintService) PrintUser(u User) {
	fmt.Printf("Usuario ** Nombre : %s, Id %v, email %s\n", u.name, u.id, u.email)
}

func (p PrintService) ListUser(lu []User) {
	fmt.Println("Lista de Usuarios")
	for _, v := range lu {
		p.PrintUser(v)
	}

}

func (p PrintService) ListBook(lb []Book) {
	fmt.Println("Lista de Libros")
	for _, v := range lb {
		p.PrintBook(v)
	}
}
func (p PrintService) ListLoan(ln Loan) {
	fmt.Println("Lista de Prestamos")
	for i, v := range ln.loans {
		fmt.Printf("Libro: %v : ", i)
		for _, l := range v {
			fmt.Printf("%v, ", l)
		}
	}
	fmt.Println()
}

// ejercicio 2
// Crea un sistema que:

//     Reciba errores desde diferentes partes del sistema (como texto).
//     Guarde los errores en memoria (en una lista, por ejemplo).
//     Envíe una alerta (puede ser un print() especial) si el error es crítico.
//     Muestre los errores recientes por consola para los desarrolladores.
//  Tu objetivo:
//     Diseña las clases aplicando el Principio de Responsabilidad Única (SRP)
//     Cada clase debe tener una única razón para cambiar.

type SaverInterface interface {
	SaveError(e Error)
	GetError() []Error
}

type LogInterface interface {
	ErrorLog(e Error, a AlertInterface)
}

type AlertInterface interface {
	SendAlert(e Error)
}

type DisplayInterface interface {
	ShowError(er SaverInterface)
}

type Error struct {
	name  string
	ttype bool
}

type ErrorRepository struct {
	error []Error
}

func (s *ErrorRepository) SaveError(e Error) {
	s.error = append(s.error, e)
}

func (s *ErrorRepository) GetError() []Error {
	return s.error
}

type ErrorLogged struct{}

func (el *ErrorLogged) ErrorLog(e Error, a AlertInterface) {

	if e.ttype {

		a.SendAlert((e))
	}
	fmt.Printf("Error Registrado...\n")
}

type Alert struct {
}

func (a *Alert) SendAlert(e Error) {
	if e.ttype {
		fmt.Printf(" Critical Error : %v\n", e.name)
	}
}

type DisplayError struct {
}

func (d *DisplayError) ShowError(er SaverInterface) {
	fmt.Printf("Lista de Errores\n")
	fmt.Printf(" \tNro \tError \t\t\t\tCritical \n")
	fmt.Printf("\t--- \t--------------------------- \t ----- \n")
	for i, v := range er.GetError() {
		fmt.Printf("\t%3v\t%20v\t %4v\n", i, v.name, v.ttype)
	}
}

type Orquester struct {
}

func (o *Orquester) StartRegister(e Error, r SaverInterface, el LogInterface, a AlertInterface) {
	r.SaveError(e)
	el.ErrorLog(e, a)

}

func (o *Orquester) DisplayRegister(r SaverInterface, d DisplayInterface) {
	d.ShowError((r))
}


func main() {
	B1 := Book{"b1", "www", 3}
	U1 := User{"U1", 111, "dd@dsds"}
	B2 := Book{"B2", "www", 3}
	U2 := User{"U2", 111, "dd@dsds"}
	L := Library{}
	L.AddBook(B1)
	L.AddBook(B2)
	L.AddUser(U1)
	L.AddUser(U2)
	L.AddLoan(&B1, U1)
	L.AddLoan(&B2, U2)
	PrintService{}.ListBook(L.books)
	PrintService{}.ListUser(L.users)
	PrintService{}.PrintBook(B1)
	PrintService{}.PrintUser(U1)
	PrintService{}.ListLoan(L.loans)
	PrintService{}.ListLoan(L.loans)
	L.AddLoan(&B1, U1)
	L.AddLoan(&B2, U1)
	L.AddLoan(&B1, U1)
	L.AddLoan(&B1, U1)
	L.AddLoan(&B2, U1)
	L.AddLoan(&B2, U1)
	PrintService{}.ListLoan(L.loans)
	PrintService{}.PrintBook(B1)
	PrintService{}.PrintBook(B2)
	L.DeleteLoan(&B2) // elimina el primer prestado, no es objetivo este paso en este ejercicio
	PrintService{}.ListLoan(L.loans)
	PrintService{}.PrintBook(B1)
	PrintService{}.PrintBook(B2)

	// ejercicio 2
	Error1 := Error{"Fallo de disco", true}
	Error2 := Error{"Disco de arranque jodio", true}
	Error3 := Error{"estas muy comodo", false}

	repo := &ErrorRepository{}
	display := &DisplayError{}
	el := &ErrorLogged{}
	alerta := &Alert{}

	O := Orquester{}
	O.StartRegister(Error1, repo, el, alerta)
	O.StartRegister(Error2, repo, el, alerta)
	O.StartRegister(Error3, repo, el, alerta)

	O.DisplayRegister(repo, display)

	
}
