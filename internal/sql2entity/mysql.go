package sql2entity

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	Engine *sql.DB
	Info   *DBInfo
}

type DBInfo struct {
	Type     string
	Host     string
	Username string
	Password string
	Charset  string
}

type TableColumn struct {
	Name       string
	Type       string
	Key        string
	Comment    string
	DataType   string
	IsNullable string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		Info: info,
	}
}

func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.Info.Username,
		m.Info.Password,
		m.Info.Host,
		m.Info.Charset)
	m.Engine, err = sql.Open(m.Info.Type, dsn)
	if err != nil {
		return err
	}
	return nil
}
