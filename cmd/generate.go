package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"stgen/converter"
	"strings"
)

const (
	PARENTFOLDER = "site"
	PUBLICDIR = "public"
)

func Generate(path string) error {
	newpath := filepath.Join(path, PARENTFOLDER)

	if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
		log.Printf("error while creating path: %+v \n", newpath)
		return err
	}

	filepath.Walk(".", func(curpath string, info os.FileInfo, err error) error {
		// ignore hidden folders and files
		if filepath.Base(curpath)[0] == '.' || strings.HasPrefix(curpath, ".") {
			return nil
		}

		// incase of walk errors
		if err != nil {
			fmt.Println("error while traversing files: ", err)
			return nil
		}

		if info.IsDir() {
			os.Mkdir(filepath.Join(newpath, PUBLICDIR), os.ModePerm)
			return nil
		} else {
			return converter.Convert(curpath)
		}
	})
	return nil
}
