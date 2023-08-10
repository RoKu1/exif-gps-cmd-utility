package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/RoKu1/exif-cmd-utility/utils"
	"github.com/RoKu1/exif-cmd-utility/writers"
)

func ProcessFile(writer writers.Writer, filepathstr string, fwg *sync.WaitGroup) error {
	defer func() {
		<-fsema
		fwg.Done()
	}()
	absolutePath, err := filepath.Abs(filepathstr)
	if err != nil {
		fmt.Printf("Error converting to absolute path: %s\n", err)
		os.Exit(1)
	}
	latitude, longitude, err := utils.ExtractGPSInfo(filepathstr)
	if err != nil {
		writer.Write([]string{absolutePath, "NotFound", "NotFound"})
	} else {
		writer.Write([]string{absolutePath, latitude, longitude})
	}
	fmt.Printf("Found and Proccesed path: %s\n", absolutePath)
	return nil
}

func ProcessDir(writer writers.Writer, dirpath string) error {

	var filepaths []string
	var dirpaths []string
	var fwg sync.WaitGroup

	defer func() {
		fwg.Wait()
		dwg.Done()
	}()

	dir, err := os.Open(dirpath)
	if err != nil {
		return fmt.Errorf("error opening directory: %s", err)
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("error reading directory entries: %s", err)
	}

	// Process each directory entry
	for _, entry := range entries {
		path := dirpath + "/" + entry.Name()

		if entry.IsDir() {
			dirpaths = append(dirpaths, path)
		} else if !utils.IsHiddenFile(entry.Name()) {
			filepaths = append(filepaths, path)
		}
	}

	for _, filepath := range filepaths {
		fsema <- true
		fwg.Add(1)
		filepath := filepath
		go ProcessFile(writer, filepath, &fwg)
	}
	<-dsema

	for _, dirpath := range dirpaths {
		dsema <- true
		dwg.Add(1)
		dirpath := dirpath
		go ProcessDir(writer, dirpath)
	}
	return nil
}
