package temperature

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_ReadTemperature(t *testing.T) {
	data := `76 01 4b 46 7f ff 0a 10 79 : crc=79 YES
76 01 4b 46 7f ff 0a 10 79 t=23375`
	tc := TC{
		Fahrenheit: true,
	}
	v, err := tc.readTemperature(strings.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	if v != 74.08 {
		t.Error("Expected 74.08 found:", v)
	}
}

func Test_InvalidTemperature(t *testing.T) {
	tc := TC{
		Fahrenheit: false,
	}

	data := `76 01 4b 46 7f ff 0a 10 79 : crc=79 YES
76 01 4b 46 7f ff 0a 10 79 t=-60375`

	_, err := tc.readTemperature(strings.NewReader(data))
	if err == nil {
		t.Error("value is out of range and should be an error")
	}

	data = `76 01 4b 46 7f ff 0a 10 79 : crc=79 YES
76 01 4b 46 7f ff 0a 10 79 t=156000`

	_, err = tc.readTemperature(strings.NewReader(data))
	if err == nil {
		t.Error("value is out of range and should be an error")
	}

}

func readFromFile(path string) (float32, error) {
	fi, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer fi.Close()
	data := make([]byte, 100)
	count, err := fi.Read(data)
	if err != nil {
		return 0, err
	}
	v := strings.TrimSpace(string(data[:count]))
	t, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, err
	}
	return float32(t), nil
}
