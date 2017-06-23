package epsongo

import "time"

type Header struct {
	Company string
	Address string
	NoTelp  string
}

type Subheader struct {
	Date     time.Time
	Cashier  string
	Waiter   string
	Customer string
	Number   string
}

type Sale struct {
	Subtotal        float64
	DiscountPercent int
	DiscountAmount  float64
	GrandTotal      float64
	Payment         float64
	TypePayment     string
	Charge          float64
	TaxPercent      int
	TaxAmount       float64
}

type Item struct {
	NameItem string
	Quantity int
	Satuan   string
	Price    float64
}

type Footer struct {
	ShowDate  bool
	ShowNote  bool
	Note      string
	WaterMark bool
}

func TemplateOne(p *Escpos, h Header, sh Subheader, sale Sale, items []Item, footer Footer) (err error) {
	p.Init()
	if err = PrintHeader(h, p); err != nil {
		return err
	}
	if err = PrintSubHeader(sh, p); err != nil {
		return err
	}
	if err = PrintItems(items, p); err != nil {
		return err
	}

	if err = PrintSale(sale, p); err != nil {
		return err
	}
	if err = PrintFooter(footer, p); err != nil {
		return err
	}
	p.Linefeed()
	p.Cut()
	p.End()
	return
}
