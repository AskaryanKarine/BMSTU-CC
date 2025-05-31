package fs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ReadProgramFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	input := strings.Join(lines, "\n")
	return input, nil
}

func SaveDOTToFile(dot, filenameDOT, filenamePNG string) error {
	file, err := os.Create(filenameDOT)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(dot)
	if err != nil {
		return err
	}
	cmd := exec.Command("dot", "-Tpng", "-o", filenamePNG, filenameDOT)
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}

	return nil
}
