package inputreader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filepath string) ([]string, error) {
	lines := make([]string, 0)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %v", filepath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
