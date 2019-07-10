package alphaVantage

import (
	"testing"
)


func TestTimeSeriesWeeklyAdjusted(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesWeeklyAdjusted("MSFT")

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Body
	date := *timeFromAdjustedTimeSeries(quotes)
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].AdjustedClose)
	assertNotZero(t, quotes[date].Volume)
}