package room

import "time"

type Ticker struct {
	Time  time.Duration
	Event string
}

func cowTicker() map[string]*Ticker {
	tickers := make(map[string]*Ticker)
	tickers["start"] = &Ticker{
		Time:  time.Second * 2,
		Event: "draw",
	}
	tickers["master"] = &Ticker{
		Time:  time.Second * 3,
		Event: "draw",
	}
	tickers["draw"] = &Ticker{
		Time:  time.Second * 10,
		Event: "compare",
	}
	tickers["compare"] = &Ticker{
		Time:  time.Second * 10,
		Event: "ready",
	}
	tickers["ready"] = &Ticker{
		Time:  time.Second * 20,
		Event: "start",
	}
	return tickers
}
