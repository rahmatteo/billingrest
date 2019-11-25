package handler

import (
	"belajargolang/billingrest/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandlerUserPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.User
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func HandlerUserDelete(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	data := model.User{ID: lastIndex}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("true"))
}

func HandlerUserPut(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonmap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := model.User{ID: lastIndex}
	err = data.Update(db, jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := data.Get(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func HandlerUserGet(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	if lastIndex == "user" {

		//query dijadiin params
		params, _ := r.URL.Query()["params"]

		data, err := model.GetAllUser(db, params...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		data := model.User{ID: lastIndex}

		err := data.Get(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonData)
	}

}
