package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "trace IP",
	Long:  `trace your IP-Address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				ShowData(ip)
			}
		} else {
			fmt.Println("Please provide an IP address to trace.")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type IP struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func ShowData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := GetData(url)

	data := IP{}

	json.Unmarshal(responseByte, &data)

	c := color.New(color.FgGreen).Add(color.Underline).Add(color.Bold)
	c.Println("Data Found:")

	c2 := color.New(color.FgCyan)
	c2.Printf("  IP: %s\n ", data.Ip)

	c3:=color.New(color.FgHiWhite)
	c3.Printf(" City: %s\n ", data.City)

	c4:=color.New(color.FgHiBlue)
	c4.Printf(" Region: %s\n ", data.Region)

	c5:=color.New(color.FgHiMagenta)
	c5.Printf(" Country: %s\n ", data.Country)

	c6:=color.New(color.FgRed)
	c6.Printf(" Location: %s\n ", data.Loc)

	c7:=color.New(color.FgHiYellow)
	c7.Printf(" Postal: %s\n ", data.Postal)

	c8:=color.New(color.FgHiRed)
	c8.Printf(" TimeZone: %s\n ", data.Timezone)
	
}

func GetData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("unable to get the reponse")
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to read the response")
	}

	return responseByte
}
