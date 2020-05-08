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
	}
}

func main() {
	engineLogFilename := flag.String("engine", "engine.txt", "Engine log filename")
	logsFilename := flag.String("logs", "logs.txt", "Logs filename")

	flag.Parse()

	fmt.Printf("engineLogFilename '%s', logsfilename '%s'\n", *engineLogFilename, *logsFilename)

	// Actual system captures a 254 byte message every 1s and writes data every 60s
	// so 254 bytes * 60 = 15420 bytes every 1 minute

	engineDataLength := 254 * 60 // 254 bytes per message * 60 seconds of data
	logsDataLength := 256
	go writeToFile(*engineLogFilename, engineDataLength)
	go writeToFile(*logsFilename, logsDataLength)

	for true {
		time.Sleep(time.Second)
	}
}
