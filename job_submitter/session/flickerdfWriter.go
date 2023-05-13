package session

import "fmt"

type flickerdfWriter struct {
	mode     string
	filePath string
}

func (fdfWriter flickerdfWriter) Mode(mode string) flickerdfWriter {
	fdfWriter.mode = mode

	return fdfWriter
}

func (fdfWriter flickerdfWriter) Csv(filePath string) {
	fdfWriter.filePath = filePath
	fmt.Println("Writing to filePath destination")
}