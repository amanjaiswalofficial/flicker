package session

import (
    "fmt"
    "os"

    "github.com/go-gota/gota/dataframe"
)

// flickerdfReader represents a reader for loading data into a flickerdf.
// It contains information about the source type, format, and properties.
type flickerdfReader struct {
    sourceType string
    format     string
    props      map[string]string
}

// Read sets the source type for the flickerdfReader and returns the modified flickerdfReader.
func (fdR flickerdfReader) Read(sourceType string) flickerdfReader {
    fdR.sourceType = sourceType
    return fdR
}

// Format sets the format for the flickerdfReader and returns the modified flickerdfReader.
func (fdR flickerdfReader) Format(format string) flickerdfReader {
    fdR.format = format
    return fdR
}

// Option sets a key-value pair of properties for the flickerdfReader and returns the modified flickerdfReader.
func (fdR flickerdfReader) Option(key string, value string) flickerdfReader {
    if fdR.props == nil {
        fdR.props = make(map[string]string)
    }
    fdR.props[key] = value
    return fdR
}

// Load loads the data from the specified source using the reader's properties and returns a flickerdf.
func (fdR flickerdfReader) Load() flickerdf {
    fdf := flickerdf{}
    if fdR.format == "csv" && fdR.sourceType == "local" {
        readPath := fdR.props["path"]
        f, err := os.Open(readPath)
        if err != nil {
            panic(err)
        }
        defer f.Close()

        df := dataframe.ReadCSV(f)
        fdf.df = df
        fdf.sourceInfo = map[string]string{
            "format":     fdR.format,
            "sourceType": fdR.sourceType,
            "path":       readPath,
        }
    }
    return fdf
}

// Stop shuts down the flickerdfReader.
func (fdR flickerdfReader) Stop() {
    fmt.Println("Shutting down flicker...")
}
