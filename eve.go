// eve
package evestatic

import (
	"github.com/morpheusxaut/evestatic/datasource"
)

type EveStatic struct {
	datasource Datasource
}

func NewEveStatic(dsType DatasourceType, dsSource string) (*EveStatic, error) {
	eve := &EveStatic{}

	switch dsType {
	case DatasourceTypeMySql:
		data, err := datasource.NewMysqlDatasource(dsSource)
		if err != nil {
			return nil, err
		}

		eve.datasource = data
		break
	case DatasourceTypeSQLite3:
		data, err := datasource.NewSqlite3Datasource(dsSource)
		if err != nil {
			return nil, err
		}

		eve.datasource = data
		break
	}

	return eve, nil
}

func (eve *EveStatic) Close() {
	eve.datasource.Close()
}

func (eve *EveStatic) RawQuery(query string, args ...interface{}) (interface{}, error) {
	return eve.datasource.RawQuery(query, args)
}

func (eve *EveStatic) GetTypeNameFromID(id int) (string, error) {
	return eve.datasource.GetTypeNameFromID(id)
}

func (eve *EveStatic) GetTypeIDFromName(name string) (int, error) {
	return eve.datasource.GetTypeIDFromName(name)
}
