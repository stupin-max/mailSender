package file_reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CSVLine struct {
	Name  string
	Email string
}

type CSVLines struct {
	Lines []CSVLine
}

func parseCSVLine(record []string) (CSVLine, error) {
	if len(record) < 2 {
		return CSVLine{}, fmt.Errorf("invalid CSV record: expected at least 2 columns, got %d", len(record))
	}
	return CSVLine{
		Name:  record[0],
		Email: record[1],
	}, nil
}

func ReadCSV(path string) (*CSVLines, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Skip header row
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	lines := &CSVLines{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV record: %w", err)
		}

		line, err := parseCSVLine(record)
		if err != nil {
			return nil, err
		}

		lines.Lines = append(lines.Lines, line)
	}
	
	return lines, nil
}
