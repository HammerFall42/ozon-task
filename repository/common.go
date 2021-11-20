package repository

import "errors"

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
