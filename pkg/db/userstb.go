package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	log "github.com/tmb-piXel/LearnEnglishBot/pkg/logger"
	m "github.com/tmb-piXel/LearnEnglishBot/pkg/models"
)

var db *sql.DB

func InitDB(postgresUrl string) {
	db = createConnection(postgresUrl)
}

// create connection with postgres db
func createConnection(postgresUrl string) *sql.DB {
	db, err := sql.Open("postgres", postgresUrl)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("connect ok")
	return db
}

func FindUser(id int64) *m.User {
	var u m.User

	sqlStatement := `SELECT * FROM users WHERE chatid=$1`

	row := db.QueryRow(sqlStatement, id)

	var (
		chatID   int64
		language string
		topic    string
		isToRu   bool
		iterWord int
	)
	err := row.Scan(&chatID, &language, &topic, &isToRu, &iterWord)

	u.SetChatID(chatID)
	u.SetLanguage(language)
	u.SetTopic(topic)
	u.SetIsToRu(isToRu)
	u.SetIterWord(iterWord)

	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		return &u
	case nil:
		return &u
	default:
		log.Errorf("Unable to scan the row. %v", err)
	}

	return &u
}

func SaveUser(u *m.User) {
	id := u.GetChatID()

	sqlUserExist := `SELECT chatid FROM users WHERE chatid = $1`

	row := db.QueryRow(sqlUserExist, id)

	var chatID int64

	err := row.Scan(&chatID)

	switch err {
	case sql.ErrNoRows:
		sqlStatement := `INSERT INTO users (chatid, language, topic, istoru, iterword) VALUES ($1, $2, $3, $4, $5)`

		_, err = db.Exec(sqlStatement,
			u.GetChatID(),
			u.GetLanguage(),
			u.GetTopic(),
			u.GetIsToRu(),
			u.GetIterWord())

		log.Println(id)
		if err != nil {
			log.Errorf("Save Unable to execute the query. %v", err)
		}

		log.Println("Inserted a single record %v", id)
	case nil:
	default:
		log.Println("User exist. %v", err)
	}
}

func UpdateUser(u *m.User) {
	sqlStatement := `UPDATE users SET language=$2, topic=$3, istoru=$4, iterword=$5 WHERE chatid=$1`

	_, err := db.Exec(sqlStatement,
		u.GetChatID(),
		u.GetLanguage(),
		u.GetTopic(),
		u.GetIsToRu(),
		u.GetIterWord())

	if err != nil {
		log.Errorf("Update Unable to execute the query. %v", err)
	}

}
