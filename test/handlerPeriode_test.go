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
)

func TestPeriodeHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var dataInsertPeriode = []model.Periode{
		model.Periode{
			Periode:   "Rahmat",
			IDRoom:    "3",
			IDPeriode: "1",
		},
	}

	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Testing Post", func(t *testing.T) {

		for _, item := range dataInsertPeriode {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/v1/billingrest/periode/", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Periode{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, item)
		}
	})

	t.Run("testing gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/v1/billingrest/periode", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Periode{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		for index, item := range got {
			comparePeriode(t, item, dataInsertPeriode[index])
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

		// params := fmt.Sprintf("filter:periode,=,%v;id_room,=,%v;id_periode,=,%v;limit:10", dataInsertPeriode[0].Periode, dataInsertPeriode[0].IDRoom, dataInsertPeriode[0].IDPeriode)
		params := fmt.Sprintf("filter:id_periode,=,%v;limit:10", dataInsertPeriode[0].IDPeriode)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/periode?%v", url.QueryEscape(params)), nil)

		// req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/simplerest/user?params=id_user,=,%s", dataInsertUser[0].ID), nil)
		if err != nil {

			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := []*model.Periode{}
		fmt.Println(got)
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got[0], dataInsertMahasiswa[0])
	})

	t.Run("testing gets 1 data", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/periode/%s", dataInsertPeriode[0].IDPeriode), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.Periode{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		comparePeriode(t, got, dataInsertPeriode[0])

	})

	t.Run("testing Put", func(t *testing.T) {
		res := httptest.NewRecorder()
		dataUpdate := map[string]interface{}{
			"periode": "periode 1",
			"id_room": "3",
		}
		jsonUpdate, err := json.MarshalIndent(dataUpdate, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/billingrest/periode/%s", dataInsertPeriode[0].IDPeriode), bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := model.Periode{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got, dataInsertMahasiswa)
	})

	t.Run("test Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/billingrest/periode/%s", dataInsertPeriode[0].IDPeriode), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("ID tidak terhapus")
		}
	})

}
