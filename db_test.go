package main

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	fmt.Println("TestDB")
	err, list := GetAllBays()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("list:", len(list))
	for i, v := range list {
		fmt.Printf("array value at [%d]=%d:%s\n\r", i, v.BayID, v.BayName)

	}

	if len(list) == 0 {
		t.Error("expected the list to contain item")
	}

}

func TestInsert(t *testing.T) {
	fmt.Println("TestInsert")

	custNum := "912e222223"
	email := "test@user.com"
	_, err := CreateCustomer(custNum, email)
	if err != nil {
		t.Error(err)
	}
	// if rowCount == 0 {
	// 	t.Error("expected rowCount to be > 0")
	// }

	// fmt.Printf("Row Count: %v\n\r", rowCount)
}
