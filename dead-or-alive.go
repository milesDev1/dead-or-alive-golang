package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var octets []string = []string{}
var valuecount int = 0
var madeIP string

type CsvLine struct {
	Column1 string
}

func main() {
	absolutePath := flag.String("path", "./office.csv", "path to csv goes here")
	flag.Parse()
	csvHandler(*absolutePath)
	pingService(octets)
}

func pingService(a []string) {
	for i := 0; i < len(octets); i++ {
		madeIP = fmt.Sprintf("172.16.%v.35", octets[i])
		out, _ := exec.Command("ping", madeIP).Output()
		if strings.Contains(string(out), "Request timed out.") {
			logService("DEAD")
		} else if strings.Contains(string(out), "bytes=32") {
			logService("ALIVE")
		} else if strings.Contains(string(out), "Destination host unreachable.") {
			logService("UNREACHABLE")
		} else {
			logService("BAD HOST")
		}
	}
}

func logService(s string) {
	logfile, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	log.Printf("%v = %v\n", madeIP, s)
	valuecount += 1
	fmt.Printf("Write operation number %d success at %v\n", valuecount, time.Now())
	return
}

func csvHandler(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	readcsv := csv.NewReader(f)
	records, err := readcsv.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range records {
		data := CsvLine{
			Column1: line[0],
		}

		split := strings.Split(data.Column1, " ")
		unprocessed := split[len(split)-1]
		processed := strings.TrimLeft(strings.TrimRight(unprocessed, "]"), "[")

		octets = append(octets, processed)
	}

	return
}
