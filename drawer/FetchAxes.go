package drawer

import (
	"sync"
	"time"

	"github.com/flamingyawn/discryptord/types"
)

// FetchAxes :: Fetch axes data
func FetchAxes(data []types.HistoMinute) types.AxesMap {
	var wg sync.WaitGroup
	// var ymin, ymax, volmin, volmax float64

	ymin, ymax := 1000000000.0, 0.0
	volmin, volmax := 1000000000.0, 0.0

	xq := make([]time.Time, len(data))
	yq := make([]float64, len(data))
	volq := make([]float64, len(data))

	for i, minute := range data {
		wg.Add(1)
		go func(i int, minute types.HistoMinute) {
			defer wg.Done()

			x := time.Unix(minute.Time, 0)
			y := minute.Close
			vol := minute.Volumeto

			if y < ymin {
				ymin = y
			}
			if y > ymax {
				ymax = y
			}
			if vol < volmin && vol > 0 {
				volmin = vol
			}
			if vol > volmax {
				volmax = vol
			}

			xq[i] = x
			yq[i] = y
			volq[i] = vol
		}(i, minute)

	}

	return types.AxesMap{
		X:      xq,
		Y:      yq,
		Vol:    volq,
		Ymin:   ymin,
		Ymax:   ymax,
		Volmin: volmin,
		Volmax: volmax,
	}
}