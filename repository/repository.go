package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UrlManager interface {
	AddNewUrl(url string) (string, error)
	GetUrl(shortened string) (string, error)
}

const (
	alphabetSize		=	63
	upperLettersStart	=	10
	lowerLettersStart	=	36
	underline			=	62
	tableName			= 	"shortened_urls"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Alphabet struct {
	order 	map[byte]int
	symbol 	map[int]byte
}

type UrlsDb struct {
	db *sqlx.DB
	currentId []byte
	alphabet Alphabet
}

func NewUrlsDb(db *sqlx.DB) *UrlsDb {
	return &UrlsDb{
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

func (r *UrlsDb) AddNewUrl(url string) (string, error) {
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

func (r UrlsDb) GetUrl(shortened string) (string, error) {
	var url string
	query := fmt.Sprintf(`SELECT url FROM %s WHERE ID = '%s'`, tableName, shortened)
	row := r.db.QueryRow(query)
	err := row.Scan(&url)
	if err == sql.ErrNoRows {
		return "", errors.New("there's no url with such shortened version")
	}

	return url, nil
}

func InitAlphabet() Alphabet {
	var sym byte
	var alphabet Alphabet
	alphabet.order = make(map[byte]int, alphabetSize)
	alphabet.symbol = make(map[int]byte, alphabetSize)

	sym = '0'
	for i := 0; i < alphabetSize; i += 1 {
		if i == upperLettersStart {
			sym = 'A'
		} else if i == lowerLettersStart {
			sym = 'a'
		} else if i == underline {
			sym = '_'
		}
		alphabet.order[sym] = i
		alphabet.symbol[i] = sym
		sym++
	}

	return alphabet
}

func IncCurId(curId []byte, alphabet Alphabet) ([]byte, error) {
	index := 9
	for ;index >= 0; index -= 1 {
		newSymOrd := alphabet.order[curId[index]] + 1
		if newSymOrd == alphabetSize {
			curId[index] = alphabet.symbol[0]
		} else {
			curId[index] = alphabet.symbol[newSymOrd]
			break
		}
	}

	if index == -1 {
		return []byte{}, errors.New("overflow")
	}

	return curId, nil
}

type InMemRepo struct {
	mp map[string]string
	currentId []byte
	alphabet Alphabet
}

func NewInMemRepo() *InMemRepo {
	return &InMemRepo{
		currentId: []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		alphabet: InitAlphabet(),
	}
}

func (r *InMemRepo) AddNewUrl(url string) (string, error) {
	contains := false
	for _, x := range r.mp {
		if x == url {
			contains = true
			break
		}
	}
	if contains {
		return "", errors.New("this url already has shortened version")
	}

	var err error
	r.currentId, err = IncCurId(r.currentId, r.alphabet)
	if err == nil {
		shortenedUrl := string(r.currentId)
		r.mp[shortenedUrl] = url
		return shortenedUrl, nil
	} else {
		return "", err
	}
}

func (r *InMemRepo) GetUrl(shortened string) (string, error) {
	if url, ok := r.mp[shortened]; ok {
		return url, nil
	}

	return "", errors.New("there's no url with such shortened version")
}