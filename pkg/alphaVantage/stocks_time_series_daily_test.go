package alphaVantage

import (
	"testing"
)

func TestTimeSeriesDaily_full(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesDaily("MSFT", SIZE_FULL)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Body
	date := *timeFromMap(quotes)
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].Volume)
}

func TestStocksTimeSeriesDaily_compact(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesDaily("MSFT", SIZE_COMPACT)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Make sure multiple quotes returned
	assertGreaterThan(t, 10, len(quotes))

	// Body
	date := *timeFromMap(quotes)
	assertStringEquals(t, "MSFT", quotes[date].Ticker)
	assertNotZero(t, float64(quotes[date].Time.Unix()))
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].Volume)
}
