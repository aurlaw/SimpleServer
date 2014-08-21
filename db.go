package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Bays struct {
	BayID   int
	BayName string
}

func CreateCustomer(custNum string, email string) (int64, error) {
	db, err := sql.Open("mysql", "MKUser:!gigahertz@tcp(vagrant-dev:3306)/mk_clev?autocommit=true")
	if err != nil {
		// panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		log.Println(err)
		return 0, err
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO go_test VALUES(?)") // ? = placeholder
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmtIns.Close()
	log.Println("Statement prepped")
	res, err := stmtIns.Exec(custNum) // Insert tuples (i, i^2)
	if err != nil {
		log.Println("Statement error")
		log.Println(err)
		return 0, err
	}
	log.Println("Statement exec")
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Println(rowCnt)
	return rowCnt, nil
}

/*
res, err := stmt.Exec("Dolly")
if err != nil {
    log.Fatal(err)
}
lastId, err := res.LastInsertId()
if err != nil {
    log.Fatal(err)
}
rowCnt, err := res.RowsAffected()
if err != nil {
    log.Fatal(err)
}

*/

func GetAllBays() (error, []Bays) {

	db, err := sql.Open("mysql", "MKUser:!gigahertz@tcp(vagrant-dev:3306)/mk_clev")
	if err != nil {
		// panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		log.Println(err)
		return err, nil
	}
	defer db.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT BayID, BayName FROM bays")
	if err != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
		log.Println(err)
		return err, nil
	}
	defer stmtOut.Close()

	bayList := make([]Bays, 0)
	rows, err := stmtOut.Query()
	if err != nil {
		log.Println(err)
		return err, nil
	}
	defer rows.Close()
	var (
		id   int
		name string
	)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
			return err, nil
		}
		b := Bays{BayID: id, BayName: name}
		bayList = append(bayList, b)
	}

	return nil, bayList
}

/*



   <add name="MySQL" connectionString="Server=vagrant-dev;Port=3306;Uid=MKUser;Pwd=!gigahertz;Database=mk_clev;" providerName="MySql.Data.MySqlClient" />


*/
