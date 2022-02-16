package main

import (
	"fmt"
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
	"github.com/pschlump/goqrsvg"
)

func main() {
	http.Handle("/", http.HandlerFunc(genqr))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func genqr(www http.ResponseWriter, req *http.Request) {

	// Pull in a paramter - 'url' if get or 'url' from post
	var uu string
	if req.Method == "GET" {
		keys, ok := req.URL.Query()["url"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'url' is missing")
			www.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(www, "Missing 'url' paramter\n")
			return
		}

		// Query()["key"] will return an array of items, we only want first item.
		uu = keys[0]
	} else if req.Method == "POST" {
		req.ParseForm()          // Parses the request body
		uu = req.Form.Get("url") // x will be "" if parameter is not set
		if uu == "" {
			log.Println("Url Param 'url' is missing")
			www.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(www, "Missing 'url' paramter\n")
			return
		}
	} else {
		www.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintf(www, "Invalid Method")
		return
	}

	s := svg.New(www)
	www.Header().Set("Content-Type", "image/svg+xml")

	// Create the QR Code in SVG
	// qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)
	qrCode, _ := qr.Encode(uu, qr.M, qr.Auto)

	// Write QR code to SVG
	qs := goqrsvg.NewQrSVG(qrCode, 5)
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)

	s.End()
}
