package model

import (
	"belajargolang/billingrest/lib"
	"database/sql"
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Nama        string    `json:"nama"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Status      string    `json:"status"`
	Foto        string    `json:"foto"`
	LevelLogin  string    `json:"level_login"`
	LatestLogin time.Time `json:"latest_login"`
}

var TbUser = `
	CREATE TABLE tbuser
	(
		id serial primary key,
		nama varchar(50),
		username varchar(50),
		password varchar(30),
		status varchar(30),
		foto varchar(225),
		level_login int,
		latest_login DATE
	);
`

//fungi deklarasi nama tabel
func (m *User) Name() string {
	return "tbuser"
}

//show field
func (m *User) Field() (fields []string, dst []interface{}) {
	fields = []string{"id", "nama", "username", "password", "status", "foto", "level_login", "latest_login"}
	dst = []interface{}{&m.ID, &m.Nama, &m.Username, &m.Password, &m.Status, &m.Foto, &m.LevelLogin, &m.LatestLogin}
	return fields, dst
}

//inisialisai primary key kalo ada
func (m *User) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id"}
	dst = []interface{}{&m.ID}
	return fields, dst
}

//conect table
func (m *User) Structur() lib.Table {
	return &User{}
}

// auto number
func (m *User) AutoNumber() bool {
	return true
}

//insert user
func (m *User) Insert(db *sql.DB) error {

	return lib.Insert(db, m)
}

func (m *User) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}

func (m *User) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}

func (m *User) Get(db *sql.DB) error {
	return lib.Get(db, m)
}

//ambil semua data User
func GetAllUser(db *sql.DB, params ...string) ([]*User, error) {
	m := &User{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	user := make([]*User, len(data))
	for index, item := range data {
		user[index] = item.(*User)
	}
	return user, nil
}
