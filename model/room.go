package model

import (
	"belajargolang/billingrest/lib"
	"database/sql"
)

type Room struct {
	IDRoom      string `json:"id_room"`
	Description string `json:"description"`
	Room        string `json:"room"`
	Foto        string `json:"foto"`
	Price       string `json:"price"`
}

var TbRoom = `
	CREATE TABLE tbroom
	(
		id_room serial primary key,
		room varchar(255),
		foto varchar(30),
		description varchar(225),
		price int
	);
`

//fungi deklarasi nama tabel
func (m *Room) Name() string {
	return "tbroom"
}

//show field
func (m *Room) Field() (fields []string, dst []interface{}) {
	fields = []string{"id_room", "room", "foto", "description", "price"}
	dst = []interface{}{&m.IDRoom, &m.Room, &m.Foto, &m.Description, &m.Price}
	return fields, dst
}

//inisialisai primary key kalo ada
func (m *Room) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id_room"}
	dst = []interface{}{&m.IDRoom}
	return fields, dst
}

//conect table
func (m *Room) Structur() lib.Table {
	return &Room{}
}

// auto number
func (m *Room) AutoNumber() bool {
	return true
}

//insert Room
func (m *Room) Insert(db *sql.DB) error {
	return lib.Insert(db, m)
}

func (m *Room) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}

func (m *Room) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}

func (m *Room) Get(db *sql.DB) error {
	return lib.Get(db, m)
}

//ambil semua data Room
func GetAllRoom(db *sql.DB, params ...string) ([]*Room, error) {
	m := &Room{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	rm := make([]*Room, len(data))
	for index, item := range data {
		rm[index] = item.(*Room)
	}
	return rm, nil
}
