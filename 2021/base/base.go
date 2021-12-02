package base

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
)

func GetLines() []string {
	_, f, _, _ := runtime.Caller(1)
	cwd := path.Join(path.Dir(f))

	file, err := os.Open(cwd + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
