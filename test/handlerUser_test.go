package test

import (
	"belajargolang/billingrest/handler"
	"belajargolang/billingrest/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestUserHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var dataInsertUser = []model.User{
		model.User{
			ID:          "1",
			Nama:        "rahmatteo",
			Username:    "Rahmat",
			Password:    "Rahmat",
			Status:      "1",
			Foto:        "Rahmat.jpg",
			LevelLogin:  "2",
			LatestLogin: time.Date(2019, 11, 11, 0, 0, 0, 0, time.UTC),
			// LatestLogin: time.Now(),
		},
	}

	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Testing Post", func(t *testing.T) {

		for _, item := range dataInsertUser {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/v1/billingrest/user/", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.User{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}
			compareUser(t, got, item)
		}
	})

	t.Run("testing gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/v1/billingrest/user", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.User{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		for index, item := range got {
			compareUser(t, item, dataInsertUser[index])
		}
	})

	t.Run("test gets with params", func(t *testing.T) {
		res := httptest.NewRecorder()
		// params := fmt.Sprintf("id,=,%s;username,=,%s", dataInsertUser[0].ID, dataInsertUser[0].Username)
		//params web handler
		//filter : contoh filter:<namafiled>,<operator>,<value> jika ada tambahan ";" contoh filter:asd,=,1;dsa,=,1
		//limit : digunakan untuk membatasi jumlah ambil data contoh limit:50
		//descending : jika true maka data akan di urutkan berdasarkan data yang terakhir masuk. atau data bernilai false
		//sort : sorting data berdasarkan field misal sort:asd ASC/sort: dsa DESC kalo lebih dari 1 : sort:asd ASC,dsa DESC

		// params := fmt.Sprintf("filter:nama,=,%v;username,=,%v;password,=,%v;status,=,%v;foto,=,%v;id,=,%v;level_login,=,%v;latest_login,=,%v;limit:10", dataInsertUser[0].Nama, dataInsertUser[0].Username, dataInsertUser[0].Password, dataInsertUser[0].Status, dataInsertUser[0].Foto, dataInsertUser[0].ID, dataInsertUser[0].LevelLogin, dataInsertUser[0].LatestLogin)
		params := fmt.Sprintf("filter:id,=,%v;limit:10", dataInsertUser[0].ID)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/user?%v", url.QueryEscape(params)), nil)

		// req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/simplerest/user?params=id_user,=,%s", dataInsertUser[0].ID), nil)
		if err != nil {

			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := []*model.User{}
		fmt.Println(got)
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got[0], dataInsertMahasiswa[0])
	})

	t.Run("testing gets 1 data", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/user/%s", dataInsertUser[0].ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.User{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		compareUser(t, got, dataInsertUser[0])

	})

	t.Run("testing Put", func(t *testing.T) {
		res := httptest.NewRecorder()
		dataUpdate := map[string]interface{}{
			"nama":         "patul",
			"username":     "patul",
			"password":     "patul",
			"status":       "patul",
			"foto":         "patul",
			"level_login":  "2",
			"latest_login": time.Now(),
		}
		jsonUpdate, err := json.MarshalIndent(dataUpdate, "", "")
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/billingrest/user/%s", dataInsertUser[0].ID), bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := model.User{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got, dataInsertMahasiswa)
	})

	t.Run("test Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/billingrest/user/%s", dataInsertUser[0].ID), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("ID tidak terhapus")
		}
	})

}
