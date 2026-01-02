package models

type Departure struct {
	Icao24                           string `json:"icao24"`
	FirstSeen                        int    `json:"firstSeen"`
	EstDepartureAirport              string `json:"estDepartureAirport"`
	LastSeen                         int    `json:"lastSeen"`
	EstArrivalAirport                string `json:"estArrivalAirport"`
	Callsign                         string `json:"callsign"`
	EstDepartureAirportHorizDistance int    `json:"estDepartureAirportHorizDistance"`
	EstDepartureAirportVertDistance  int    `json:"estDepartureAirportVertDistance"`
	EstArrivalAirportHorizDistance   int    `json:"estArrivalAirportHorizDistance"`
	EstArrivalAirportVertDistance    int    `json:"estArrivalAirportVertDistance"`
	DepartureAirportCandidatesCount  int    `json:"departureAirportCandidatesCount"`
	ArrivalAirportCandidatesCount    int    `json:"arrivalAirportCandidatesCount"`
}
