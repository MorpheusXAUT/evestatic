// mysql
package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDatasource struct {
	source string
	con    *sql.DB
}

func NewMysqlDatasource(src string) (*MysqlDatasource, error) {
	mysql := &MysqlDatasource{
		source: src,
	}

	c, err := sql.Open("mysql", mysql.source)
	if err != nil {
		return nil, err
	}

	mysql.con = c

	return mysql, nil
}

func (mysql *MysqlDatasource) Close() {
	mysql.con.Close()
}

func (mysql *MysqlDatasource) RawQuery(query string, args ...interface{}) (interface{}, error) {
	if err := mysql.con.Ping(); err != nil {
		return nil, err
	}

	return mysql.con.Query(query, args)
}

func (mysql *MysqlDatasource) GetTypeNameFromID(id int) (string, error) {
	if err := mysql.con.Ping(); err != nil {
		return "", err
	}

	rows, err := mysql.con.Query("SELECT `typeName` FROM `invTypes` WHERE `typeID` = ?", id)
	if err != nil {
		return "", err
	}

	if !rows.Next() {
		return "", fmt.Errorf("No rows found for typeID %d", id)
	}

	var name string

	err = rows.Scan(&name)
	if err != nil {
		return "", nil
	}

	return name, nil
}

func (mysql *MysqlDatasource) GetTypeIDFromName(name string) (int, error) {
	if err := mysql.con.Ping(); err != nil {
		return -1, err
	}

	rows, err := mysql.con.Query("SELECT `typeID` FROM `invTypes` WHERE `typeName` LIKE ?", name)
	if err != nil {
		return -1, err
	}

	if !rows.Next() {
		return -1, fmt.Errorf("No rows found for typeName %q", name)
	}

	var id int

	err = rows.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
