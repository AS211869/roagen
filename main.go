package main

import (
	"encoding/json"
	"os"
	"fmt"
	"flag"
	"strings"
	"time"
)

type RipeRpkiExportMetadata struct {
	Generated int `json:"generated"`
	GeneratedTime string `json:"generatedTime"`
}

type RipeRpkiExportRoa struct {
	ASN string `json:"asn"`
	Prefix string `json:"prefix"`
	MaxLength int `json:"maxLength"`
	TA string `json:"ta"`
}

type RipeRpkiExport struct {
	Metadata RipeRpkiExportMetadata `json:"metadata"`
	ROAs []RipeRpkiExportRoa `json:"roas"`
}

func main() {
	dataFile := flag.String("data", "", "Location of the RIPE RPKI validator export file.")
	outDir := flag.String("out", "", "Directory where ROAgen should save the ROA files. Directory must already exist")
	flag.Parse()

	if (len(*dataFile) == 0) {
		fmt.Println("Data option is required")
		os.Exit(10)
	}

	if (len(*outDir) == 0) {
		fmt.Println("Out option is required")
		os.Exit(10)
	}

	// https://stackoverflow.com/a/16466189
	file, _ := os.Open(*dataFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	data := RipeRpkiExport{}
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}

	roa4, roa4Err := os.Create(fmt.Sprintf("%s/roa4.conf", *outDir))
	if roa4Err != nil {
		fmt.Println(err)
		os.Exit(11)
	}
	defer roa4.Close()

	roa6, roa6Err := os.Create(fmt.Sprintf("%s/roa6.conf", *outDir))
	if roa6Err != nil {
		fmt.Println(err)
		os.Exit(11)
	}
	defer roa6.Close()

	now := time.Now()
	genTime := fmt.Sprintf("# ROA generated at %s\n\n", now.Format(time.RFC1123Z))

	roa4.Write([]byte(genTime))
	roa6.Write([]byte(genTime))

	for _, roa := range data.ROAs {
		roaData := []byte(fmt.Sprintf("route %s max %d as %s; # %s\n", roa.Prefix, roa.MaxLength, strings.Replace(roa.ASN, "AS", "", 1), roa.TA))

		if strings.Contains(roa.Prefix, ":") {
			roa6.Write(roaData)
		} else {
			roa4.Write(roaData)
		}
	}
}