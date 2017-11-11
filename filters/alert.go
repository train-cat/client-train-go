package filters

type Alert struct {
	Pagination
	CodeTrain *string `query:"code_train"`
	StationID *int    `query:"station_id"`
}
