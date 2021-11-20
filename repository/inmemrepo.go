package repository

import "errors"

type InMemRepo struct {
	mp map[string]string
	currentId []byte
	alphabet Alphabet
}

func NewInMemRepo() *InMemRepo {
	return &InMemRepo{
		mp : make(map[string]string),
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
