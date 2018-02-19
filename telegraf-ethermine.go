package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const baseUrl = "https://api.ethermine.org"

type historyResponse struct {
	Status string
	Data   []historyRecord
}

type historyRecord struct {
	Time             uint64
	ReportedHashrate float64
	CurrentHashrate  float64
	AverageHashrate  float64
	ValidShares      uint32
	InvalidShares    uint32
	StaleShares      uint32
	ActiveWorkers    uint32
}

func httpGet(url string) (body []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return body, err
}

func ethermineMinerHistory(minerAddress string) (response *historyResponse, err error) {
	responseText, err := httpGet(baseUrl + "/miner/" + minerAddress + "/history")
	if err != nil {
		return
	}

	response = &historyResponse{}
	err = json.Unmarshal(responseText, &response)
	if err != nil {
		return
	}

	return response, err
}

func usage() {

}

func main() {
	minerAddress := flag.String("address", "", "The ether account address to query, without the leading 0x.")
	flag.Parse()

	if *minerAddress == "" {
		flag.Usage()
		return
	}

	history, err := ethermineMinerHistory(*minerAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to download miner history from %s, %s.\n", baseUrl, err)
		return
	}

	last := history.Data[len(history.Data)-1]
	fmt.Printf(
		"ethermine,address=%s reported_hashrate=%f,current_hashrate=%f,average_hashrate=%f,valid_shares=%di,invalid_shares=%di,stale_shares=%di,active_workers=%di %d000000000\n",
		*minerAddress,
		last.ReportedHashrate,
		last.CurrentHashrate,
		last.AverageHashrate,
		last.ValidShares,
		last.InvalidShares,
		last.StaleShares,
		last.ActiveWorkers,
		last.Time)
}
