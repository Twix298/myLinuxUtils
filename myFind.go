package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	searchFiles := flag.Bool("f", false, "prints only files")
	searchDir := flag.Bool("d", false, "prints only directories")
	searchSyml := flag.Bool("sl", false, "prints only symlink")
	extension := flag.String("ext", "", "prints files with extension")
	flag.Parse()
	file := os.Args[len(os.Args)-1]
	if !*searchDir && !*searchFiles && !*searchSyml {
		*searchSyml = true
		*searchDir = true
		*searchFiles = true
	}
	err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

		if *searchDir && info.IsDir() {
			fmt.Println(path)
		}
		if !info.IsDir() {
			if *searchFiles {
				if *extension != "" {
					if filepath.Ext(path) == "."+*extension {
						fmt.Println(path)
					}
				} else {
					fmt.Println(path)
				}
			}
			if *searchSyml && (info.Mode()&os.ModeSymlink) != 0 {
				target, err := os.Readlink(path)
				_, errOpen := os.Open(target)
				if err != nil && errOpen != nil {
					fmt.Printf("%s -> [broken]\n", path)
				} else {
					fmt.Printf("%s -> %s\n", path, target)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
