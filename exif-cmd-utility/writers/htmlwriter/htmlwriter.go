package writers

import (
	"fmt"
	"os"
	"sync"
)

type HTMLWriter struct {
	mutex    *sync.Mutex
	htmlFile *os.File
}

func NewHTMLWriter(htmlFile *os.File) (*HTMLWriter, error) {
	return &HTMLWriter{htmlFile: htmlFile, mutex: &sync.Mutex{}}, nil
}

func (w *HTMLWriter) WriteHeader() {
	w.mutex.Lock()
	w.htmlFile.WriteString("<html>\n<head><title>EXIF Data</title></head>\n<body>\n")
	w.htmlFile.WriteString("<h1>EXIF Data</h1>\n<table border=\"1\">\n")
	w.htmlFile.WriteString("<tr><th>Image File Path</th><th>GPS Position Latitude</th><th>GPS Position Longitude</th></tr>\n")
	w.mutex.Unlock()
}

func (w *HTMLWriter) Write(row []string) {
	w.mutex.Lock()
	w.htmlFile.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n", row[0], row[1], row[2]))
	w.mutex.Unlock()
}

func (w *HTMLWriter) Flush() {
	w.mutex.Lock()
	w.htmlFile.WriteString("</table>\n</body>\n</html>")
	w.mutex.Unlock()
}
