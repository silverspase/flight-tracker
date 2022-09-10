package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFlightPath(t *testing.T) {
	uc := New()

	var testCases = []struct {
		msg            string
		input          [][]string
		expectedOutput []string
		error          error
	}{
		{
			msg:            "success:visitEachAirportOneTime",
			input:          [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			expectedOutput: []string{"SFO", "EWR"},
			error:          nil,
		},
		{
			msg:            "success:visitOneAirportSeveralTimes",
			input:          [][]string{{"IND", "EWR"}, {"EWR", "IND"}, {"IND", "ATL"}, {"ATL", "GSO"}},
			expectedOutput: []string{"IND", "GSO"},
			error:          nil,
		},
		{
			msg:            "error:oneFlightShouldContain2AirportsError",
			input:          [][]string{{"IND"}},
			expectedOutput: nil,
			error:          errors.New(oneFlightShouldContain2AirportsError),
		},
		{
			msg:            "error:theFlightListIsNotConsistentError",
			input:          [][]string{{"IND", "EWR"}, {"ATL", "GSO"}},
			expectedOutput: nil,
			error:          errors.New(theFlightListIsNotConsistentError),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			res, err := uc.GetFlightPath(tc.input)
			assert.Equalf(t, tc.error, err, tc.msg)
			assert.Equalf(t, tc.expectedOutput, res, tc.msg)
		})
	}
}
