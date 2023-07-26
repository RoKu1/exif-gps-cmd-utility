# EXIF GPS info extraction utility for files in Go
This command-line utility written in Go extracts EXIF data from images and writes the following attributes to a CSV or HTML file:

- Image file path
- GPS position latitude
- GPS position longitude

## Prerequisites
Before using this utility, ensure you have the following:

- Go installed on your system (at least Go 1.19 or higher).
- Image files with EXIF data (e.g., JPEG files) in the specified directory and sub-directory.

## Installation
1. Change to the project directory: 
    ```bash
    cd exif-cmd-utilty
    ```


2. Install the required Go packages: 
    ```bash
    go mod tidy
    ```


# Usage
To use the utility, run the following command:
   
```bash
go run main.go -i <input_directory> -o <output_csv_file> -w <optionalparam>
```

    -i <input_directory>

    -o <output_csv_file>

    -n <optionalfield-fileworkers-default-1


Replace <input_directory> with the path to the directory that contains the images with EXIF data. Replace <output_csv/html_file> with the desired path and name of the CSV file where the extracted attributes will be written.

#### For example:

```bash
go run ./ -i ./images -o output.csv
go run ./ -i ./images -o output.html

go run ./ -i ./images -o output.csv -n 10
go run ./ -i ./images -o output.html -n 10
```

## Running Using Binary
Aditionaly you can copy the binary of the utility as per your system from the workdir called exif-cmd-utilty-<your os> and run it anywhere of your choice as below, this is the prefered method

#### For example for mac system:
```bash
./exif-cmd-utilty-mac -i ./images -o output.csv
./exif-cmd-utilty-mac -i ./images -o output.html

./exif-cmd-utilty-mac -i ./images -o output.csv -n 10
./exif-cmd-utilty-mac -i ./images -o output.html -n 10
```

# Notes 
This utility uses the "go-exif" library to extract EXIF data from the images. Ensure that the images you provide have valid EXIF data for GPS position latitude and longitude.

The CSV file generated will have three columns: "Image File Path," "GPS Position Latitude," and "GPS Position Longitude."

The utility will exclude hidden files (files that start with a dot) and only process image files (e.g., JPEG files).

If the EXIF data for GPS position latitude and longitude is not available for an image, the corresponding fields in the CSV file will be empty.

# Acknowledgments
The "go-exif" library used in this utility is created by Daniel Soprea. Visit the library's repository [here](https://github.com/dsoprea/go-exif/).