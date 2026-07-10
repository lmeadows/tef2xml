package tabedit_parser

import (
	"testing"
)

func TestParseTefFile(t *testing.T) {
	mockData := []map[string]string{
		{"col1": "value1", "col2": "value2"},
		{"col1": "value3", "col2": "value4"},
	}

	filePath := "mock_tef_data.txt"
	err := writeMockFile(filePath, mockData)
	if err != nil {
		t.Fatalf("Error writing mock file: %v", err)
	}
	defer os.Remove(filePath)

	data, err := ParseTefFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing TEF file: %v", err)
	}

	for i, row := range data {
		expectedRow := mockData[i]
		for colName, expectedValue := range expectedRow {
			actualValue, exists := row[colName]
			if !exists || actualValue != expectedValue {
				t.Errorf("Mismatch in row %d: %v", i+1, row)
			}
		}
	}

	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Error removing mock file: %v", err)
	}
}

func writeMockFile(filePath string, data []map[string]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return writer.Error()
}
