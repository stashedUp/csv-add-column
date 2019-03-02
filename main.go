package main

/*
* Version 0.0.1
* Compatible on command line - tested on mac
 */

/*** WORKFLOW ***/
/*
* 1- Read command line args - pass csv file
* 2- Loop through file, collect properties on each line (currently app is hardcoded to accept four columns)
* 3- Check if each email address is valid
* 4- Send information to email library
* 5- Connect to SES using AWS Credentials
* 6- Sends message using AWS SES
* 7- Returns OK message if sent was successful
 */

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pborman/getopt"
)

type Org struct {
	OrgName  string
	Position string
	Name     string
	Email    string
}

func main() {

	helpFlag := getopt.BoolLong("help", 'h', "displays help message")
	filePath := getopt.StringLong("file", 'f', "", "--file", "Path to csv file")
	firstCol := getopt.StringLong("first", 'A', "", "--first", "Add or replace first column")
	secondCol := getopt.StringLong("second", 'B', "", "--second", "Add or replace second column")
	thirdCol := getopt.StringLong("third", 'C', "", "--third", "Add or replace third column")
	forthCol := getopt.StringLong("forth", 'D', "", "--forth", "Add or replace forth column")

	getopt.Parse()

	if *helpFlag {
		usageMessage()
	} else if *filePath == "" {
		fmt.Println("You must provide a csv file")
		usageMessage()
	}

	csvFile, errorFile := os.Open(*filePath)

	if errorFile != nil {
		log.Fatal("Unable to open file")
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	now := time.Now()
	secs := now.Unix()

	rand := fmt.Sprintf("%v", secs)

	file, err := os.Create("./outgoingdata/" + rand + "result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	w := csv.NewWriter(file)
	fmt.Println("Attempting recompose csv")

	//var organizations []Org
	line, _ := reader.Read()
	fmt.Println(line)
	for {

		var org Org
		var organizations []string

		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		var orgname, position, name, email int

		if *firstCol != "" {
			org.OrgName = *firstCol
			position = 0
			name = 1
			email = 2
		} else {
			org.OrgName = line[orgname]
		}

		organizations = append(organizations, org.OrgName)

		if *secondCol != "" {
			org.Position = *secondCol
			name = 1
			email = 2
		} else {
			org.Position = line[position]
		}

		organizations = append(organizations, org.Position)

		if *thirdCol != "" {
			org.Name = *thirdCol
			email = 2
		} else {
			org.Name = line[name]
		}

		organizations = append(organizations, org.Name)

		if *forthCol != "" {
			org.Email = *forthCol
		} else {
			org.Email = line[email]
		}

		organizations = append(organizations, org.Email)
		// out, _ := json.Marshal(org)
		if err := w.Write(organizations); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func usageMessage() {
	fmt.Print("\n\n")
	getopt.PrintUsage(os.Stderr)
	fmt.Println("Pass path to csv file to -f or --file; example ./main -f test.csv")
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
