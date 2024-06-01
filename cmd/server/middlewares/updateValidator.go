package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mikeloangel/sysgauge/internal/metrixunits"
)

func UpdateValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/update/")
		parts := strings.Split(path, "/")

		if len(parts) < 3 {
			http.Error(w, "Some required param are not set", http.StatusNotFound)
			return
		}

		metricType, metricName, metricValue := parts[0], parts[1], parts[2]

		if metricType != string(metrixunits.Gauge) && metricType != string(metrixunits.Counter) {
			http.Error(w, fmt.Sprintf("Metrc type %s is not supported", metricType), http.StatusBadRequest)
			return
		}

		var valueF float64
		var valueI int64
		var err error

		switch metricType {
		case string(metrixunits.Gauge):
			if valueF, err = strconv.ParseFloat(metricValue, 64); err != nil {
				http.Error(w, "Invalid gauge metric for value", http.StatusBadRequest)
				return
			}
		case string(metrixunits.Counter):
			if valueI, err = strconv.ParseInt(metricValue, 10, 64); err != nil {
				http.Error(w, "Invalid counter metric for value", http.StatusBadRequest)
				return
			}
		}

		metricUnit := metrixunits.MetricUnit{
			Type:   metrixunits.MetricType(metricType),
			Name:   metricName,
			ValueF: &valueF,
			ValueI: &valueI,
		}

		ctx := context.WithValue(r.Context(), MetricUnitKey, metricUnit)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
