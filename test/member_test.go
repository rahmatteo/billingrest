package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
)

func TestInsertMember(t *testing.T) {
	var dataInsertMember = []model.Member{
		// db := PrepareTest(t)
		// defer db.Close()
		// data := []*model.User{
		model.Member{Username: "Rahmat", Password: "Rahmat", Status: "2", NamaMember: "Rahmat@gmail.com", NoHp: "80890", Email: "Rahmat@gmail.com", Alamat: "jl rahmat", Foto: "rahmat.jpg", IDMember: "4", Flag: "1", Nik: "3123213"},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get Member", func(t *testing.T) {
		for _, dataInsert := range dataInsertMember {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Member{IDMember: dataInsert.IDMember}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareMember(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"username": "patuy",
		}

		dataUpdate := model.Member{IDMember: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllMember(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Member{IDMember: item.IDMember}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareMember(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "username,=,DIA"
		result, err := model.GetAllMember(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Member{IDMember: item.IDMember}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareMember(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Member{IDMember: "2"}
		// m := model.User{ID: dataInsertUser[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func compareMember(t *testing.T, got, want model.Member) {
	if got.IDMember != want.IDMember {
		t.Fatalf("got : %s want :%s id_member tidak sama", got.IDMember, want.IDMember)
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
	if got.NamaMember != want.NamaMember {
		t.Fatalf("got :%s want :%s nama_member tidak Sama", got.NamaMember, want.NamaMember)
	}
	if got.NoHp != want.NoHp {
		t.Fatalf("got :%s want :%s no_hp tidak Sama", got.NoHp, want.NoHp)
	}
	if got.Email != want.Email {
		t.Fatalf("got :%s want :%s email tidak Sama", got.Email, want.Email)
	}
	if got.Alamat != want.Alamat {
		t.Fatalf("got :%s want :%s alamat tidak Sama", got.Alamat, want.Alamat)
	}
	if got.Foto != want.Foto {
		t.Fatalf("got :%s want :%s foto tidak Sama", got.Foto, want.Foto)
	}
	if got.Flag != want.Flag {
		t.Fatalf("got :%s want :%s flag tidak Sama", got.Flag, want.Flag)
	}
	if got.Nik != want.Nik {
		t.Fatalf("got :%s want :%s nik tidak Sama", got.Nik, want.Nik)
	}
}
