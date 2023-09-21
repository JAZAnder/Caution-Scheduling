package main

import (
	"database/sql"
	"encoding/json"
	//"fmt"

	//"fmt"
	"net/http"
	"strconv"

	//"github.com/bytedance/sonic/decoder"
	"github.com/gorilla/mux"
)

func (a *App) getLab(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid lab Id")
		return
	}

	l := lab{Id: id}
	if err := l.getLab(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Lab not Found")
		default : 
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, l)

}

func (a *App) getLabs(w http.ResponseWriter, r *http.Request){
	labs, err := getLabs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, labs)
}

func (a *App) createLab(w http.ResponseWriter, r *http.Request){
	var l lab

	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

    if len(l.Name) <= 0 || len(l.Location) <= 0 {
		l.print()
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

	err := l.createLab(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	

	respondWithJSON(w, http.StatusCreated, l)
}

func (a *App) updateLab(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid lab Id")
        return
    }

	var l lab
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&l); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    l.Id = id

    if err := l.updateLab(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, l)

}

func (a *App) deleteLab(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid lab ID")
        return
    }

    l := lab{Id: id}
    if err := l.deleteLab(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
