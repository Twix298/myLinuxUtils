package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func mustCheckCountFlag(lFlag, wFlag, mFlag bool) {
	var count int
	if lFlag {
		count++
	}
	if wFlag {
		count++
	}
	if mFlag {
		count++
	}
	if count > 1 {
		log.Fatal("Count flags not equal 1")
	}
}

func calcCountCharacter(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error open file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var countWords int
	for scanner.Scan() {
		countWords++
	}
	fmt.Println(strconv.Itoa(countWords) + "\t" + path)
}

func calcCountWords(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error open file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var countWords int
	for scanner.Scan() {
		countWords++
	}
	fmt.Println(strconv.Itoa(countWords) + "\t" + path)
}

func calcCountLine(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error open file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Счетчик строк
	lineCount := 0

	// Считываем файл построчно и увеличиваем счетчик для каждой строки
	for scanner.Scan() {
		lineCount++
	}
	fmt.Println(strconv.Itoa(lineCount) + "\t" + path)
}

func main() {
	lFlag := flag.Bool("l", false, "prints count line")
	wFlag := flag.Bool("w", false, "prints count symbols")
	mFlag := flag.Bool("m", false, "prints count word")
	flag.Parse()
	mustCheckCountFlag(*lFlag, *wFlag, *mFlag)
	if !*lFlag && !*wFlag && !*mFlag {
		*wFlag = true
	}
	var wg sync.WaitGroup
	if *lFlag {
		for i := 2; i < len(os.Args); i++ {
			file := os.Args[i]
			wg.Add(1)
			go calcCountLine(file, &wg)

		}
	} else if *wFlag {
		for i := 2; i < len(os.Args); i++ {
			file := os.Args[i]
			wg.Add(1)
			go calcCountWords(file, &wg)

		}
	} else if *mFlag {
		for i := 2; i < len(os.Args); i++ {
			file := os.Args[i]
			wg.Add(1)
			go calcCountCharacter(file, &wg)

		}
	}
	wg.Wait()
}
