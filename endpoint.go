package traincat

// List of endpoint used by the client
const (
	EndpointLogin = "/login"

	// Station
	EndpointStations = "/stations"
	EndpointStation  = EndpointStations + "/%d"

	// Trip
	EndpointTrips = "/trips"
	EndpointTrip  = EndpointTrips + "/%d"
	endpointTerminus = "/terminus"

	// StopTime
	stopsTime                 = "/stops_time"
	stopTime                  = stopsTime + "/%d"
	EndpointStationStopsTime  = EndpointStation + stopsTime

	// Alerts
	EndpointAlerts        = "/alerts"
	EndpointStationAlerts = EndpointStation + stopTime + EndpointAlerts

	// Action
	EndpointActions = "/actions"
)
