package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"
)

var (
	// Version placeholder for -v option
	Version = "0"
	// CommitID placeholder for -v option
	CommitID = "0"
)

type result struct {
	Sales []sale `json:"sales"`
	Total int    `json:"total"`
	Query string `json:"query"`
}

type sale struct {
	Date    string `json:"date"`
	Price   string `json:"price"`
	Address string `json:"address"`
}

func main() {

	nextYear := strconv.Itoa(time.Now().Year() + 1)

	county := flag.String("c", "galway", "select county")
	yearFrom := flag.String("yf", "2010", "select year to search from")
	yearTo := flag.String("yt", nextYear, "select year to search to")
	output := flag.String("o", "table", "output to table (default) or json")
	version := flag.Bool("v", false, "prints current version")

	flag.Parse()

	if *version {
		fmt.Println(Version + " - " + CommitID)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "missing required address argument\n")
		os.Exit(2)
	}

	address := strings.Join(flag.Args(), " ")

	host := "www.propertypriceregister.ie"

	fullHost := "https://" + host

	path := "/Website/npsra/PPR/npsra-ppr.nsf/PPR-By-Date"

	query := fmt.Sprintf("&Start=1&Query=[dt_execution_date]>=01/01/%s AND [dt_execution_date]<01/01/%s AND [address]=*%s* AND [dc_county]=%s", *yearFrom, *yearTo, address, *county)

	url := fullHost + path + query

	var sales []sale

	c := colly.NewCollector(
		colly.AllowedDomains(host),
	)

	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	c.OnHTML("#content", func(content *colly.HTMLElement) {

		optionalNavURLs := content.ChildAttrs("#rcColRt > a", "href")

		content.ForEach("table.resultsTable tr", func(_ int, el *colly.HTMLElement) {

			if el.ChildText("td:nth-child(1)") != "" {

				sale := sale{}

				sale.Date = el.ChildText("td:nth-child(1)")
				sale.Price = el.ChildText("td:nth-child(2)")
				sale.Address = el.ChildText("td:nth-child(3)")

				sales = append(sales, sale)
			}
		})

		if len(optionalNavURLs) > 0 {
			if len(optionalNavURLs) > 1 {
				c.Visit(fullHost + optionalNavURLs[1])
			} else {
				c.Visit(fullHost + optionalNavURLs[0])
			}
		}
	})

	err := c.Post(url, map[string]string{})

	if err != nil {
		log.Fatal(err)
	}

	result := result{sales, len(sales), query}

	if *output == "json" {

		var jsonData []byte
		jsonData, err = json.MarshalIndent(result, "", "    ")

		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(jsonData))
	} else {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Date", "Address", "Price"})
		table.SetRowLine(true)

		for _, s := range sales {
			table.Append([]string{s.Date, s.Address, s.Price})
		}

		table.Render()
	}
}
