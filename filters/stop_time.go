package filters

// StopTime filters
type StopTime struct {
	Pagination
	TrainThroughStationID *int    `query:"train_through_station_id"`
	ScheduledBefore       *string `query:"scheduled_before"`
	ScheduledAfter        *string `query:"scheduled_after"`
	ScheduledAt           *string `query:"scheduled_at"`
}
