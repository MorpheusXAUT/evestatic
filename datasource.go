// database
package evestatic

type Datasource interface {
	Close()
	RawQuery(query string, args ...interface{}) (interface{}, error)
	GetTypeNameFromID(id int) (string, error)
	GetTypeIDFromName(name string) (int, error)
}
