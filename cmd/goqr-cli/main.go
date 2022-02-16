package main

import (
	"flag"
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
	"github.com/pschlump/goqrsvg"
)

var Url = flag.String("url", "", "URL to encode into the QR code.")

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "goqr-cli: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Printf("Extra arguments are not supported [%s]\n", fns)
		os.Exit(1)
	}

	genqr(*Url, os.Stdout)
}

func genqr(uu string, fp *os.File) {
	s := svg.New(fp)

	// Create the QR Code in SVG
	qrCode, _ := qr.Encode(uu, qr.M, qr.Auto)

	// Write QR code to SVG
	qs := goqrsvg.NewQrSVG(qrCode, 5)
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)

	s.End()
}
