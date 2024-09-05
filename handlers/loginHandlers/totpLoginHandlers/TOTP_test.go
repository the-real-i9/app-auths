package totpLoginHandlers

import (
	"fmt"
	"testing"
)

func TestGenerateBarcodeAndSetupKey(t *testing.T) {
	barcodeImageURL, setupkey := generateBarcodeAndSetupKey("oluwarinolasam@gmail.com")

	fmt.Println(barcodeImageURL, "\n\n", setupkey)
}
