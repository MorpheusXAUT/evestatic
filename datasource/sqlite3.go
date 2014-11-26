// sqlite3
package datasource

type Sqlite3Datasource struct {
}

func NewSqlite3Datasource(source string) *Sqlite3Datasource {
	sqlite3 := &Sqlite3Datasource{}

	return sqlite3
}

func (sqlite3 *Sqlite3Datasource) RawQuery(query string) string {
	return ""
}

func (sqlite3 *Sqlite3Datasource) GetNameFromID(id int) string {
	return ""
}

func (sqlite3 *Sqlite3Datasource) GetIDFromName(name string) int {
	return 0
}
