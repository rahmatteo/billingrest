package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
)

func TestInsertRoom(t *testing.T) {
	var dataInsertRoom = []model.Room{
		// db := PrepareTest(t)
		// defer db.Close()
		// data := []*model.User{
		model.Room{Description: "Desc Rahmat", Room: "Room Rahmat", Foto: "Rahmat.jpg", Price: "8000000"},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get Room", func(t *testing.T) {
		for _, dataInsert := range dataInsertRoom {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Room{IDRoom: dataInsert.IDRoom}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareRoom(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"description": "desc 1",
			"room":        "room 1",
			"foto":        "room.jpg",
			"price":       "2000000",
		}

		dataUpdate := model.Room{IDRoom: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllRoom(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Room{IDRoom: item.IDRoom}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareRoom(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "id_room,=,3"
		result, err := model.GetAllRoom(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Room{IDRoom: item.IDRoom}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareRoom(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Room{IDRoom: "2"}
		// m := model.User{ID: dataInsertUser[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func compareRoom(t *testing.T, got, want model.Room) {
	if got.IDRoom != want.IDRoom {
		t.Fatalf("got : %s want :%s id_room tidak sama", got.IDRoom, want.IDRoom)
	}
	if got.Description != want.Description {
		t.Fatalf("got :%s want :%s description tidak Sama", got.Description, want.Description)
	}
	if got.Room != want.Room {
		t.Fatalf("got :%s want :%s room tidak Sama", got.Room, want.Room)
	}
	if got.Foto != want.Foto {
		t.Fatalf("got :%s want :%s foto tidak Sama", got.Foto, want.Foto)
	}
	if got.Price != want.Price {
		t.Fatalf("got :%s want :%s price tidak Sama", got.Price, want.Price)
	}
}
