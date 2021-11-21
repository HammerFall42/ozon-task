package repository

import "github.com/spf13/viper"

type UrlManager interface {
	AddNewUrl(url string) (string, error)
	GetUrl(shortened string) (string, error)
}

type Repo struct {
	rep UrlManager
}

func InitRepo(mode byte) (Repo, error) {
	var r Repo
	if mode == 'd' {
		db, err := NewPostgresCon(Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: viper.GetString("db.password"),
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