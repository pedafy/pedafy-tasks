package apiv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *APIv1) getAllStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := a.databaseHandler.GetAllStatus()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if len(status) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"data":""}`))
	} else {
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonStatus)
		}
	}
}

func (a *APIv1) getStatusByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	id, _ := strconv.Atoi(ids)
	status, err := a.databaseHandler.GetStatusByID(id)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if status.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonStatus)
		}
	}
}

func (a *APIv1) getStatusByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	status, err := a.databaseHandler.GetStatusByName(name)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if status.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonStatus)
		}
	}
}
