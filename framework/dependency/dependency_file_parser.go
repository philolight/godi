package dependency

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseConfigFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		line := strings.Trim(scanner.Text(), " \t\n\r")
		if strings.HasPrefix(line, "{}") {
			New(line)
		} else if strings.HasSuffix(line, "{") {
			client := strings.Trim(line[0:len(line)-1], CutSet)
			parseBean(client, scanner, lines)
		}
	}

	return nil
}

const CutSet = " \t\n\r"

func parseBean(client string, scanner *bufio.Scanner, lines int) error {
	for scanner.Scan() {
		lines++
		line := strings.Trim(scanner.Text(), CutSet)
		if line == "}" {
			return nil
		}

		equalIdx := strings.Index(line, "=")

		if equalIdx == -1 {
			return fmt.Errorf("%d: no equal : %s", lines, line)
		}

		field := strings.Trim(line[0:equalIdx], CutSet)
		subject := strings.Trim(line[equalIdx+1:], CutSet)

		if err := Set(client, field, subject); err != nil {
			return err
		}
	}

	return fmt.Errorf("parse config file error : no }")
}
