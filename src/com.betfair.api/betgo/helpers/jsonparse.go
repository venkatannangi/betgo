package helpers

import (
	"encoding/json"
	"fmt"
)


//market type definitions
type MaketType struct {
	MarketType string
	MarketCount int32
}

type MarketTypeResponse struct {
	Result []MaketType
}

//competitions definitions
type Competition struct {
	Id string
	Name string
}

type CompetitionCount struct {
	Competition Competition
	MarketCount int32
}

type CompetitionsResponse struct {
	Result []CompetitionCount
}

type RunnerPrice struct{
	Price float32
	Size float32
}

type RunnerExchangePrices struct{
	AvailableToBack []RunnerPrice
	AvailableToLay  []RunnerPrice
	TradedVolume []RunnerPrice
}

type RunnerPriceInfo struct {
	SelectionId int32
	Handicap float32
	Status string
	TotalMatched float32
	Ex RunnerExchangePrices

}

type MarketBook struct {
	MarketId string
	IsMarketDataDelayed bool
	Status string
	BetDelay int32
	BspReconciled bool
	Complete bool
	Inplay bool
	NumberOfWinners int32
	NumberOfRunners int32
	NumberOfActiveRunners int32
	TotalMatched float32
	TotalAvailable float32
	CrossMatching bool
	RunnersVoidable bool
	Version int64
	Runners []RunnerPriceInfo
}

type MarketBookResponse struct {
	Result []MarketBook
}

func ParseMarketTypeResponse(response string) []MarketTypeResponse{
	var jsonBlob = []byte(response)
	var responses []MarketTypeResponse

	err := json.Unmarshal(jsonBlob, &responses)

	if err != nil {
		fmt.Println("error:" ,err)
	}
	//fmt.Printf("%+v", responses)

	return 	responses
}


func ParseCompetitionsResponse(response string) []CompetitionsResponse{
	var jsonBlob = []byte(response)
	var responses []CompetitionsResponse

	err := json.Unmarshal(jsonBlob, &responses)

	if err != nil {
		fmt.Println("error:" ,err)
	}
	//fmt.Printf("%+v", responses)

	return 	responses
}

func ParseMarketBookResponse(response string) []MarketBookResponse{
	var jsonBlob = []byte(response)
	var responses []MarketBookResponse

	err := json.Unmarshal(jsonBlob, &responses)

	if err != nil {
		fmt.Println("error:" ,err)
	}
	//fmt.Printf("%+v", responses)

	return 	responses
}
