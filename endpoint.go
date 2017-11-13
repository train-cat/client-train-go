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
	stops                    = "/stops"
	EndpointStationStop      = EndpointStation + stops
	EndpointTrainStop        = EndpointTrain + stops
	EndpointStationTrainStop = EndpointStation + EndpointTrain + stops

	// Alerts
	EndpointAlert = "/alerts"
)
