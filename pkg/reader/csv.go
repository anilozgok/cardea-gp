package reader

import (
	"encoding/csv"
	"os"
)

func CSV(fileToRead string) ([][]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Comma = ','
	reader.Comment = '#'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
