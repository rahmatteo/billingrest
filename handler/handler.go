package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

var db *sql.DB

func RegisDB(sqlDB *sql.DB) {
	if db != nil {
		panic("db telah terdaftar")
	}
	db = sqlDB
}

func LastIndex(r *http.Request) string {
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	// fmt.Sprintf("ini data url %v", dataURL)
	lastIndex := dataURL[len(dataURL)-1]
	return lastIndex
}

func SS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/htmll; charset=utf-8; application/json")
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	fmt.Println(dataURL[3])
	switch dataURL[3] {

	case "user":
		switch r.Method {
		case http.MethodGet:
			HandlerUserGet(w, r)
		case http.MethodPost:
			HandlerUserPost(w, r)
		case http.MethodPut:
			HandlerUserPut(w, r)
		case http.MethodDelete:
			HandlerUserDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "periode":
		switch r.Method {
		case http.MethodGet:
			HandlerPeriodeGet(w, r)
		case http.MethodPost:
			HandlerPeriodePost(w, r)
		case http.MethodPut:
			HandlerPeriodePut(w, r)
		case http.MethodDelete:
			HandlerPeriodeDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "invoice":
		switch r.Method {
		case http.MethodGet:
			HandlerInvoiceGet(w, r)
		case http.MethodPost:
			HandlerInvoicePost(w, r)
		case http.MethodPut:
			HandlerInvoicePut(w, r)
		case http.MethodDelete:
			HandlerInvoiceDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "member":
		switch r.Method {
		case http.MethodGet:
			HandlerMemberGet(w, r)
		case http.MethodPost:
			HandlerMemberPost(w, r)
		case http.MethodPut:
			HandlerMemberPut(w, r)
		case http.MethodDelete:
			HandlerMemberDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "room":
		switch r.Method {
		case http.MethodGet:
			HandlerRoomGet(w, r)
		case http.MethodPost:
			HandlerRoomPost(w, r)
		case http.MethodPut:
			HandlerRoomPut(w, r)
		case http.MethodDelete:
			HandlerRoomDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	// case "nilai":
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		HandlerNilaiGet(w, r)
	// 	case http.MethodPost:
	// 		HandlerNilaiPost(w, r)
	// 	case http.MethodPut:
	// 		HandlerNilaiPut(w, r)
	// 	case http.MethodDelete:
	// 		HandlerNilaiDelete(w, r)
	// 	default:
	// 		w.Write([]byte("method tidak ditemukan"))

	// 	}
	// case "matakuliah":
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		HandlerMatakuliahGet(w, r)
	// 	case http.MethodPost:
	// 		HandlerMatakuliahPost(w, r)
	// 	case http.MethodPut:
	// 		HandlerMatakuliahPut(w, r)
	// 	case http.MethodDelete:
	// 		HandlerMatakuliahDelete(w, r)
	// 	default:
	// 		w.Write([]byte("method tidak ditemukan"))

	// 	}
	default:
		w.Write([]byte("request not found"))
	}
}

// package handler

// import (
// 	"belajargolang/simplerest/model"
// 	"database/sql"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"net/url"
// 	"strconv"
// 	"strings"
// )

// type SimpleRestHandler struct {
// 	Pattern string
// }

// func UrlPath(u *url.URL, pattern string) []string {
// 	urlpath := u.RawPath
// 	if urlpath == "/v1/simplerest/" {
// 		urlpath = u.Path
// 	}
// 	pathpattern := strings.TrimPrefix(urlpath, pattern)
// 	path := strings.Split(pathpattern, "/")
// 	return path
// }

// //Query Fields
// func QueryFields(query url.Values) ([]string, error) {
// 	qry := query.Get("fields")
// 	var flds []string
// 	if qry != "" {
// 		flds = strings.Split(qry, "")
// 	}
// 	return flds, nil
// }

// //Query Cursor
// func QueryCursor(query url.Values) (set bool, cursor dba.Cursor, err error) {
// 	qry := query.Get("cursor")
// 	if qry != "" {
// 		cursor, err = dba.Decode(qry)
// 		set = true
// 	}

// 	return set, cursor, err
// }

// //Query Limit
// func QueryLimit(query url.Values) (int, error) {
// 	qry := query.Get("limit")
// 	var lmt int
// 	var err error

// 	if qry != "" {
// 		lmt, err = strconv.Atoi(qry)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	return lmt, nil
// }

// //Query Sort
// func QuerySort(query url.Values) ([]string, bool, error) {
// 	qry := query.Get("sort")
// 	var srt []string
// 	descending := true
// 	if qry != "" {
// 		fieldsrt := strings.Split(qry, "")
// 		srt = strings.Split(fieldsrt[0], ",")
// 		if len(srt) < 1 {
// 			return nil, false, fmt.Errorf("tidak memenuhi persyaratan")
// 		}
// 		if fieldsrt[1] != "DESC" {
// 			descending = false
// 		}
// 	}
// 	return srt, descending, nil
// }

// func QueryFilter(query url.Values) ([]dba.Filter, error) {
// 	qry := query.Get("Filters")

// 	var fil []dba.filter
// 	if qry != "" {
// 		fieldspr := strings.Split(qry, ";")

// 		for _, tc = range fieldspr {
// 			parameter := strings.Split(tc, ",")
// 			if len(parameter) != 3 {
// 				return nil, fmt.Errorf("tidak memenuhi persyaratan")
// 			}
// 			valueparameter, err := url.PathUnEscape(parameter[2])
// 			if err != nil {
// 				return nil, fmt.Errorf("tidak memenuhi persyaratan")
// 			}
// 			b := dba.Filter{
// 				Field: parameter[0],
// 				Op:    parameter[1],
// 				Value: valueparameter,
// 			}

// 			fil = append(fil, b)
// 		}
// 	}

// 	return fil, nil
// }

// func newTx(db *sql.DB) *sql.Tx {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return nil
// 	}
// 	return tx
// }

// func WebHandler(pattern string) http.Handler {
// 	wh := webhandler.new(SimpleRestHandler{Pattern: pattern})
// 	return wh
// }

// func (h SimpleRestHandler) Handle(w http.ResponseWriter, r *http.Request) webHandler.Response {
// 	var res webhandler.Response

// 	path := strings.TrimPrefix(r.URL.Path, "")
// 	splitPath := strings.Split(path, "/")
// 	urls := splitPath[0]
// 	switch {
// 	case urls == "User":
// 		switch {
// 		case r.Method == http.MethodGet:
// 			res = h.handleGet(w, r)

// 		case r.Method == http.MethodPost:
// 			res = h.handlePost(w, r)

// 		case r.Method == http.MethodPut:
// 			res = h.handlePut(w, r)

// 		case r.Method == http.MethodDelete:
// 			res = h.handleDelete(w, r)
// 		}
// 	default:
// 		res.ErrMessage = "Not Found"
// 		return res
// 	}
// 	return res
// }

// func (h SimpleRestHandler) handleGet(w http.ResponseWriter, r *http.Request) webHandler.Response {
// 	res := webhandler.Response{}
// 	paths := UrlPath(r.URL, h.pattern)

// 	db, err := webhandler.DBFromContext(r.Context())
// 	if err != nil {
// 		res.Error(err, http.StatusIntervalServerErorr)
// 		return res
// 	}

// 	if len(paths) == 2 {
// 		id, err := url.PathUnEscape(paths[1])
// 		if err != nil {
// 			res.Error(err, http.StatusInternalServerError)
// 			return res
// 		}

// 		if id != "" {
// 			m := &model.User{id_user: id}
// 			err := m.Get(db)
// 			if err != nil {
// 				res.Error(err, http.StatusOK)
// 				return res
// 			}
// 			res.Data = m
// 			return res
// 		}
// 	}

// 	var values = r.URL.Query()
// 	set, cursor, err := QueryCursor(values)
// 	if err != nil {
// 		res.Error(err, http.StatusBadRequest)
// 	}
// 	if !set {
// 		fields, err := QueryFields(values)
// 		if err != nil {
// 			res.Error(err, http.StatusBadRequest)
// 			return res
// 		}
// 		limit, err := QueryLimit(values)
// 		if err != nil {
// 			res.Error(err, http.StatusBadRequest)
// 			return res
// 		}
// 		if limit == 0 {
// 			limit = 1000
// 		}
// 		sort, asc, err := QuerySort(values)
// 		if err != nil {
// 			res.Error(err, http.StatusBadRequest)
// 			return res
// 		}
// 		filter, err := QueryFilter(values)
// 		if err != nil {
// 			res.Error(err, http.StatusBadRequest)
// 			return res
// 		}
// 		cursor = dba.Cursor{
// 			Fields:     fields,
// 			Filters:    filter,
// 			OrderBy:    sort,
// 			Descending: asc,
// 			Limit:      limit,
// 		}
// 	}

// 	var data []dba.Table
// 	m := &model.User{}
// 	data, cursor, err = dba.Fetch(db, m, "", cursor)
// 	if err != nil {
// 		res.Error(err, http.StatusOK)
// 		return res
// 	}
// 	res.Data = data
// 	res.Cursor = cursor
// 	return res
// }

// func (h SimpleRestHandler) handlePost(w http.ResponseWriter, r *http.Request) webHandler.Response {
// 	res := webhandler.Response{}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		res.Error(err, StatusBadRequest)
// 		return res
// 	}

// 	user := &model.User{}
// 	if err := json.Unmarshal(body, mahasiswa); err != nil {
// 		res.Error(err, http.StatusBadRequest)
// 		return res
// 	}

// 	db, err := webhandler.DBFromContext(r.Context())
// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	tx, err := db.Begin()
// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}
// 	if err = user.Insert(tx); err != nil {
// 		res.Error(err, http.StatusOK)
// 		tx.Rollback()
// 		return res
// 	}
// 	if err = tx.Commit(); err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	res.Data = user
// 	return res
// }

// func (h SimpleRestHandler) handlePut(w http.ResponseWriter, r *http.Request) webHandler.Response {
// 	res := webhandler.Response{}
// 	body, err := ioutil.ReadAll(r.body)
// 	if err != nil {
// 		res.Error(err, http.StatusBadRequest)
// 		return res
// 	}
// 	data := make(map[string]interface{})
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		res.Error(err, http.StatusBadRequest)
// 		return res
// 	}

// 	paths := UrlPath(r.URL, h.Pattern)
// 	last, err := url.PathUnEscape(paths[1])
// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}
// 	if last == "" || last == "user" {
// 		res.Error(errors.New("id tidak boleh kosong"), http.StatusOK)
// 		return res
// 	}
// 	tx, err := db.Begin()
// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	user := model.User{id_user: last}
// 	if err = user.Get(db); err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	data, err = user.Update(tx, data)
// 	if err != nil {
// 		tx.Rollback()
// 		res.Error(err, http.StatusOK)
// 		return res
// 	}

// 	if err = tx.Commit(); err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	res.Data = data
// 	return res
// }

// func (h SimpleRestHandler) handleDelete(w http.ResponseWriter, r *http.Request) webHandler.Response {
// 	res := webhandler.Response{}
// 	paths := UrlPath(r.URL, h.pattern)
// 	last, err := url.PathUnescape(paths[1])
// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}
// 	if last == "" || last == "user" {
// 		res.Error(errors.New("id user tidak boleh kosong"), http.StatusOK)
// 		return res
// 	}

// 	db, err := webhandler.DBFromContext(r.Context())
// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	tx, err := db.Begin()

// 	if err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	user := model.User{id_user: last}
// 	if err = user.Get(db); err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	data, err = user.Update(tx, data)
// 	if err != nil {
// 		tx.Rollback()
// 		res.Error(err, http.StatusOK)
// 		return res
// 	}

// 	if err = tx.Commit(); err != nil {
// 		res.Error(err, http.StatusInternalServerError)
// 		return res
// 	}

// 	res.Data = data
// 	return res

// }
