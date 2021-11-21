package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestDbRepo_AddNewUrl(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		db        *sqlx.DB
		currentId []byte
		alphabet  Alphabet
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock	func()
		want    string
		wantErr bool
	}{
		{
			"add",
			fields{
				db,
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				InitAlphabet(),
			},
			args{
				"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit",
			},
			func() {
				rows := sqlmock.NewRows([]string{})
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			"0000000001",
			false,
		},
		{
			"same",
			fields{
				db,
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
				InitAlphabet(),
			},
			args{
				"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit",
			},
			func() {
				rows := sqlmock.NewRows([]string{"ID"}).AddRow("0000000001")
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			"",
			true,
		},
		{
			"overflow",
			fields{
				db,
				[]byte{'_', '_', '_', '_', '_', '_', '_', '_', '_', '_'},
				InitAlphabet(),
			},
			args{
				"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit",
			},
			func() {
				rows := sqlmock.NewRows([]string{})
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			"",
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DbRepo{
				db:        tt.fields.db,
				currentId: tt.fields.currentId,
				alphabet:  tt.fields.alphabet,
			}
			tt.mock()
			got, err := r.AddNewUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNewUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddNewUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbRepo_GetUrl(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		db        *sqlx.DB
		currentId []byte
		alphabet  Alphabet
	}
	type args struct {
		shortened string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock 	func()
		want    string
		wantErr bool
	}{
		{
			"get",
			fields{
				db,
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
				InitAlphabet(),
			},
			args{
				"0000000001",
			},
			func() {
				rows := sqlmock.NewRows([]string{"url"}).AddRow("https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit")
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit",
			false,
		},
		{
			"absent",
			fields{
				db,
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				InitAlphabet(),
			},
			args{
				"0000000001",
			},
			func() {
				rows := sqlmock.NewRows([]string{})
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			},
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := DbRepo{
				db:        tt.fields.db,
				currentId: tt.fields.currentId,
				alphabet:  tt.fields.alphabet,
			}
			tt.mock()
			got, err := r.GetUrl(tt.args.shortened)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
