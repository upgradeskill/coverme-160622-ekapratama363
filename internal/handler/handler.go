package handler

import (
	"encoding/json"
	"fmt"
	"main/internal/dto"
	"main/internal/storage"
	"net/http"
)

func responseJson(w http.ResponseWriter, message []byte, httpCode int) {
	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(httpCode)
	w.Write(message)
}

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			message := []byte(`{"message": "method harus GET"}`)
			responseJson(w, message, http.StatusMethodNotAllowed)
			return
		}

		users := storage.User

		fmt.Println(users)
		var data []dto.User

		for _, user := range users {
			data = append(data, user)
		}

		response, err := json.Marshal(data)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		responseJson(w, response, http.StatusOK)
	}
}

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			message := []byte(`{"message": "method harus POST"}`)
			responseJson(w, message, http.StatusMethodNotAllowed)
			return
		}

		var newUser dto.User

		payload := r.Body    //menampung request
		defer r.Body.Close() //close setelah menampung request

		err := json.NewDecoder(payload).Decode(&newUser)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		storage.User[newUser.ID] = newUser

		message := []byte(`{"message": "success create data"}`)
		responseJson(w, message, http.StatusCreated)
	}
}

func ShowUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			message := []byte(`{"message": "method harus GET"}`)
			responseJson(w, message, http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			message := []byte(`{"message": "id parameter harus diisi"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		data, isAvailable := storage.User[id]
		if !isAvailable {
			message := []byte(`{"message": "data tidak ditemukan"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(&data)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(w, message, http.StatusInternalServerError)
			return
		}

		responseJson(w, response, http.StatusOK)
	}
}

func UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			message := []byte(`{"message": "method harus PUT"}`)
			responseJson(w, message, http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			message := []byte(`{"message": "id parameter harus diisi"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		data, isAvailable := storage.User[id]
		if !isAvailable {
			message := []byte(`{"message": "data tidak ditemukan"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		var newUser dto.User
		payload := r.Body

		defer r.Body.Close()

		err := json.NewDecoder(payload).Decode(&newUser)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(w, message, http.StatusInternalServerError)
			return
		}

		data.Name = newUser.Name

		storage.User[data.ID] = data

		response, err := json.Marshal(&data)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(w, message, http.StatusInternalServerError)
			return
		}

		responseJson(w, response, http.StatusOK)
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			message := []byte(`{"message": "method harus DELETE"}`)
			responseJson(w, message, http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			message := []byte(`{"message": "id parameter harus diisi"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		_, isAvailable := storage.User[id]
		if !isAvailable {
			message := []byte(`{"message": "data tidak ditemukan"}`)
			responseJson(w, message, http.StatusBadRequest)
			return
		}

		delete(storage.User, id)

		message := []byte(`{"message": "berhasil hapus data"}`)
		responseJson(w, message, http.StatusOK)
	}
}
