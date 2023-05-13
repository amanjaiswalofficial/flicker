package session

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type flickerdfReader struct {
    sourceType string
    format string
    props map[string]string
}

func (fdR flickerdfReader) Read(sourceType string) flickerdfReader {
    fdR.sourceType = sourceType
	return fdR
}

func (fdR flickerdfReader) Format(format string) flickerdfReader {
    fdR.format = format
    return fdR
}

func (fdR flickerdfReader) Option(key string, value string) flickerdfReader {
    if fdR.props == nil {
        fdR.props = make(map[string]string)
    }
    fdR.props[key] = value
    return fdR
}

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
       fdf.df= df
    }   
    return fdf
}

func (fdR flickerdfReader) Stop() {
    fmt.Println("Shutting down flicker...")
}