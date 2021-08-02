package sybase

import (
	"database/sql"
	"database/sql/driver"
)

type SyBaseClient struct {
	conn   driver.Connector
	schema string
}

type Client interface {
	Query(sql string) (*sql.Rows, error)
}

func NewSyBaseClient(conn driver.Connector, schema string) *SyBaseClient {
	return &SyBaseClient{
		conn:   conn,
		schema: schema,
	}
}

var db *sql.DB

func (c *SyBaseClient) getDB() *sql.DB {
	if db == nil {
		db = sql.OpenDB(c.conn)
	}
	return db
}

func (c *SyBaseClient) moveToSchema(db *sql.DB) error {
	_, err := db.Exec("use " + c.schema)
	return err
}

//TODO(pozuecoa): deprecate
func (c *SyBaseClient) Query(sql string) (*sql.Rows, error) {
	db := c.getDB()
	if err := c.moveToSchema(db); err != nil {
		return nil, err
	}
	return db.Query(sql)
}

func (c *SyBaseClient) QueryMap(sql string) ([]map[string]interface{}, error) {
	rows, err := c.Query(sql)
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ToMap(rows)
}
