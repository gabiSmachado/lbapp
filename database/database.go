package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gabiSmachado/lbapp/datamodel"
	_ "github.com/go-sql-driver/mysql"
)

func CurrentId(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT MAX(id) FROM intents").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS intents (
		id int primary key AUTO_INCREMENT,
		name varchar(255))`)

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}

	log.Printf("Creating table successfully!\n")
	return nil
}

func Insert(db *sql.DB, intent datamodel.Intent) (int, error) {
	_, err := db.Exec(`INSERT INTO intents (name) VALUES (?)`, intent.Name)
	if err != nil {
		log.Printf("Error %s when inserting in table", err)
		return 0, err
	}
	id, _ := CurrentId(db)
	return id, nil
}

func ListIntents(db *sql.DB) ([]datamodel.Intent, error) {
	rows, err := db.Query("SELECT * FROM intents")
	if err != nil {
		log.Printf("Error %s when listing intents", err)
		return nil, err
	}
	defer rows.Close()

	var intents []datamodel.Intent
	var name string
	var id int
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Printf("Error %s when listing intents", err)
		}
		intent := datamodel.Intent{Idx: id, Name: name}
		intents = append(intents, intent)
	}
	if err = rows.Err(); err != nil {
		//log.Printf("Error %s when listing intents", err)
		return nil, err
	}

	return intents, nil

}

func DeleteIntent(db *sql.DB, id int) error {
	_, err := db.Query("DELETE FROM intents WHERE id = ?", id)
	return err
}

func DBconnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/intents")
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		fmt.Println("Connection successful!")
		return db, nil
	}

}

func IntentShow(db *sql.DB, id int) (datamodel.Intent, error) {
	var name string
	err := db.QueryRow("SELECT * FROM intents WHERE id = ?", id).Scan(&id, &name)
	intent := datamodel.Intent{Idx: id, Name: name}

	if err != nil {
		log.Printf("Error %s when selecting intent", err)
		return intent, err
	}
	return intent, nil
}

/* func main() {
	db, err := DBconnect()
	if err != nil {
		println(err)
	}
	defer db.Close()

	createTable(db)

	//Insert(db, intent)

	//ShowIntent(db, 1)
	//ListIntents(db)
	//DeleteIntent(db, 1)
}
*/
