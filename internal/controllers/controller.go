package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"vg-sage/internal/models"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func getFilter(r *http.Request) (error, models.Filter) {
	filterParsed := models.Filter{}
	query := r.URL.Query()
	filter := query.Get("filter")

	if filter == "" {
		return nil, filterParsed
	}
	err := json.Unmarshal([]byte(filter), &filterParsed)
	// fuerza el código de empresa en función del entorno.
	cod := os.Getenv("CODIGO_BBDD")
	if cod == "" {
		cod = "1"
	}
	// filterParsed.Where["CodigoEmpresa"] = 9999
	return err, filterParsed
}

func FindAll(point interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err, filterParsed := getFilter(r)
		if err != nil {
			fmt.Println(err)
		}
		models.Find(filterParsed, point)
		respondJSON(w, http.StatusOK, point)
	}
}

func Find(point interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, filterParsed := getFilter(r)
		models.FindOne(filterParsed, point)
		respondJSON(w, http.StatusOK, point)
	}
}

func FindByArticleCode(point interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := getID(r)
		where := map[string]interface{}{
			"CodigoArticulo": id,
		}
		err := models.FindOneQuery(where, point)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			fmt.Println(err)

			return
		}
		respondJSON(w, http.StatusOK, point)
	}
}

func getID(r *http.Request) (string, error) {
	params := mux.Vars(r)
	id := params["id"]
	return id, nil
}

func getBody(point interface{}, r *http.Request) (interface{}, error) {
	body := point
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}
	return body, nil
}

func Patch(point interface{}, patch interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := getID(r)

		uid, _ := uuid.FromString(id)

		body, _ := getBody(patch, r)

		models.Patch(point, uid, body)

		fmt.Println(uid)
		fmt.Println(body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
