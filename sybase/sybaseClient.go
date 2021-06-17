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

func (c *SyBaseClient) getDB() *sql.DB {
	db := sql.OpenDB(c.conn)
	return db
}

func (c *SyBaseClient) moveToSchema(db *sql.DB) error {
	_, err := db.Exec("use " + c.schema)
	return err
}

func (c *SyBaseClient) Query(sql string) (*sql.Rows, error) {
	db := c.getDB()
	if err := c.moveToSchema(db); err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Query(sql)
}
