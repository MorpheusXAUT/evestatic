// database
package evestatic

type Datasource interface {
	RawQuery(query string) string
	GetNameFromID(id int) string
	GetIDFromName(name string) int
}
