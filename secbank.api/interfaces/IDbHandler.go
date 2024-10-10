package interfaces

type IDbHandler interface {
	Execute(statement string)
	Query(statement string, dest interface{}) error
	Insert(entity interface{}, tableName string) error
	Update(id int, tableName string, updateData map[string]interface{}) error
	Delete(id int, tableName string) error
	Get(id int, tableName string, dest interface{}) error
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
