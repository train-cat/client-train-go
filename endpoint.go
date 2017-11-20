package traincat

// List of endpoint used by the client
const (
	EndpointLogin = "/login"

	// Station
	EndpointStations = "/stations"
	EndpointStation  = EndpointStations + "/%d"

	// Train
	EndpointTrains = "/trains"
	EndpointTrain  = EndpointTrains + "/%s"

	// Stops
	stops                     = "/stops"
	stop                      = "/stops/%d"
	EndpointStationStop       = EndpointStation + stops
	EndpointTrainStop         = EndpointTrain + stops
	EndpointStationTrainStops = EndpointStation + EndpointTrain + stops
	EndpointStationTrainStop  = EndpointStation + stop

	// Alerts
	EndpointAlerts        = "/alerts"
	EndpointStationAlerts = EndpointStation + EndpointTrain + EndpointAlerts

	// Action
	EndpointActions = "/actions"
)
