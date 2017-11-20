package filters

// Stop filters
type Stop struct {
	Pagination
	TerminusID            *int    `query:"terminus_id"`
	TrainThroughStationID *uint   `query:"train_through_station_id"`
	ScheduledBefore       *string `query:"scheduled_before"`
	ScheduledAfter        *string `query:"scheduled_after"`
	ScheduledAt           *string `query:"scheduled_at"`
}
