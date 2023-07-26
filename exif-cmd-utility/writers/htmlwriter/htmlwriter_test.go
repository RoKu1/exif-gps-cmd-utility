package writers

import (
	"os"
	"testing"
)

func TestHTMLWriter(t *testing.T) {
	tempHTMLFile, err := os.CreateTemp("", "test.html")
	if err != nil {
		t.Fatalf("Failed to create temporary HTML file: %v", err)
	}
	defer os.Remove(tempHTMLFile.Name())

	htmlWriter, err := NewHTMLWriter(tempHTMLFile)
	if err != nil {
		t.Fatalf("Failed to create HTML writer: %v", err)
	}

	htmlWriter.WriteHeader()

	dataRows := [][]string{
		{"file1.jpg", "12.345678", "98.765432"},
		{"file2.jpg", "34.567890", "78.901234"},
	}

	for _, row := range dataRows {
		htmlWriter.Write(row)
	}

	htmlWriter.Flush()

	if err := tempHTMLFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary HTML file: %v", err)
	}

	htmlContents, err := os.ReadFile(tempHTMLFile.Name())
	if err != nil {
		t.Fatalf("Failed to read HTML file contents: %v", err)
	}

	expectedContents := "<html>\n<head><title>EXIF Data</title></head>\n<body>\n" +
		"<h1>EXIF Data</h1>\n<table border=\"1\">\n" +
		"<tr><th>Image File Path</th><th>GPS Position Latitude</th><th>GPS Position Longitude</th></tr>\n" +
		"<tr><td>file1.jpg</td><td>12.345678</td><td>98.765432</td></tr>\n" +
		"<tr><td>file2.jpg</td><td>34.567890</td><td>78.901234</td></tr>\n" +
		"</table>\n</body>\n</html>"
	if string(htmlContents) != expectedContents {
		t.Errorf("Unexpected HTML file contents.\nExpected:\n%s\n\nGot:\n%s", expectedContents, string(htmlContents))
	}
}
