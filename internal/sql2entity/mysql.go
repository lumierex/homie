package sql2entity

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

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
	ColumnName    string
	ColumnType    string
	ColumnKey     string
	ColumnComment string
	DataType      string
	IsNullable    string
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

// 获取数据库下的某张表的元信息
func (m *DBModel) Columns(dbName, tableName string) ([]*TableColumn, error) {
	// table_schema 数据库名称 https://blog.csdn.net/qq_42778001/article/details/120035616
	query := "SELECT " + "COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FORM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME= ? "
	rows, err := m.Engine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn

	for rows.Next() {

		var column TableColumn
		rows.Scan(
			&column.ColumnName,
			&column.ColumnType,
			&column.ColumnKey,
			&column.ColumnComment,
			&column.DataType,
			&column.IsNullable,
		)
		columns = append(columns, &column)
	}
	log.Println(columns)
	return columns, nil

}
