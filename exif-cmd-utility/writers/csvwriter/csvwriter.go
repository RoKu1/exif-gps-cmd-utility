package writers

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
)

type CSVWriter struct {
	mutex     *sync.Mutex
	csvWriter *csv.Writer
}

func NewCsvWriter(csvFile *os.File) (*CSVWriter, error) {
	w := csv.NewWriter(csvFile)
	return &CSVWriter{csvWriter: w, mutex: &sync.Mutex{}}, nil
}

func (w *CSVWriter) WriteHeader() {
	w.mutex.Lock()
	header := []string{"Image File Path", "GPS Position Latitude", "GPS Position Longitude"}
	err := w.csvWriter.Write(header)
	if err != nil {
		log.Fatalf("error writing to csv %s", err)
	}
	w.mutex.Unlock()
}

func (w *CSVWriter) Write(row []string) {
	w.mutex.Lock()
	err := w.csvWriter.Write(row)
	if err != nil {
		log.Fatalf("error writing to csv %s", err)
	}
	w.mutex.Unlock()
}

func (w *CSVWriter) Flush() {
	w.mutex.Lock()
	w.csvWriter.Flush()
	w.mutex.Unlock()
}
