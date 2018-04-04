package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	lib "hnanalysis"
)

type reData struct {
	re  *regexp.Regexp
	str string
}

type hnData struct {
	nHN  int
	hits map[reData]int
}

func processCSV(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	reader := csv.NewReader(file)
	//reader.Comma = ';'
	row := 0
	timeIndex := -1
	textIndex := -1
	data := make(map[time.Time]hnData)
	var rexps []reData
	rexps = append(rexps, reData{str: "Kubernetes", re: regexp.MustCompile(`(?im)(kubernetes|k8s)`)})
	rexps = append(rexps, reData{str: "Mesos", re: regexp.MustCompile(`(?im)mesos`)})
	rexps = append(rexps, reData{str: "Cloud Foundry", re: regexp.MustCompile(`(?im)cloud\s+foundry`)})
	rexps = append(rexps, reData{str: "Docker Swarm", re: regexp.MustCompile(`(?im)docker\s+swarm`)})
	rexps = append(rexps, reData{str: "OpenStack", re: regexp.MustCompile(`(?im)openstack`)})
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		row++
		if row == 1 {
			for k, v := range record {
				switch v {
				case "time":
					timeIndex = k
				case "text":
					textIndex = k
				}
			}
			continue
		}
		utm, err := strconv.ParseInt(record[timeIndex], 10, 64)
		if err != nil {
			return err
		}
		tm := lib.MonthStart(time.Unix(utm, 0))
		text := record[textIndex]
		d, ok := data[tm]
		if !ok {
			h := make(map[reData]int)
			for _, rexp := range rexps {
				if rexp.re.MatchString(text) {
					h[rexp] = 1
					fmt.Printf("%v: %s first match\n", tm, rexp.str)
				} else {
					h[rexp] = 0
				}
			}
			data[tm] = hnData{
				nHN:  1,
				hits: h,
			}
		} else {
			d.nHN++
			for _, rexp := range rexps {
				if rexp.re.MatchString(text) {
					d.hits[rexp]++
					fmt.Printf("%v: %s #%d match (all posts: %d)\n", tm, rexp.str, d.hits[rexp], d.nHN)
				}
			}
			data[tm] = d
		}
	}
	fmt.Printf("data: %+v\n%d rows\n", data, row)
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
