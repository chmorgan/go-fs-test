package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func writeToFile(filename string, dataLength int) {
	interval := time.Millisecond * 500

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for true {
		str := fmt.Sprintf("+++%s---\n", randSeq(dataLength))

		data := []byte(str)
		f.Write(data)

		f.Sync()

		fmt.Printf("wrote %d bytes to %s\n", len(data), filename)

		time.Sleep(interval)
	}
}

func main() {
	engineLogFilename := flag.String("engine", "engine.txt", "Engine log filename")
	logsFilename := flag.String("logs", "logs.txt", "Logs filename")

	flag.Parse()

	fmt.Printf("engineLogFilename '%s', logsfilename '%s'\n", *engineLogFilename, *logsFilename)

	engineDataLength := 64
	logsDataLength := 256
	go writeToFile(*engineLogFilename, engineDataLength)
	go writeToFile(*logsFilename, logsDataLength)

	for true {
		time.Sleep(time.Second)
	}
}
