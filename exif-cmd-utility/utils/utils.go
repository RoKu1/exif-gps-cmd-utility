package utils

import (
	"flag"
	"fmt"
	"os"

	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

func ExtractGPSInfo(filePath string) (latitude, longitude string, err error) {

	rawExif, err := exif.SearchFileAndExtractExif(filePath)
	if err != nil {
		return "0.0", "0.0", err
	}

	ifdmapping, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return "0.0", "0.0", nil
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(ifdmapping, ti, rawExif)
	if err != nil {
		return "0.0", "0.0", err
	}

	ifd, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return "0.0", "0.0", err
	}

	gi, err := ifd.GpsInfo()
	// fmt.Println(gi, gi.Latitude.Decimal())
	if err != nil {
		return "0.0", "0.0", nil
	}
	return fmt.Sprintf("%.6f", gi.Latitude.Decimal()), fmt.Sprintf("%.6f", gi.Longitude.Decimal()), err
}

func PrintUsage() {
	fmt.Println("Usage: exif-gpsinfo-reader [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// Function to check if the file name is hidden (starts with a dot)
func IsHiddenFile(name string) bool {
	return len(name) > 0 && name[0] == '.'
}
