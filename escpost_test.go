package epsongo

import (
	"bufio"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	p.Init()
	p.SetSmooth(1)
	p.SetFontSize(2, 3)
	p.SetFont("A")
	_, err = p.Write("test ")
	if err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	p.SetFont("B")
	p.Write("test2 ")
	p.SetFont("C")
	p.Write("test3 ")
	p.Formfeed()

	p.SetFont("B")
	p.SetFontSize(1, 1)

	p.SetEmphasize(1)
	p.SetAlign("center")
	p.Write("halle")
	p.Formfeed()

	p.SetUnderline(1)
	p.SetFontSize(4, 4)
	p.Write("halle garis")
	p.Linefeed()

	p.SetReverse(1)
	p.SetFontSize(2, 4)
	p.Write("halle")
	p.Formfeed()

	p.SetFont("C")
	p.SetFontSize(8, 8)
	p.Write("halle")
	p.FormfeedN(5)

	p.Cut()
	p.End()

	w.Flush()
}
