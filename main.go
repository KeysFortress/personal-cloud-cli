package main

import (
	"fmt"

	"personal_cloud_cli/implementations"
)

func main() {
	qrService := implementations.QRService{}

	qr, err := qrService.Get("This is a test")

	if err != nil {
		return
	}

	fmt.Println(qr)
}
