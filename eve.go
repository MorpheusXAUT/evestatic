// eve
package evestatic

import (
	"github.com/morpheusxaut/evestatic/datasource"
)

type EveStatic struct {
	datasource Datasource
}

func NewEveStatic(dsType DatasourceType, dsSource string) *EveStatic {
	eve := &EveStatic{}

	switch dsType {
	case DatasourceTypeMySql:
		eve.datasource = datasource.NewMysqlDatasource(dsSource)
		break
	case DatasourceTypeSQLite3:
		eve.datasource = datasource.NewSqlite3Datasource(dsSource)
		break
	}

	return eve
}
