package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/matthieukhl/go-atc/internal/models"
)

const (
	departuresEndpoint = "/flights/departure/"
)

func GetFlightDepartures(client Client, airportICAO string) error {

	// Query params
	// begin := time.Now().UTC().AddDate(0, 0, -1).Unix()
	// end := time.Now().UTC().AddDate(0, 0, 1).Unix()

	begin := "1517227200"
	end := "1517230800"

	endpoint := baseUrl + departuresEndpoint

	reqUrl, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	// Add query parameters
	query := reqUrl.Query()
	query.Add("airport", airportICAO)
	// query.Add("begin", strconv.Itoa(int(begin)))
	// query.Add("end", strconv.Itoa(int(end)))
	query.Add("begin", begin)
	query.Add("end", end)
	reqUrl.RawQuery = query.Encode()
	fmt.Println(reqUrl.String())
	fmt.Println(client.Config.OpenSkyApiKey)

	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)

	bearer := "Bearer " + client.Config.OpenSkyApiKey
	req.Header.Add("Authorization", bearer)

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		fmt.Printf("Error with status code: %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	departures := []models.Departure{}
	err = json.Unmarshal(data, &departures)

	fmt.Println(departures)

	return nil
}
