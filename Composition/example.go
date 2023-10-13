package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID      int
	Name    string
	Price   float64
	MadeIn  string
	Barcode string
	Color   string
	Model   int
}

type ElectronicProduct struct {
	Product    //embedded-struct
	ScreenSize int
	Ram        int
	Cpu        string
	OS         string
}

type Mobile struct {
	ElectronicProduct //embedded-struct
	SimNumber         int
	CameraType        string
	NetworkType       string
}

type Laptop struct {
	ElectronicProduct //embedded-struct
	UsbPortNumber     int
	KeyBoardType      string
	HasCDRom          string
}

func main() {
	hpLaptop := new(Laptop)

	samsungPhone := new(Mobile)

	samsungPhone.Name = "Samsung S21"
	hpLaptop.OS = "Ubuntu"

	samsungJson, _ := json.Marshal(samsungPhone)
	hpJson, _ := json.Marshal(hpLaptop)

	fmt.Println(string(samsungJson))
	fmt.Println(string(hpJson))
}
