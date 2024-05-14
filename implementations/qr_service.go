package implementations

import (
	"bytes"

	qrcode "github.com/skip2/go-qrcode"
)

type QRService struct {
	bitmap [][]bool
}

func (q *QRService) toUpscaledString(inverseColor bool, upscaleFactor int) string {
	bits := q.bitmap

	var buf bytes.Buffer

	// Iterate over each row
	for _, row := range bits {
		// Iterate over each pixel in the row
		for _, pixel := range row {
			// Duplicate the pixel according to the upscale factor
			for i := 0; i < upscaleFactor; i++ {
				if pixel != inverseColor {
					buf.WriteString(" ")
				} else {
					buf.WriteString("█")
				}
			}
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

func (g *QRService) generateQRCode(data string) ([][]bool, error) {
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	return qr.Bitmap(), nil
}

func (g *QRService) Get(data string) (string, error) {
	bitmap, err := g.generateQRCode("Hello World")
	if err != nil {
		return "", err
	}
	g.bitmap = bitmap

	res := g.toUpscaledString(true, 2)
	return res, nil
}
