package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/readytotouch/readytotouch/internal/protos/organizers"
)

func main() {
	// review("./public/logos/original/")
	review("./public/logos/adapted/")
}

func review(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	imageExistsMap := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			panic("dir unexpected")
		}

		fileName := file.Name()
		if fileName == ".gitkeep" {
			continue
		}

		ext := filepath.Ext(file.Name())

		switch ext {
		case ".png", ".jpg", ".jpeg":
		default:
			panic(fmt.Sprintf("unexpected ext: %s", ext))
		}

		alias := strings.TrimSuffix(fileName, ext)

		imageExistsMap[fileName] = alias
	}

	aliasImageMap, err := fetchAliasImageMap("./aliases_right.txt")
	if err != nil {
		panic(err)
	}

	assertCorrectnessAliasImageMap(aliasImageMap, organizers.CompanyAliasMap, imageExistsMap)
}

func assertCorrectnessAliasImageMap(aliasImageMap map[string]string, aliasMap map[string]int64, imageExistsMap map[string]string) {
	for alias, image := range aliasImageMap {
		if _, ok := aliasMap[alias]; !ok {
			panic(fmt.Sprintf("alias not found: %s", alias))
		}

		if _, ok := imageExistsMap[image]; !ok {
			panic(fmt.Sprintf("image not found: %s", image))
		}
	}
}

func fetchAliasImageMap(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Ініціалізація мапи
	dataMap := make(map[string]string)

	// Читання рядків з файлу
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := parts[0]
		value := parts[1]

		if _, ok := dataMap[key]; ok {
			return nil, fmt.Errorf("duplicate key: %s", key)
		}

		dataMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dataMap, nil
}
