package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
	"time"
)

func TestInsertUser(t *testing.T) {
	var dataInsertUser = []model.User{
		// db := PrepareTest(t)
		// defer db.Close()
		// data := []*model.User{
		model.User{Nama: "Rahmat", Username: "Rahmat", Password: "Rahmat", Status: "active", Foto: "Rahmat.jpg", LevelLogin: "3", LatestLogin: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC)},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get User", func(t *testing.T) {
		for _, dataInsert := range dataInsertUser {

			err = dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.User{ID: dataInsert.ID}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareUser(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"nama":         "patuyl",
			"username":     "patuy",
			"password":     "patuy",
			"status":       "stat patuy",
			"foto":         "patuy.jpg",
			"level_login":  "2",
			"latest_login": time.Now(),
		}

		dataUpdate := model.User{ID: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

		// dataUpdate := dataInsertUser[0]
		// if err := dataUpdate.Update(db, update); err != nil {
		// 	t.Fatal(err)
		// }

		// got := model.User{ID: dataUpdate.ID}
		// if err := got.Get(db); err != nil {
		// 	t.Fatal(err)
		// }
		// compareUser(t, got, dataUpdate)

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllUser(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.User{ID: item.ID}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareUser(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "username,=,DIA"
		result, err := model.GetAllUser(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.User{ID: item.ID}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareUser(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.User{ID: "2"}
		// m := model.User{ID: dataInsertUser[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func compareUser(t *testing.T, got, want model.User) {
	if got.Nama != want.Nama {
		t.Fatalf("got : %s want :%s nama tidak sama", got.Nama, want.Nama)
	}
	if got.Username != want.Username {
		t.Fatalf("got :%s want :%s username tidak Sama", got.Username, want.Username)
	}
	if got.Password != want.Password {
		t.Fatalf("got :%s want :%s password tidak Sama", got.Password, want.Password)
	}
	if got.Status != want.Status {
		t.Fatalf("got :%s want :%s status tidak Sama", got.Status, want.Status)
	}
	if got.Foto != want.Foto {
		t.Fatalf("got :%s want :%s foto tidak Sama", got.Foto, want.Foto)
	}
	if got.ID != want.ID {
		t.Fatalf("got : %s want :%s id tidak sama", got.ID, want.ID)
	}
	if got.LevelLogin != want.LevelLogin {
		t.Fatalf("got : %v want :%v lavel_login tidak sama", got.LevelLogin, want.LevelLogin)
	}
	// if got.LatestLogin != want.LatestLogin {
	// 	t.Fatalf("got : %s want :%s latest_login tidak sama", got.LatestLogin, want.LatestLogin)
	// }

	// if len(fields) == 0 {
	// 	fields, _ = got.Fields()
	// }
	// for _, field := range fields {
	// 	if field == "id_user" && got.ID != want.ID {
	// 		t.Errorf("Got ID : %v want : %v", got.ID, want.ID)
	// 	}
	// 	if field == "username" && got.Username != want.Username {
	// 		t.Errorf("Got Username : %v want : %v", got.Username, want.Username)
	// 	}
	// 	if field == "password" && got.Password != want.Password {
	// 		t.Errorf("Got Password : %v want : %v", got.Password, want.Username)
	// 	}
	// 	if field == "repassword" && got.Repassword != want.Repassword {
	// 		t.Errorf("Got Repassword : %v want : %v", got.Repassword, want.Repassword)
	// 	}
	// 	if field == "email" && got.Email != want.Email {
	// 		t.Errorf("Got Email : %v want : %v", got.Email, want.Email)
	// 	}
	// }
}

// package test

// import (
// 	"belajargolang/simplerest/model"
// 	"testing"
// )

// func TestUser(t *testing.T) {
// 	db := PrepareTest(t)
// 	defer db.Close()
// 	data := []*model.User{
// 		&model.User{ID: "55415572", Username: "Rahmat", Password: "Rahmat", Repassword: "Rahmat", Email: "Rahmat@gmail.com"},
// 		&model.User{ID: "55415572", Username: "Saputra", Password: "Saputra", Repassword: "Saputra", Email: "Saputra@gmail.com"},
// 		&model.User{ID: "55415572", Username: "Matteo", Password: "Matteo", Repassword: "Matteo", Email: "Matteo@gmail.com"},
// 	}
// 	t.Run("Test Insert User", func(t *testing.T) {
// 		for _, item := range data {
// 			fields, _ := item.Fields()
// 			tx, err := db.Begin()
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			if err := item.Insert(tx); err != nil {
// 				t.Fatalf("Insert data : %v error : %v", item, err)
// 			}
// 			if err := tx.Commit(); err != nil {
// 				t.Fatal(err)
// 			}
// 			got := &model.User{ID: item.ID}
// 			if err := got.Get(db); err != nil {
// 				t.Fatal(err)
// 			}
// 			CompareUser(t, got, item, fields)
// 		}
// 	})

// 	t.Run("Test Update User", func(t *testing.T) {
// 		changes := []map[string]interface{}{
// 			{"nama": "Verudi"},
// 		}
// 		dataUpdate := data[0]
// 		for _, change := range changes {
// 			_, err := dataUpdate.Update(db, change)
// 			if err != nil {
// 				t.Fatalf("Update error : %v", err)
// 			}
// 			got := &model.User{ID: data[0].ID}
// 			if err := got.Get(db); err != nil {
// 				t.Fatal(err)
// 			}
// 		}
// 	})

// 	t.Run("Test Delete User", func(t *testing.T) {
// 		for _, item := range data {
// 			if err := item.Delete(db); err != nil {
// 				t.Errorf("Delete id : %v error : %v", item.ID, err)
// 			}
// 		}
// 	})
// }

// func CompareUser(t *testing.T, got, want *model.User, fields []string) {
// 	if len(fields) == 0 {
// 		fields, _ = got.Fields()
// 	}
// 	for _, field := range fields {
// 		if field == "id_user" && got.ID != want.ID {
// 			t.Errorf("Got ID : %v want : %v", got.ID, want.ID)
// 		}
// 		if field == "username" && got.Username != want.Username {
// 			t.Errorf("Got Username : %v want : %v", got.Username, want.Username)
// 		}
// 		if field == "password" && got.Password != want.Password {
// 			t.Errorf("Got Password : %v want : %v", got.Password, want.Username)
// 		}
// 		if field == "repassword" && got.Repassword != want.Repassword {
// 			t.Errorf("Got Repassword : %v want : %v", got.Repassword, want.Repassword)
// 		}
// 		if field == "email" && got.Email != want.Email {
// 			t.Errorf("Got Email : %v want : %v", got.Email, want.Email)
// 		}
// 	}
// }
