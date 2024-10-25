package infrastructures

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"secbank.api/utils"
)

type SQLHandler struct {
	Conn *sqlx.DB
}

func (handler *SQLHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SQLHandler) Query(statement string, dest interface{}) error {

	err := handler.Conn.Select(dest, statement)

	return err
}

func (handler *SQLHandler) QueryWithParamSingleRow(statement string, dest interface{}, args ...interface{}) error {

	err := handler.Conn.Get(dest, statement, args...)

	return err
}

func (handler *SQLHandler) Insert(entity interface{}, tableName string) error {
	dialect := goqu.Dialect("postgres")

	entityWithoutID, err := utils.StructToMapWithoutID(entity, "id")

	if err != nil {
		return err
	}

	insert := dialect.Insert(tableName).Rows(entityWithoutID)

	sql, args, err := insert.ToSQL()

	if err != nil {
		return err
	}

	row := handler.Conn.QueryRow(sql, args...)

	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (handler *SQLHandler) Delete(id int, tableName string) error {
	dialect := goqu.Dialect("postgres")

	deleteQuery := dialect.Delete(tableName).Where(goqu.Ex{"id": id})

	sql, args, err := deleteQuery.ToSQL()
	if err != nil {
		return fmt.Errorf("error generating SQL: %v", err)
	}

	result, err := handler.Conn.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing DELETE: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	return nil
}

func (handler *SQLHandler) Update(id int, tableName string, updateData map[string]interface{}) error {
	dialect := goqu.Dialect("postgres")

	updateQuery := dialect.Update(tableName).Set(updateData).Where(goqu.Ex{"id": id})

	sql, args, err := updateQuery.ToSQL()
	if err != nil {
		return fmt.Errorf("error generating SQL: %v", err)
	}

	result, err := handler.Conn.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing UPDATE: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (handler *SQLHandler) Get(id int, tableName string, dest interface{}) error {
	dialect := goqu.Dialect("postgres")

	selectQuery := dialect.From(tableName).Where(goqu.Ex{"id": id})

	sql, args, err := selectQuery.ToSQL()
	if err != nil {
		return fmt.Errorf("error generating SQL: %v", err)
	}

	err = handler.Conn.Get(dest, sql, args...)
	if err != nil {
		return fmt.Errorf("error fetching record: %v", err)
	}

	return nil
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}
