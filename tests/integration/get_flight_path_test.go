package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"flightTracker/internal/dep_container"

	"github.com/stretchr/testify/suite"
)

const (
	host = "http://localhost:8080"
)

type GetFlightPathTestSuite struct {
	suite.Suite
}

func TestAddLotToUserFavoritesTestSuite(t *testing.T) {
	suite.Run(t, new(GetFlightPathTestSuite))
}

func (s *GetFlightPathTestSuite) TestGetFlightPath_ok() {
	url := fmt.Sprintf("%s/%s", host, dep_container.GetFlightPathEndpoint)

	requstData := [][]string{{"IND", "EWR"}, {"EWR", "IND"}, {"IND", "ATL"}, {"ATL", "GSO"}}

	jsonValue, err := json.Marshal(requstData)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()
}
