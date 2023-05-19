package session

import "fmt"

// flickerdfWriter represents a writer for writing data from a flickerdf.
// It contains information about the write mode and file path.
type flickerdfWriter struct {
	mode     string
	filePath string
}

// Mode sets the write mode for the flickerdfWriter and returns the modified flickerdfWriter.
func (fdfWriter flickerdfWriter) Mode(mode string) flickerdfWriter {
	fdfWriter.mode = mode
	return fdfWriter
}

// Csv sets the file path for writing the data in CSV format using the flickerdfWriter.
func (fdfWriter flickerdfWriter) Csv(filePath string) {
	fdfWriter.filePath = filePath
	fmt.Println("Writing to file path destination")
}
