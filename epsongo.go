package epsongo

import (
	"fmt"
	"strings"
	"time"

	"github.com/leekchan/accounting"
)

func money(price float64) string {
	acc := accounting.Accounting{
		Symbol:    "Rp.",
		Precision: 0,
	}
	return acc.FormatMoneyFloat64(price)
}

func PrintHeader(header Header, p *Escpos) (err error) {
	if err = p.SetAlign("center"); err != nil {
		return err
	}
	if _, err = p.Write(header.Company); err != nil {
		return err
	}
	p.Linefeed()
	if _, err = p.Write(header.Address); err != nil {
		return err
	}

	p.Linefeed()
	if _, err = p.Write(fmt.Sprintf("Telp: %s", header.NoTelp)); err != nil {
		return err
	}
	p.Linefeed()
	p.lineBorderDouble()
	p.Linefeed()

	return
}

func PrintSubHeader(sb Subheader, p *Escpos) (err error) {
	if err = p.SetAlign("left"); err != nil {
		return err
	}

	subheader := fmt.Sprintf(`Date : %s
Kasir : %s
Waiter : %s
Customer : %s
Number : %s`, sb.Date.Format(time.ANSIC),
		sb.Cashier,
		sb.Waiter,
		sb.Customer,
		sb.Number)

	if _, err = p.Write(subheader); err != nil {
		return err
	}
	p.Linefeed()
	p.lineBorderDouble()
	p.Linefeed()

	return
}

func PrintItems(items []Item, p *Escpos) (err error) {
	for _, item := range items {
		if err = printItem(item, p); err != nil {
			return
		}

	}
	p.lineBorderSingle()
	return nil
}

func printItem(item Item, p *Escpos) (err error) {
	if err = p.SetAlign("left"); err != nil {
		return err
	}
	name := fmt.Sprintf("%s %d %s", item.NameItem, item.Quantity, item.Satuan)
	if _, err = p.Write(name); err != nil {
		return
	}
	p.Linefeed()

	if err = p.SetAlign("right"); err != nil {
		return err
	}
	price := money(item.Price)
	if _, err = p.Write(price); err != nil {
		return
	}
	p.Linefeed()
	return nil
}

func percentString(percent int) string {
	return fmt.Sprintf("%d", percent) + "%"
}

func justifySale(st []string) (maxlength int, ss []string) {
	for _, s := range st {
		if len(s) > maxlength {
			maxlength = len(s)
		}
	}

	for _, s := range st {
		ss = append(ss, justifySaleDot(s, maxlength))
	}

	return maxlength, ss
}

func justifySaleDot(line string, length int) string {
	dotlength := length - len(line)
	dot := ""
	for i := 0; i < dotlength+3; i++ {
		dot += "."
	}

	line = strings.Replace(line, "...", dot, -1)
	return strings.Replace(line, ".", " ", -1)
}

func PrintSale(sale Sale, p *Escpos) (err error) {
	if err = p.SetAlign("right"); err != nil {
		return err
	}

	text := []string{
		fmt.Sprintf("Sub Total   :...%v", money(sale.Subtotal)),
		fmt.Sprintf("Discout %s  :...%v", percentString(sale.DiscountPercent), money(sale.DiscountAmount)),
		fmt.Sprintf("Pajak %s   :...%v", percentString(sale.TaxPercent), money(sale.TaxAmount)),
		"-------------------------",
	}

	_, text = justifySale(text)
	for _, t := range text {
		if _, err = p.Write(t); err != nil {
			return
		}
		p.Linefeed()
	}

	p.Linefeed()

	payment := []string{
		fmt.Sprintf("Grand Total   :...%s", money(sale.GrandTotal)),
		"-------------------------",
		fmt.Sprintf("%s   :...%v", sale.TypePayment, money(sale.Payment)),
		"-------------------------",
		fmt.Sprintf("Kembali   :...%s", money(sale.Charge)),
	}

	_, payment = justifySale(payment)
	for _, t := range payment {
		if _, err = p.Write(t); err != nil {
			return
		}
		p.Linefeed()
	}
	p.Linefeed()

	return nil
}

func PrintFooter(footer Footer, p *Escpos) (err error) {

	p.Linefeed()
	if err = p.SetAlign("center"); err != nil {
		return err
	}
	p.Linefeed()

	if footer.ShowNote {
		_, err = p.Write(footer.Note)
		if err != nil {
			return
		}
	}

	p.Linefeed()
	if footer.WaterMark {
		_, err = p.Write("power by tokoin.id")
		if err != nil {
			return
		}
	}
	p.Linefeed()
	return nil
}

func (p *Escpos) lineBorderDouble() (err error) {
	var border string
	for i := 0; i < 40; i++ {
		border += "="
	}

	_, err = p.Write(border)
	return err
}

func (p *Escpos) lineBorderSingle() (err error) {
	border := ``
	for i := 0; i < 40; i++ {
		border += "-"
	}

	_, err = p.Write(border)
	p.Linefeed()

	return err
}
