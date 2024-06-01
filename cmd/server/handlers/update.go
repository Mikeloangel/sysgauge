package handlers

import (
	"net/http"

	"github.com/Mikeloangel/sysgauge/cmd/server/memstorage"
	"github.com/Mikeloangel/sysgauge/cmd/server/middlewares"
	"github.com/Mikeloangel/sysgauge/internal/metrixunits"
)

func Update(w http.ResponseWriter, r *http.Request) {
	metricUnit, ok := r.Context().Value(middlewares.MetricUnitKey).(metrixunits.MetricUnit)

	if !ok {
		http.Error(w, "Update metric unit error", http.StatusInternalServerError)
		return
	}

	memstorage.Update(metricUnit)

	w.Write([]byte("Ok!\n"))
}
