// mysql
package datasource

type MysqlDatasource struct {
}

func NewMysqlDatasource(source string) *MysqlDatasource {
	mysql := &MysqlDatasource{}

	return mysql
}

func (mysql *MysqlDatasource) RawQuery(query string) string {
	return ""
}

func (mysql *MysqlDatasource) GetNameFromID(id int) string {
	return ""
}

func (mysql *MysqlDatasource) GetIDFromName(name string) int {
	return 0
}
