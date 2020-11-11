package main

import (
	"log"
	"os"
	"time"
)

func main()  {
	if len(os.Args) < 2 {
		log.Fatal("You need to provide a file name.")
	}
	fileName := os.Args[1]
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(fileName, currentTime, currentTime)
		if err != nil {
			log.Fatal(err)
		}
	}
}
