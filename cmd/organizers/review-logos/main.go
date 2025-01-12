package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
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

	var (
		sizeCountMap       = make(map[string]int)
		extCountMap        = make(map[string]int)
		totalSize          = int64(0)
		aliasFoundCountMap = make(map[bool]int)
	)

	for _, file := range files {
		if file.IsDir() {
			panic("dir unexpected")
		}

		fileName := file.Name()
		if fileName == ".gitkeep" {
			continue
		}

		filePath := filepath.Join(dir, fileName)

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			panic(err)
		}

		fileSize := fileInfo.Size()

		ext := filepath.Ext(file.Name())

		alias := strings.TrimSuffix(fileName, ext)

		var (
			width  int
			height int
		)
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			imgFile, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer imgFile.Close()

			img, _, err := image.Decode(imgFile)
			if err != nil {
				panic(err)
			}

			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
		}

		extCountMap[ext] += 1

		_, ok := organizers.CompanyAliasMap[alias]
		aliasFoundCountMap[ok] += 1
		delete(organizers.CompanyAliasMap, alias)

		if ok {
			continue
		}

		fmt.Println(fileName)
		/*
			fmt.Printf("Name: %s\n", fileName)
			fmt.Printf("Size: %d bytes\n", fileSize)
			fmt.Printf("Alias: %s (%t) \n", alias, ok)
			fmt.Printf("Size: %dx%d pixels\n", width, height)
			fmt.Println()
		*/
		sizeCountMap[fmt.Sprintf("%dx%d", width, height)] += 1
		totalSize += fileSize
	}
	/*
		aliases := make([]string, 0, len(organizers.CompanyAliasMap))
		for alias := range organizers.CompanyAliasMap {
			aliases = append(aliases, alias)
		}
		sort.Slice(aliases, func(i, j int) bool {
			return aliases[i] < aliases[j]
		})

		for _, alias := range aliases {
			fmt.Println(alias)
		}

		fmt.Println(sizeCountMap)
		fmt.Println(totalSize)
		fmt.Println(aliasFoundCountMap)
		fmt.Println(extCountMap)
	*/
}
