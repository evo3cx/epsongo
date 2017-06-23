package epsongo

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSetPrinter(t *testing.T) {

	if err := SetPrimaryPrintter("003"); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
}

func TestStringToPrinter(t *testing.T) {
	testCase := []struct {
		in  string
		out Printer
	}{
		{
			in: "Bus 001 Device 003: ID 04b8:0202 Seiko Epson Corp. Receipt Printer M129C/TM-T70",
			out: Printer{
				bus:          "001",
				deviceNumber: "003",
				id:           "04b8:0202",
				path:         "/dev/usb/lp0",
			},
		},
	}

	fmt.Print(testCase)
}

func TestGetListPrinter(t *testing.T) {}

func TestRunPrinter(t *testing.T) {}

func TestPrintHeader(t *testing.T) {
	h := Header{
		Company: "Tokoin",
		Address: "Jalan Gatot Subroto 127 Depok",
		NoTelp:  "192038128310",
	}

	f, err := os.OpenFile("tmbaprn:/ESDPRT001", os.O_RDONLY|os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	p.Init()
	if err = PrintHeader(h, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	p.Cut()
	p.End()

	w.Flush()

}

func TestPrintSubheader(t *testing.T) {
	h := Header{
		Company: "Tokoin",
		Address: "Jalan Gatot Subroto 127 Depok",
		NoTelp:  "192038128310",
	}

	sh := Subheader{
		Date:     time.Now(),
		Cashier:  "Reza",
		Waiter:   "Umum",
		Customer: "Umum",
		Number:   "212",
	}

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDONLY|os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	p.Init()
	if err = PrintHeader(h, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	if err = PrintSubHeader(sh, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	p.Linefeed()
	p.Cut()
	p.End()

	w.Flush()

}

func TestPrintItems(t *testing.T) {
	item := []Item{
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
	}

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDONLY|os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	p.Init()
	if err = PrintItems(item, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	p.Linefeed()
	p.Cut()
	p.End()

	w.Flush()
}

func TestFooter(t *testing.T) {
	h := Header{
		Company: "Tokoin",
		Address: "Jalan Gatot Subroto 127 Depok",
		NoTelp:  "192038128310",
	}

	sh := Subheader{
		Date:     time.Now(),
		Cashier:  "Reza",
		Waiter:   "Umum",
		Customer: "Umum",
		Number:   "212",
	}

	footer := Footer{
		ShowNote:  true,
		Note:      "Selamat Datang Kembali",
		WaterMark: true,
	}

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDONLY|os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	p.Init()
	if err = PrintHeader(h, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	if err = PrintSubHeader(sh, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	if err = PrintFooter(footer, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}
	p.Linefeed()
	p.Cut()
	p.End()

	w.Flush()
}

func TestPrintSale(t *testing.T) {

	sale := Sale{
		Subtotal:        3000,
		DiscountPercent: 10,
		DiscountAmount:  300,
		GrandTotal:      2700,
		Payment:         5000,
		TypePayment:     "cash",
		Charge:          2300,
		TaxPercent:      0,
	}

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDONLY|os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	if err = PrintSale(sale, p); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}

	p.Cut()
	p.End()

	w.Flush()
}

func TestTemplateOne(t *testing.T) {
	h := Header{
		Company: "Tokoin",
		Address: "Jalan Gatot Subroto 127 Depok",
		NoTelp:  "192038128310",
	}

	sh := Subheader{
		Date:     time.Now(),
		Cashier:  "Reza",
		Waiter:   "Umum",
		Customer: "Umum",
		Number:   "212",
	}

	sale := Sale{
		Subtotal:        3000,
		DiscountPercent: 10,
		DiscountAmount:  300,
		GrandTotal:      2700,
		Payment:         5000,
		TypePayment:     "cash",
		Charge:          2300,
		TaxPercent:      0,
	}

	footer := Footer{
		ShowNote:  true,
		Note:      "Selamat Datang Kembali",
		WaterMark: true,
	}

	items := []Item{
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
		Item{
			Price:    5000,
			Quantity: 1,
			Satuan:   "Pc",
			NameItem: "Pepsi",
		},
	}

	f, err := os.OpenFile("tmbaprn:/ESDPRT001", os.O_RDONLY|os.O_RDWR, 666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	p := New(w)

	if err = TemplateOne(p, h, sh, sale, items, footer); err != nil {
		t.Fatalf("expect nil get err %s", err)
	}

	w.Flush()
}
