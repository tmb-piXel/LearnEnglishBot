package repositories

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	m "github.com/tmb-piXel/LearnEnglishBot/pkg/models"
)

// var users = make(map[int64]*m.User)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// create connection with postgres db
func createConnection() *sql.DB {
	POSTGRES_URL := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Open the connection
	db, err := sql.Open("postgres", POSTGRES_URL)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

func FindUser(id int64) *m.User {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var u m.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE chatid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	var (
		chatID   int64
		language string
		topic    string
		isToRu   bool
		iterWord int
	)
	// unmarshal the row object to user
	err := row.Scan(&chatID, &language, &topic, &isToRu, &iterWord)

	u.SetChatID(chatID)
	u.SetLanguage(language)
	u.SetTopic(topic)
	u.SetIsToRu(isToRu)
	u.SetIterWord(iterWord)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return &u
	case nil:
		return &u
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return &u
	// return users[id]
}

func SaveUser(u *m.User) {
	id := u.GetChatID()
	// users[id] = u

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	sqlUserExist := `SELECT chatid FROM users WHERE chatid = $1`

	row := db.QueryRow(sqlUserExist, id)

	var chatID int64

	// unmarshal the row object to user
	err := row.Scan(&chatID)

	switch err {
	case sql.ErrNoRows:
		// create the insert sql query
		// returning userid will return the id of the inserted user
		sqlStatement := `INSERT INTO users (chatid, language, topic, istoru, iterword) VALUES ($1, $2, $3, $4, $5)`

		// execute the sql statement
		// Scan function will save the insert id in the id
		_, err = db.Exec(sqlStatement,
			u.GetChatID(),
			u.GetLanguage(),
			u.GetTopic(),
			u.GetIsToRu(),
			u.GetIterWord())

		fmt.Println(id)
		if err != nil {
			log.Fatalf("Save Unable to execute the query. %v", err)
		}

		fmt.Printf("Inserted a single record %v", id)
	case nil:
	default:
		log.Printf("User exist. %v", err)
	}
}

func UpdateUser(u *m.User) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE users SET language=$2, topic=$3, istoru=$4, iterword=$5 WHERE chatid=$1`

	_, err := db.Exec(sqlStatement,
		u.GetChatID(),
		u.GetLanguage(),
		u.GetTopic(),
		u.GetIsToRu(),
		u.GetIterWord())

	if err != nil {
		log.Fatalf("Update Unable to execute the query. %v", err)
	}

}

func SaveWords(chatID int64, original, translated string) {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO words (chatid, originalword, transletedword, istoru) VALUES($1, $2, $3, (SELECT istoru FROM users where chatid = $1))`

	_, err := db.Exec(sqlStatement,
		chatID,
		original,
		translated,
	)

	if err != nil {
		log.Fatalf("Update Unable to execute the query. %v", err)
	}

}
