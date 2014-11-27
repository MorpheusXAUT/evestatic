// sqlite3
package datasource

type Sqlite3Datasource struct {
}

func NewSqlite3Datasource(source string) (*Sqlite3Datasource, error) {
	sqlite3 := &Sqlite3Datasource{}

	return sqlite3, nil
}

func (sqlite3 *Sqlite3Datasource) Close() {

}

func (sqlite3 *Sqlite3Datasource) RawQuery(query string, args ...interface{}) (interface{}, error) {
	return "", nil
}

func (sqlite3 *Sqlite3Datasource) GetTypeNameFromID(id int) (string, error) {
	return "", nil
}

func (sqlite3 *Sqlite3Datasource) GetTypeIDFromName(name string) (int, error) {
	return 0, nil
}
