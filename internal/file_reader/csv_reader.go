package file_reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type CsvLine struct {
	Name  string
	Email string
}

type CsvLines struct {
	Lines []CsvLine
}

func line2Model(str []string) *CsvLine {
	csvLine := &CsvLine{
		Name:  str[0],
		Email: str[1],
	}
	return csvLine
}

func FileReader(path string) (*CsvLines, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	lines := &CsvLines{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		line := line2Model(record)

		lines.Lines = append(lines.Lines, *line)

		if err != nil {
			log.Fatal(err)
		}
	}
	return lines, err
}
