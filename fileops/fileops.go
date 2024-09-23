package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func File(balance float64, fileName string) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceTxt), 0644)
}

func ReadFile(FileName string) (float64, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		return 10000, errors.New("Failed to find a file!")
	}

	accountBalance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 10000, errors.New("failed to convert!")
	}

	return accountBalance, nil
}
