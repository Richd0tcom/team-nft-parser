package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type NFT struct {
	Format           string      `json:"format"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	MintingTool      string      `json:"minting_tool"`
	SensitiveContent bool        `json:"sensitive_content"`
	SeriesNumber     int         `json:"series_number"`
	SeriesTotal      int         `json:"series_total"`
	Attributess      []Trait     `json:"attributes"`
	Collection       Collections `json:"collection"`
	Gender           string      `json:"gender"`
	Uuid             string      `json:"uuid"`
}

type Collections struct {
	Name       string       `json:"name"`
	Id         string       `json:"id"`
	Attributes []Attributes `json:"attributes"`
}

type Attributes struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func main() {
	fmt.Println(`
	8888	888 888888888888 888888888888b.       8888888bb                                           
	8888b   888 888            	888           888   Y88b                                            
	88888b  888 888            	888           888    888                                            
	888Y88b 888 8888888       	888           888   d88P  8888b.  888d888 .d8888b   .d88b.  888d888 
	888 Y88b888 888            	888           8888888P"      "88b 888P"   88K      d8P  Y8b 888P"   
	888  Y88888 888            	888           888        .d888888 888     "Y8888b. 88888888 888     
	888   Y8888 888            	888           888        888  888 888          X88 Y8b.     888     
	888    Y888 888            	888           888        "Y888888 888      88888P'  "Y8888  888     `)

	var inputFileName string
	flag.StringVar(&inputFileName, "input", "", "A csv input file for the nfts")
	flag.Parse()

	outputCsv := filepath.Base(inputFileName)

	csvFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening csv file: ", err)
		panic(err)
	}
	reader := csv.NewReader(csvFile)

	//create a new output file.
	var filename = outputCsv
	var extension = filepath.Ext(filename)
	fmt.Println(extension)
	var newCsvBase = filename[0 : len(filename)-len(extension)]

	file, err := os.Create(newCsvBase + ".output.csv")
	if err != nil {
		fmt.Println("Error creating csv file: ", err)
		panic(err)
	}
	writer := csv.NewWriter(file)
	if _, err := os.Stat("nft"); os.IsNotExist(err) {

		os.Mkdir("nft", 0777)
		fmt.Println("Directory created, Parsing nfts...")
	} else {

		fmt.Println("Directory already exists,Parsing nfts...")
	}

	var nfts NFT

	line, _ := reader.Read()
	line = append(line, "sha256")
	writer.Write(line)

	counter := 1

	for {

		line, errors := reader.Read()
		if errors == io.EOF {
			break
		} else if errors != nil {
			log.Fatal(errors)
		}
		teamName := getTeamName(counter)

		num, _ := strconv.Atoi(line[1])
		// var attri []string

		atr := strings.Split(line[6], ";")
		items := []Trait{}
		for _, ele := range atr {

			attri := strings.Split(ele, ":")
			// if len(attri)>0{
			// 	attri = attri[:len(attri)-1]
			// }

			if len(attri) == 2 {
				at := Trait{
					TraitType: attri[0],
					Value:     attri[1],
				}
				fmt.Println(at)
				items = append(items, at)
			}

		}
		nfts = NFT{
			Format: "CHIP-0007",

			Name:             line[3],
			Description:      line[4],
			MintingTool:      teamName,
			SensitiveContent: false,
			SeriesNumber:     num,
			SeriesTotal:      420,
			Gender:           line[5],
			Attributess:      items,
			Uuid:             line[7],
			Collection: Collections{
				Name: "Zuri NFT Tickets for Free Lunch",
				Id:   "b774f676-c1d5-422e-beed-00ef5510c64d",
				Attributes: []Attributes{
					{
						Type:  "description",
						Value: "Rewards for accomplishments during HNGi9.",
					},
				},
			},
		}

		nftJson, _ := json.MarshalIndent(nfts, "", "  ")
		// nftJsonFile:= fmt.Sprintf()
		err = os.WriteFile("nft/"+line[2]+".json", nftJson, 0777)
		if err != nil {
			panic(err)
		}
		f, err := os.Open("nft/" + line[2] + ".json")

		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		x := h.Sum(nil)
		line = append(line, fmt.Sprintf("%x", x))
		_ = writer.Write(line)

		sha256.New().Reset()

		counter++
	}
	writer.Flush()
	//...................................
	//Writing struct type to a JSON file
	//...................................

	//Next use the flag package to prompt the user for input

}

func getTeamName(counter int) string {
	var teamName string
	switch {
	case counter > 400:
		teamName = "Team Boot"

	case counter > 380:
		teamName = "Team Axe"
	case counter > 360:
		teamName = "Team VBelt"

	case counter > 340:
		teamName = "Team Ruler"

	case counter > 320:
		teamName = "Team Sandpaper"

	case counter > 300:
		teamName = "Team Scale"

	case counter > 280:
		teamName = "Team Hydraulics"

	case counter > 260:
		teamName = "Team Powerdrill"

	case counter > 240:
		teamName = "Team Axle"

	case counter > 220:
		teamName = "Team Tape"

	case counter > 200:
		teamName = "Team Gear"

	case counter > 180:
		teamName = "Team Crankshaft"

	case counter > 160:
		teamName = "Team Headlight"

	case counter > 140:
		teamName = "Team Plug"

	case counter > 120:
		teamName = "Team Brainbox"

	case counter > 100:
		teamName = "Team Chisel"

	case counter > 80:
		teamName = "Team Prybar"

	case counter > 60:
		teamName = "Team Grit"

	case counter > 40:
		teamName = "Team Engine"

	case counter > 20:
		teamName = "Team Clutch"

	default:
		teamName = "Team Bevel"
	}
	return teamName
}

type Trait struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}
