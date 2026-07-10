package tabedit_parser

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ParseTefFile reads a .tef file and returns its contents as a slice of maps.
func ParseTefFile(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []map[string]string
	for _, row := range rows {
		dataMap := make(map[string]string)
		for i, col := range row {
			dataMap[fmt.Sprintf("col%d", i+1)] = col
		}
		data = append(data, dataMap)
	}

	return data, nil
}
