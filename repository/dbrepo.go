package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbRepo struct {
	db *sqlx.DB
	currentId []byte
	alphabet Alphabet
}

func NewDbRepo(db *sqlx.DB) *DbRepo {
	return &DbRepo{
		db: db,
		alphabet: InitAlphabet(),
		currentId: []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
	}
}

func NewPostgresCon(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *DbRepo) AddNewUrl(url string) (string, error) {
	var shortenedUrl string
	query := fmt.Sprintf(`SELECT ID FROM %s WHERE url = '%s'`, tableName, url)
	row := r.db.QueryRow(query)
	err := row.Scan(&shortenedUrl)
	if err == sql.ErrNoRows {
		r.currentId, err = IncCurId(r.currentId, r.alphabet)
		if err == nil {
			shortenedUrl = string(r.currentId)
			query = fmt.Sprintf(`INSERT INTO %s VALUES ('%s', '%s')`, tableName, shortenedUrl, url)
			row = r.db.QueryRow(query)
			return shortenedUrl, nil
		} else {
			return "", err
		}
	}

	return "", errors.New("this url already has shortened version")
}

func (r DbRepo) GetUrl(shortened string) (string, error) {
	var url string
	query := fmt.Sprintf(`SELECT url FROM %s WHERE ID = '%s'`, tableName, shortened)
	row := r.db.QueryRow(query)
	err := row.Scan(&url)
	if err == sql.ErrNoRows {
		return "", errors.New("there's no url with such shortened version")
	}

	return url, nil
}