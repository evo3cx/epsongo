package epsongo

import (
	"fmt"
	"os"
	"os/exec"
)

type Printer struct {
	id           string
	name         string
	bus          string
	path         string
	deviceNumber string
	f            *os.File
}

var primaryPrintter, secondaryPrinnter Printer

func SetPrimaryPrintter(deviceNumber string) (err error) {
	cmd := exec.Command("lsusb", "-s", deviceNumber)

	b, err := cmd.Output()
	if err != nil {
		return
	}

	fmt.Println(string(b))
	return
}

func openPrinter() {

}

func deviceDescToPrinter(s string) Printer {
	p := Printer{}
	return p
}
