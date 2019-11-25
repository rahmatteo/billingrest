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

func TestMemberHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var dataInsertMember = []model.Member{
		model.Member{
			Username:   "Rahmat",
			Password:   "Rahmat",
			Status:     "23",
			NamaMember: "1Rahmat",
			NoHp:       "8881281",
			Email:      "Rahmat@gmail.com",
			Alamat:     "Jl. menteng",
			Foto:       "Rahmat.jpg",
			IDMember:   "1",
			Flag:       "1",
			Nik:        "112321",
		},
	}

	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Testing Post", func(t *testing.T) {

		for _, item := range dataInsertMember {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/v1/billingrest/member/", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Member{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}
			compareMember(t, got, item)
		}
	})

	t.Run("testing gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/v1/billingrest/member", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Member{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		for index, item := range got {
			compareMember(t, item, dataInsertMember[index])
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

		params := fmt.Sprintf("filter:id_member,=,%v;limit:10", dataInsertMember[0].IDMember)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/member?%v", url.QueryEscape(params)), nil)

		// req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/simplerest/user?params=id_user,=,%s", dataInsertUser[0].ID), nil)
		if err != nil {

			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := []*model.Member{}
		fmt.Println(got)
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got[0], dataInsertMahasiswa[0])
	})

	t.Run("testing gets 1 data", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/billingrest/member/%s", dataInsertMember[0].IDMember), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.Member{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		compareMember(t, got, dataInsertMember[0])

	})

	t.Run("testing Put", func(t *testing.T) {
		res := httptest.NewRecorder()
		dataUpdate := map[string]interface{}{
			"username":    "patul",
			"password":    "patul",
			"status":      "1",
			"nama_member": "patul",
			"no_hp":       "0808",
			"email":       "patul@gmail.com",
			"alamat":      "jl. patul",
			"foto":        "patul.jpg",
			"flag":        "1",
			"nik":         "123123",
		}
		jsonUpdate, err := json.MarshalIndent(dataUpdate, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/billingrest/member/%s", dataInsertMember[0].IDMember), bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := model.Member{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareMahasiswa(t, got, dataInsertMahasiswa)
	})

	t.Run("test Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/billingrest/member/%s", dataInsertMember[0].IDMember), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("ID tidak terhapus")
		}
	})

}
