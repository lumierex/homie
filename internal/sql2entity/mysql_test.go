package sql2entity

import (
	"database/sql"
	"log"
	"testing"
)

func TestDBModel_Connect(t *testing.T) {
	type fields struct {
		Engine *sql.DB
		Info   *DBInfo
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "mysql",
			fields: fields{
				Engine: &sql.DB{},
				Info: &DBInfo{
					Type:     "mysql",
					Username: "root",
					Password: "12345678",
					Host:     "localhost",
					Charset:  "utf8mb4",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DBModel{
				Engine: tt.fields.Engine,
				Info:   tt.fields.Info,
			}
			if err := m.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("DBModel.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
			tableColumns, err := m.Columns("mtl", "user")
			if err != nil {
				t.Errorf("DBModel.Columns() error = %v", err)
			}
			for _, column := range tableColumns {
				log.Println(*column)
			}
		})
	}
}
