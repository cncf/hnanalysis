package main

import (
	"encoding/csv"
	"fmt"
	lib "hnanalysis"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// reData holds regexp and its string "name"
// It is read from "jobs.yaml" file
type reData struct {
	re    *regexp.Regexp `yaml:""`
	Str   string         `yaml:"name"`
	ReStr string         `yaml:"regexp"`
}

// jobs keep all jobs to process
// It is read from "jobs.yaml" file
type jobs struct {
	Jobs []reData `yaml:"jobs"`
}

// hnData holds number of hacker news posts and
// maps of reData hits
// there can be up to nHN hits for each reData regexp
type hnData struct {
	nHN  int
	hits map[reData]int
}

// processCSV Read "ifn" CSV file (BigQuery output)
// analyses ita nd saves results to "ofn"
func processCSV(ifn, ofn string) error {
	// if called with DEBUG=1 some extra info will be displayed
	debug := os.Getenv("DEBUG") != ""

	// Read CSV and close file
	iFile, err := os.Open(ifn)
	if err != nil {
		return err
	}
	defer func() { _ = iFile.Close() }()
	reader := csv.NewReader(iFile)
	//reader.Comma = ';'

	// Process CSV data
	rows := 0
	timeIndex := -1
	textIndex := -1

	// Main data structure
	data := make(map[time.Time]hnData)

	// Read defined jobs
	bytes, err := ioutil.ReadFile("jobs.yaml")
	if err != nil {
		return err
	}

	// Parse YAML
	var allJobs jobs
	err = yaml.Unmarshal(bytes, &allJobs)
	if err != nil {
		return err
	}
	for i, job := range allJobs.Jobs {
		allJobs.Jobs[i].re = regexp.MustCompile(job.ReStr)
	}
	rexps := allJobs.Jobs

	// Months
	tms := make(map[time.Time]struct{})
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		rows++
		// Get indexes of data rows we need
		if rows == 1 {
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

		// time is stored as unix seconds since epoch
		utm, err := strconv.ParseInt(record[timeIndex], 10, 64)
		if err != nil {
			return err
		}

		// Group by months
		tm := lib.MonthStart(time.Unix(utm, 0))
		tms[tm] = struct{}{}
		text := record[textIndex]

		// Analyse post text for given regexps
		d, ok := data[tm]
		if !ok {
			h := make(map[reData]int)
			for _, rexp := range rexps {
				if rexp.re.MatchString(text) {
					h[rexp] = 1
					if debug {
						fmt.Printf("%v: %s first match\n", tm, rexp.Str)
					}
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
					if debug {
						fmt.Printf("%v: %s #%d match (all posts: %d)\n", tm, rexp.Str, d.hits[rexp], d.nHN)
					}
				}
			}
			data[tm] = d
		}
	}

	// Sort all months found
	tmAry := lib.TimeAry{}
	for tm := range tms {
		tmAry = append(tmAry, tm)
	}
	sort.Sort(tmAry)
	if debug {
		fmt.Printf("data: %+v\n", data)
		fmt.Printf("dates: %+v\n", tmAry)
	}

	// Write output CSV
	oFile, err := os.Create(ofn)
	if err != nil {
		return err
	}
	defer func() { _ = oFile.Close() }()
	writer := csv.NewWriter(oFile)
	defer writer.Flush()

	// Header row "Month", "Hacker News Total" + all regexps processed
	hdr := []string{"Month", "Hacker News Total"}
	for _, rexp := range rexps {
		hdr = append(hdr, rexp.Str)
	}
	err = writer.Write(hdr)
	if err != nil {
		return err
	}

	// Data rows
	for _, tm := range tmAry {
		d, ok := data[tm]
		if !ok {
			fmt.Printf("WARNING: Missing data for %v\n", tm)
			continue
		}
		row := []string{lib.ToYMDDate(tm), strconv.Itoa(d.nHN)}
		for _, rexp := range rexps {
			row = append(row, strconv.Itoa(d.hits[rexp]))
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	// All OK, return
	fmt.Printf("Processed %d rows\n", rows)
	return nil
}

func main() {
	dtStart := time.Now()
	if len(os.Args) < 3 {
		fmt.Printf("%s: required input CSV file name (BigQuery output) and output CSV file name\n", os.Args[0])
		return
	}
	err := processCSV(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	dtEnd := time.Now()
	fmt.Printf("Time: %v\n", dtEnd.Sub(dtStart))
}
