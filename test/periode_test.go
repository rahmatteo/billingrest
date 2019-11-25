package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
)

func TestInsertPeriode(t *testing.T) {
	var dataInsertPeriode = []model.Periode{
		// db := PrepareTest(t)
		// defer db.Close()
		// data := []*model.User{
		model.Periode{Periode: "periode 1", IDRoom: "11"},
		model.Periode{Periode: "periode 2", IDRoom: "12"},
		model.Periode{Periode: "periode 3", IDRoom: "13"},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get Periode", func(t *testing.T) {
		for _, dataInsert := range dataInsertPeriode {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Periode{IDPeriode: dataInsert.IDPeriode}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"periode": "20",
			"id_room": "5",
		}

		dataUpdate := model.Periode{IDPeriode: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllPeriode(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Periode{IDPeriode: item.IDPeriode}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "periode,=,DIA"
		result, err := model.GetAllPeriode(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Periode{IDPeriode: item.IDPeriode}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Periode{IDPeriode: "2"}
		// m := model.Periode{ID: dataInsertPeriode[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func comparePeriode(t *testing.T, got, want model.Periode) {
	if got.Periode != want.Periode {
		t.Fatalf("got : %s want :%s periode tidak sama", got.Periode, want.Periode)
	}
	if got.IDRoom != want.IDRoom {
		t.Fatalf("got : %s want :%s id_room tidak sama", got.IDRoom, want.IDRoom)
	}

}
