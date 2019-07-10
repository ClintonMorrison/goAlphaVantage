package alphaVantage

import "testing"

func TestTimeSeriesIntraday_full(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesIntraday("MSFT", INTERVAL_30, SIZE_FULL)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	date := *timeFromMap(quotes)
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].Volume)
}

func TestTimeSeriesIntraday_compact(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesIntraday("MSFT", INTERVAL_1, SIZE_COMPACT)

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