package to

import (
	"time"
)

type EventTypeRequest struct {
	Filter MarketFilter
	Locale string
}

type MarketFilter struct {
	ExchangeIds []string
	TextQuery string
	EventTypeIds []string
	EventIds []string
	CompetitionIds []string
	Venues []string
	MarketBettingTypes []string
	MarketCountries []string
	MarketTypeCodes []string
	WithOrders []string
	BspOnly bool
	TurnInPlayEnabled bool
	InPlayOnly bool
	MarketStartTime timerange
}

type timerange struct {
	From time.Time
	To time.Time
}
