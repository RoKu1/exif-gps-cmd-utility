package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/RoKu1/exif-cmd-utility/utils"
	"github.com/RoKu1/exif-cmd-utility/writers"
	csvwriter "github.com/RoKu1/exif-cmd-utility/writers/csvwriter"
	htmlwriter "github.com/RoKu1/exif-cmd-utility/writers/htmlwriter"
)

var fsema chan bool

var dsema chan bool

var dwg sync.WaitGroup

func main() {

	// Mandatory Parameters
	inputDirPath := flag.String("i", "", "Path to the input file (mandatory)")
	outputFilePath := flag.String("o", "", "Path to the output file (mandatory)")

	// Optional Flags
	helpFlag := flag.Bool("help", false, "Show usage information")
	maxFileWorkers := flag.Int("w", 1, "File Worker Count")

	flag.Parse()

	if *helpFlag {
		utils.PrintUsage()
		os.Exit(0)
	}

	if *maxFileWorkers > 10000 || *maxFileWorkers == 0 {
		*maxFileWorkers = 100
	}

	fsema = make(chan bool, *maxFileWorkers)
	dsema = make(chan bool, 1)

	// Check if the input file/directory exists and has sufficient permissions
	// fmt.Println(*inputDirPath)
	isDir, err := utils.IsDirectory(*inputDirPath)
	if err != nil || !isDir {
		log.Fatalf("Invalid input directory: %s", err)
	}

	isHTML := strings.HasSuffix(*outputFilePath, ".html")
	isCSV := strings.HasSuffix(*outputFilePath, ".csv")

	writeFile, err := os.Create(*outputFilePath)
	if err != nil {
		log.Fatalf("Error creating CSV file: %s", err)
	}
	defer writeFile.Close()

	var writer writers.Writer
	if isCSV {
		writer, err = csvwriter.NewCsvWriter(writeFile)
		if err != nil {
			log.Fatalf("Error creating CSVWriter file: %s", err)
		}
	} else if isHTML {
		writer, err = htmlwriter.NewHTMLWriter(writeFile)
		if err != nil {
			log.Fatalf("Error creating CSVWriter file: %s", err)
		}
	} else {
		log.Fatalf("Only .html and .csv supported for output file")
	}
	writer.WriteHeader()
	defer writer.Flush()
	dwg.Add(1)
	dsema <- true
	go ProcessDir(writer, *inputDirPath)
	dwg.Wait()
	close(fsema)
	close(dsema)

	if err != nil {
		log.Fatalf("Error walking through the directory: %s", err)
	}
}
