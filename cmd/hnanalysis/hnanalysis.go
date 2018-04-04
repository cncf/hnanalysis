package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	lib "hnanalysis"
)

func processCSV(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	reader := csv.NewReader(file)
	//reader.Comma = ';'
	row := 0
	idIndex := -1
	parentIndex := -1
	timeIndex := -1
	titleIndex := -1
	textIndex := -1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		row++
		//fmt.Printf("Record %d is '%v' and has %d fields: ", row+1, record, len(record))
		if row == 1 {
			for k, v := range record {
				switch v {
				case "id":
					idIndex = k
				case "parent":
					parentIndex = k
				case "time":
					timeIndex = k
				case "title":
					titleIndex = k
				case "text":
					textIndex = k
				}
			}
			fmt.Printf("Data indices: %d,%d,%d,%d,%d\n", idIndex, parentIndex, timeIndex, titleIndex, textIndex)
			continue
		}
		id := record[parentIndex]
		if id == "" {
			id = record[idIndex]
		}
		utm, err := strconv.ParseInt(record[timeIndex], 10, 64)
		if err != nil {
			return err
		}
		tm := lib.MonthStart(time.Unix(utm, 0))
		fmt.Printf("%v\n", tm)
	}
	fmt.Printf("%d rows\n", row)
	return nil
}

func main() {
	dtStart := time.Now()
	if len(os.Args) < 2 {
		fmt.Printf("%s: required CSV file name (BigQuery output)\n", os.Args[0])
		return
	}
	err := processCSV(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	dtEnd := time.Now()
	fmt.Printf("Time: %v\n", dtEnd.Sub(dtStart))
}
