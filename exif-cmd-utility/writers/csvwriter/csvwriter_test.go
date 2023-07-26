package writers

import (
	"os"
	"testing"
)

func TestCSVWriter(t *testing.T) {
	// Create a temporary CSV file for testing
	tempCSVFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary CSV file: %v", err)
	}
	defer os.Remove(tempCSVFile.Name())

	// Create a new CSVWriter
	csvWriter, err := NewCsvWriter(tempCSVFile)
	if err != nil {
		t.Fatalf("Failed to create CSV writer: %v", err)
	}

	// Write the header
	csvWriter.WriteHeader()

	// Write some data rows
	dataRows := [][]string{
		{"file1.jpg", "12.345678", "98.765432"},
		{"file2.jpg", "34.567890", "78.901234"},
	}

	for _, row := range dataRows {
		csvWriter.Write(row)
	}

	// Flush the writer
	csvWriter.Flush()

	// Close the CSV file
	if err := tempCSVFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary CSV file: %v", err)
	}

	// Read the contents of the CSV file
	csvContents, err := os.ReadFile(tempCSVFile.Name())
	if err != nil {
		t.Fatalf("Failed to read CSV file contents: %v", err)
	}

	// Verify the contents of the CSV file
	expectedContents := "Image File Path,GPS Position Latitude,GPS Position Longitude\nfile1.jpg,12.345678,98.765432\nfile2.jpg,34.567890,78.901234\n"
	if string(csvContents) != expectedContents {
		t.Errorf("Unexpected CSV file contents.\nExpected:\n%s\n\nGot:\n%s", expectedContents, string(csvContents))
	}
}
