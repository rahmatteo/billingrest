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

func TestInvoiceHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var dataInsertInvoice = []model.Invoice{
		model.Invoice{
			IDInvoice:       "1",
			IDRoom:          "1",
			IDPeriode:       "1",
			Quantity:        "30",
			Charge:          "67888",
			Description:     "Rahmat",
			Total:           "5000000",
			TransactionDate: time.Date(2019, 11, 11, 0, 0, 0, 0, time.UTC),
			// TransactionDate: time.Now(),
		},
	}

	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Testing Post", func(t *testing.T) {

		for _, item := range dataInsertInvoice {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/v1/billingrest/invoice/", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Invoice{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, item)
		}
	})

	t.Run("testing gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/v1/billingrest/invoice", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Invoice{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		for index, item := range got {
			compareInvoice(t, item, dataInsertInvoice[index])
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

		params := fmt.Sprintf("filter:id_invoice,=,%v;limit:10", dataInsertInvoice[0].IDInvoice)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/invoice?%v", url.QueryEscape(params)), nil)

		// req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/simplerest/user?params=id_user,=,%s", dataInsertUser[0].ID), nil)
		if err != nil {

			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := []*model.Invoice{}
		fmt.Println(got)
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got[0], dataInsertMahasiswa[0])
	})

	t.Run("testing gets 1 data", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/invoice/%s", dataInsertInvoice[0].IDInvoice), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.Invoice{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		compareInvoice(t, got, dataInsertInvoice[0])

	})

	t.Run("testing Put", func(t *testing.T) {
		res := httptest.NewRecorder()
		dataUpdate := map[string]interface{}{
			"id_room":          "3",
			"id_periode":       "4",
			"quantity":         "5",
			"charge":           "6000000",
			"description":      "desc1323",
			"total":            "67000",
			"transaction_date": time.Now(),
			// "transaction_date": "2019-11-20",
		}
		jsonUpdate, err := json.MarshalIndent(dataUpdate, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/billingrest/invoice/%s", dataInsertInvoice[0].IDInvoice), bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := model.Invoice{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got, dataInsertMahasiswa)
	})

	t.Run("test Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/billingrest/invoice/%s", dataInsertInvoice[0].IDInvoice), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("ID tidak terhapus")
		}
	})

}
