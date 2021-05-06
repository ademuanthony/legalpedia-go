package postgres

import (
	"context"
	"fmt"
)

const (
	
)

var (
	createTableScripts = map[string]string{
	}

	tableOrder = []string{
	}

	// createIndexScripts is a map of table name to a collection of index on the table
	createIndexScripts = map[string][]string{
	}
)

func (pg *PgDb) CreateTables(ctx context.Context) error {
	tx, err := pg.db.Begin()
	if err != nil {
		return err
	}
	for _, tableName := range tableOrder {
		if exist := pg.TableExists(tableName); exist {
			continue
		}
		_, err := tx.Exec(createTableScripts[tableName])
		if err != nil {
			_ = tx.Rollback()
			log.Errorf("an error occurred while running %s", createTableScripts[tableName])
			return err
		}
		for _, createScript := range createIndexScripts[tableName] {
			_, err := tx.Exec(createScript)
			if err != nil {
				_ = tx.Rollback()
				log.Errorf("an error occurred while running %s", createIndexScripts[tableName])
				return err
			}
		}

		return tx.Commit()
	}
	return nil
}

func (pg *PgDb) TableExists(name string) bool {
	rows, err := pg.db.Query(`SELECT relname FROM pg_class WHERE relname = $1`, name)
	if err == nil {
		defer func() {
			if e := rows.Close(); e != nil {
				log.Error("Close of Query failed: ", e)
			}
		}()
		return rows.Next()
	}
	return false
}

func (pg *PgDb) DropTables() error {
	for tableName := range createTableScripts {
		if err := pg.dropTable(tableName); err != nil {
			return err
		}
	}
	return nil
}

func (pg *PgDb) dropTable(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, name))
	return err
}

func (pg *PgDb) dropIndex(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP INDEX IF EXISTS %s;`, name))
	return err
}
