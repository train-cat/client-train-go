package traincat

const (
	EndpointLogin = "/login"

	// Station
	EndpointStations = "/stations"
	EndpointStation = EndpointStations + "/%d"

	// Train
	EndpointTrains = "/trains"
	EndpointTrain = EndpointTrains + "/%s"

	// Station - Train
	EndpointStationTrainStop = EndpointStation + EndpointTrain + "/stops"
)
