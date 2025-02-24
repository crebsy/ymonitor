package workers

import (
	"log"
	"strconv"
	"yearn/ymonitor/config"
	"yearn/ymonitor/prom"
	"yearn/ymonitor/requests"

	"github.com/prometheus/client_golang/prometheus"
)

func SimpleMonitor(hosts chan config.Host, hostType string) {
	for host := range hosts {
		url := host.Url.Url
		log.Printf("querying %s %s, network %s\n", hostType, host.Name, host.Network)
		res, stats, err := requests.DoGetRequest(url.String())
		if err != nil {
			log.Print(err)
			continue
		}

		prom.Observe(stats, prometheus.Labels{
			"host":    host.Name,
			"network": host.Network,
			"code":    strconv.Itoa(res.StatusCode),
			"type":    hostType,
		})
	}
}
