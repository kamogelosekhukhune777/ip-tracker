package cmd

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  "Trace the IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				getData(ip)
			}
		} else {
			fmt.Printf("Please provide IP to trace")
		}
	},
}

//07:00

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP           string `json:"ip"`
	HostName     string `json:"hostName"`
	AnyCast      bool   `json:"anyCast"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Location     string `json:"location"`
	Organisation string `json:"organisation"`
	Postal       string `json:"postal"`
	Timezone     string `json:"timezone"`
	ReadMe       string `json:"readMe"`
}

func getData(url string) {
	ip := strings.Trim(url, " ")
	urls := "http://ipinfo.io/" + ip + "/geo"

	respByte := showData(urls + ip)

	var data Ip
	err := json.Unmarshal(respByte, &data)
	if err != nil {
		log.Println("unable to unmarshal the data")
		return
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND:")
	fmt.Printf(
		"IP: %s\nCITY: %s\nREGION: %s\nCOUNTRY: %s\nLOCATION: %s\n TIMEZONE: %s\n POSTAL: %s\n",
		data.IP, data.City, data.Region, data.Country, data.Location, data.Timezone, data.Postal,
	)
}

func showData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("unable to get the response")
	}

	responseByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("unable to read the response")
	}

	return responseByte
}
