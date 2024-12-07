package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type TCReader struct {
	file      *os.File
	reader    *bufio.Scanner
	delimiter string
}

func NewTCReader(filepath string, delimiter string) (*TCReader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
		return nil, err
	}

	tcReader := &TCReader{
		reader:    scanner,
		delimiter: delimiter,
	}

	return tcReader, nil
}

func (tcr *TCReader) Close() {
	if tcr.file != nil {
		tcr.file.Close()
	}
}

func (tcr *TCReader) Next() []string {
	row := tcr.reader.Text()
	fields := strings.Split(row, tcr.delimiter)
	return fields
}

func (tcr *TCReader) Scan() bool {
	isReadable := tcr.reader.Scan()
	return isReadable
}
