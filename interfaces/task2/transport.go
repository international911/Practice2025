package main

import "fmt"

type Transport interface {
	Move() string
	Stop() string
}

type Airplane struct {
	Model string
}

type Boat struct {
	Name string
}

func (a Airplane) Move() string {
	return fmt.Sprintf("Самолет %s взлетает", a.Model)
}

func (a Airplane) Stop() string {
	return fmt.Sprintf("Самолет %s приземляется", a.Model)
}

func (b Boat) Move() string {
	return fmt.Sprintf("Лодка %s плывет", b.Name)
}

func (b Boat) Stop() string {
	return fmt.Sprintf("Лодка %s останавливается у причала", b.Name)
}

func operateTransport(t Transport) {
	fmt.Println(t.Move())
	fmt.Println(t.Stop())
}

func main() {
	airplane := Airplane{Model: "Boeing 737"}
	boat := Boat{Name: "Северный Ветер"}

	operateTransport(airplane)
	operateTransport(boat)
}
