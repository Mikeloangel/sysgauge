package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Mikeloangel/sysgauge/cmd/server/handlers"
	"github.com/Mikeloangel/sysgauge/cmd/server/memstorage"
	"github.com/Mikeloangel/sysgauge/cmd/server/middlewares"
	"github.com/Mikeloangel/sysgauge/internal/config"
)

func main() {
	initStorage()
	serve()
}

func initStorage() {
	memstorage.InitMemstorage()
}

func serve() {
	mux := http.NewServeMux()

	mux.Handle(`/update/`,
		middlewares.Conveyor(
			http.HandlerFunc(handlers.Update),
			middlewares.UpdateValidator,
			middlewares.Post,
		),
	)

	mux.HandleFunc(`/get/`, func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/get/")
		parts := strings.Split(path, "/")
		mtype := parts[0]
		mkey := parts[1]

		if mtype == "counter" {
			v, e := memstorage.GetCounter(mkey)
			if e != nil {
				http.Error(w, "No key for counter", http.StatusNotFound)
				return
			}
			io.WriteString(w, fmt.Sprint(v))
			return
		}

		if mtype == "gauge" {
			v, e := memstorage.GetGauge(mkey)
			if e != nil {
				http.Error(w, "No key for gauge", http.StatusNotFound)
				return
			}
			io.WriteString(w, fmt.Sprint(v))
			return
		}

		http.Error(w, "Wrong type", http.StatusBadRequest)
	})

	err := http.ListenAndServe(config.ServerPort, mux)

	if err != nil {
		panic(err)
	}
}
