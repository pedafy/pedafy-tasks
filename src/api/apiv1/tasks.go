package apiv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-tasks/src/database"
)

func (a *APIv1) taskGetAllHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var tasks []database.Tasks
	order, ok := r.URL.Query()["sort"]

	if ok && len(order[0]) > 0 {
		tasks, err = a.databaseHandler.GetAllTasksByOrder(order[0])
	} else {
		tasks, err = a.databaseHandler.GetAllTasks()
	}
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if len(tasks) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonTasks, err := json.Marshal(tasks)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonTasks)
		}
	}
}

func (a *APIv1) taskGetAllByFilterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tasks, err := a.databaseHandler.GetAllTasksByFilter(vars["id_kind"], vars["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if len(tasks) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"data":[]}`))
	} else {
		jsonTasks, err := json.Marshal(tasks)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			fmt.Fprintf(w, `{"data":%s}`, jsonTasks)
		}
	}
}

func (a *APIv1) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task database.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"error":"malormated data"}`))
		return
	}

	newTask, err := a.databaseHandler.NewTask(task)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else {
		jsonTask, err := json.Marshal(newTask)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, `{"data":%s}`, jsonTask)
		}
	}
}

func (a *APIv1) modifyTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	ID, _ := strconv.Atoi(ids)
	var task database.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"error":"malormated data"}`))
		return
	}

	modifiedTask, err := a.databaseHandler.ModifyTask(task, ID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if modifiedTask.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"error":"unknown data"}`))
	} else {
		jsonTask, err := json.Marshal(modifiedTask)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"data":%s}`, jsonTask)
		}
	}
}

func (a *APIv1) archiveTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	ID, _ := strconv.Atoi(ids)

	task, err := a.databaseHandler.ArchiveTask(ID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"unavailable"}`))
	} else if task.ID == 0 {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"error":"unknown data"}`))
	} else {
		jsonTask, err := json.Marshal(task)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error":"unavailable"}`))
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"data":%s}`, jsonTask)
		}
	}
}
