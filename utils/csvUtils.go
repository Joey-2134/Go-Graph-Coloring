package utils

import (
	"encoding/csv"
	"os"
	"strconv"
)

func WriteConflictsCSV(outPath string, conflictsPerIteration []int) error {
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	fieldnames := []string{"iteration", "conflicts"}
	if err := writer.Write(fieldnames); err != nil {
		return err
	}

	for i, conflicts := range conflictsPerIteration {
		row := []string{strconv.Itoa(i), strconv.Itoa(conflicts)}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return writer.Error()
}
