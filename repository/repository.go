package repository

type UrlManager interface {
	AddNewUrl(url string) (string, error)
	GetUrl(shortened string) (string, error)
}

type Repo struct {
	rep UrlManager
}

func InitRepo(mode byte, path string) (Repo, error) {
	var r Repo
	if mode == 'd' {
		db, err := NewPostgresCon(Config{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			DBName:   "postgres",
			SSLMode:  "disable",
			Password: "postgrepass",
		})
		if err == nil {
			r.rep = NewDbRepo(db)
		} else {
			return r, err
		}
	} else if mode == 'i' {
		r.rep = NewInMemRepo()
	}

	return r, nil
}

func (r *Repo) CallAddNewUrl(url string) (string, error) {
	return r.rep.AddNewUrl(url)
}

func (r *Repo) CallGetUrl(shortened string) (string, error) {
	return r.rep.GetUrl(shortened)
}