package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}
func createArchive(file string, buf *os.File, chanalErr chan<- error, wg *sync.WaitGroup) {
	gw := gzip.NewWriter(buf)
	tw := tar.NewWriter(gw)

	err := addToArchive(tw, file)
	if err != nil {
		chanalErr <- err
	} else {
		chanalErr <- nil
	}
	defer wg.Done()
	tw.Close()
	gw.Close()
}

func archive(args []string, fewFiles bool) {
	var targetDir string
	if fewFiles {
		targetDir = os.Args[2]
	}
	var wg sync.WaitGroup
	chanalErr := make(chan error)
	for _, file := range args {
		currentTime := time.Now().Unix()
		filename := filepath.Base(file) // Получаем имя файла с расширением
		name := filename[:len(filename)-len(filepath.Ext(filename))]
		archiveName := name + "_" + strconv.Itoa(int(currentTime)) + ".tar.gz"
		out, err := os.Create(targetDir + archiveName)
		if err != nil {
			log.Fatal("Error writing archive:", err)
		}
		defer out.Close()
		wg.Add(1)
		go createArchive(file, out, chanalErr, &wg)
		errCh := <-chanalErr
		if errCh != nil {
			log.Fatalln("Error creating archive:", err)
		}
	}
	wg.Wait()
	close(chanalErr)
	fmt.Println("Archive created successfully")
}

func main() {
	fewFiles := flag.Bool("a", false, "use for archive few files")
	flag.Parse()
	var args []string
	if *fewFiles {
		args = append(args, os.Args[3:]...)
	} else {
		args = append(args, os.Args[1:]...)
	}
	archive(args, *fewFiles)
}
