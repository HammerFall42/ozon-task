package repository

import (
	"testing"
)

func TestInMemRepo_AddNewUrl(t *testing.T) {
	type fields struct {
		mp        map[string]string
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
		want    string
		wantErr bool
	}{
		{
			"add",
			fields{map[string]string{},
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				InitAlphabet()},
				args{"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit"},
				"0000000001",
				false,
		},
		{
			"same",
			fields{map[string]string{"0000000001": "https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit"},
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
				InitAlphabet()},
			args{"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit"},
			"",
			true,
		},
		{
			"overflow",
			fields{map[string]string{},
				[]byte{'_', '_', '_', '_', '_', '_', '_', '_', '_', '_'},
				InitAlphabet()},
			args{"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemRepo{
				mp:        tt.fields.mp,
				currentId: tt.fields.currentId,
				alphabet:  tt.fields.alphabet,
			}
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

func TestInMemRepo_GetUrl(t *testing.T) {
	type fields struct {
		mp        map[string]string
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
		want    string
		wantErr bool
	}{
		{
			"get",
			fields{
				map[string]string{"0000000001": "https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit"},
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '1'},
				InitAlphabet(),
			},
			args{
				"0000000001",
			},
			"https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit",
			false,
		},
		{
			"absent",
			fields{
				map[string]string{},
				[]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
				InitAlphabet(),
			},
			args{
				"0000000001",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemRepo{
				mp:        tt.fields.mp,
				currentId: tt.fields.currentId,
				alphabet:  tt.fields.alphabet,
			}
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
