package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
	"time"
)

func TestInsertInvoice(t *testing.T) {
	var dataInsertInvoice = []model.Invoice{
		// db := PrepareTest(t)
		// defer db.Close()
		// data := []*model.User{
		model.Invoice{IDRoom: "2", IDPeriode: "3", Quantity: "30", Charge: "10000", Description: "urutan invoice pertama", Total: "20000000", TransactionDate: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC)},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get Invoice", func(t *testing.T) {
		for _, dataInsert := range dataInsertInvoice {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Invoice{IDInvoice: dataInsert.IDInvoice}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"id_room":          "2",
			"id_periode":       "5",
			"quantity":         "40",
			"charge":           "700000",
			"description":      "keterangan update test",
			"total":            "80000000",
			"transaction_date": time.Now(),
		}

		dataUpdate := model.Invoice{IDInvoice: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllInvoice(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Invoice{IDInvoice: item.IDInvoice}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "id_room,=,2"
		result, err := model.GetAllInvoice(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Invoice{IDInvoice: item.IDInvoice}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Invoice{IDInvoice: "2"}
		// m := model.User{ID: dataInsertUser[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func compareInvoice(t *testing.T, got, want model.Invoice) {
	if got.IDInvoice != want.IDInvoice {
		t.Fatalf("got : %s want :%s id_invoice tidak sama", got.IDInvoice, want.IDInvoice)
	}
	if got.IDRoom != want.IDRoom {
		t.Fatalf("got :%s want :%s id_room tidak Sama", got.IDRoom, want.IDRoom)
	}
	if got.Quantity != want.Quantity {
		t.Fatalf("got :%s want :%s quantity tidak Sama", got.Quantity, want.Quantity)
	}
	if got.Charge != want.Charge {
		t.Fatalf("got :%s want :%s charge tidak Sama", got.Charge, want.Charge)
	}
	if got.Description != want.Description {
		t.Fatalf("got :%s want :%s description tidak Sama", got.Description, want.Description)
	}
	if got.Total != want.Total {
		t.Fatalf("got :%s want :%s total tidak Sama", got.Total, want.Total)
	}
	// if got.TransactionDate != want.TransactionDate {
	// 	t.Fatalf("got :%v want :%v transaction_date tidak Sama", got.TransactionDate, want.TransactionDate)
	// }
	if got.IDPeriode != want.IDPeriode {
		t.Fatalf("got :%s want :%s id_periode tidak Sama", got.IDPeriode, want.IDPeriode)
	}
}
