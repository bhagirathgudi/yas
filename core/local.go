package main

import (
    "log"
    "os"
)

func CreateAndChngeDir() string{
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Chdir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func CleanUpDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}

