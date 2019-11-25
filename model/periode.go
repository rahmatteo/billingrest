package model

import (
	"belajargolang/billingrest/lib"
	"database/sql"
)

type Periode struct {
	Periode   string `json:"periode"`
	IDRoom    string `json:"id_room"`
	IDPeriode string `json:"id_periode"`
}

var TbPeriode = `
	CREATE TABLE tbperiode
	(
		id_periode serial primary key,
		periode varchar(50),
		id_room int
	);
`

//fungi deklarasi nama tabel
func (m *Periode) Name() string {
	return "tbperiode"
}

//show field
func (m *Periode) Field() (fields []string, dst []interface{}) {
	fields = []string{"id_periode", "periode", "id_room"}
	dst = []interface{}{&m.IDPeriode, &m.Periode, &m.IDRoom}
	return fields, dst
}

//inisialisai primary key kalo ada
func (m *Periode) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id_periode"}
	dst = []interface{}{&m.IDPeriode}
	return fields, dst
}

//conect table
func (m *Periode) Structur() lib.Table {
	return &Periode{}
}

// auto number
func (m *Periode) AutoNumber() bool {
	return true
}

//insert Periode
func (m *Periode) Insert(db *sql.DB) error {
	return lib.Insert(db, m)
}

func (m *Periode) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}

func (m *Periode) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}

func (m *Periode) Get(db *sql.DB) error {
	return lib.Get(db, m)
}

//ambil semua data Periode
func GetAllPeriode(db *sql.DB, params ...string) ([]*Periode, error) {
	m := &Periode{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	period := make([]*Periode, len(data))
	for index, item := range data {
		period[index] = item.(*Periode)
	}
	return period, nil
}
