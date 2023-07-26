package writers

// Writer is the interface that every writer should implement.
type Writer interface {
	Write(row []string)

	WriteHeader()

	Flush()
}
