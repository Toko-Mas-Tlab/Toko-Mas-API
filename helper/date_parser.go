package helper

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func DateParseToBahasa(t time.Time) string {
	// init locale
	indonesian := language.MustParse("id")
	printer := message.NewPrinter(indonesian)

	// date format
	formattedDate := t.Format("2 January 2006")

	// parse to bahasa
	bulan := t.Month()
	bulanIndonesia := printer.Sprintf("%v", bulan)

	result := fmt.Sprintf("%s %s %d", formattedDate[:2], bulanIndonesia, t.Year())

	return result
}
