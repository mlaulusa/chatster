package route

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mlaulusa/chatster/model"
)

func GetRoomRouter() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/room").Methods("GET").HandlerFunc(getAll)
	router.PathPrefix("/room").Methods("POST").HandlerFunc(save)

	return router
}

func save(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	err := json.NewDecoder(r.Body).Decode(&room)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = model.SaveRoom(&room)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)

}

func getAll(w http.ResponseWriter, r *http.Request) {
	rooms, err := model.GetAllRooms()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)

}
