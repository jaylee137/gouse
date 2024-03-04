package connection

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"
// )

// type MySQLConnection struct {
// 	db *sql.DB
// }

// func NewMySQLConnection(uri string) (*MySQLConnection, error) {
// 	db, err := Connect(uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MySQLConnection{
// 		db: db,
// 	}, nil
// }

// func Connect(uri string) (*sql.DB, error) {
// 	db, err := sql.Open("mysql", uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func (connection *MySQLConnection) Close() error {
// 	return connection.db.Close()
// }

// func (connection *MySQLConnection) CreateDatabase(dbName string) (*sql.DB, error) {
// 	_, err := connection.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return connection.db, nil
// }

// func (connection *MySQLConnection) CreateTable(dbName string, tableName string, columns string) error {
// 	_, err := connection.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s)", dbName, tableName, columns))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (connection *MySQLConnection) Insert(dbName string, tableName string, columns string, values string) error {
// 	_, err := connection.db.Exec(fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", dbName, tableName, columns, values))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (connection *MySQLConnection) Query(query string) (*sql.Rows, error) {
// 	return connection.db.Query(query)
// }

// func (connection *MySQLConnection) QueryRow(query string) *sql.Row {
// 	return connection.db.QueryRow(query)
// }

// func (connection *MySQLConnection) Exec(query string) (sql.Result, error) {
// 	return connection.db.Exec(query)
// }

// func (connection *MySQLConnection) Prepare(query string) (*sql.Stmt, error) {
// 	return connection.db.Prepare(query)
// }

// func (connection *MySQLConnection) Begin() (*sql.Tx, error) {
// 	return connection.db.Begin()
// }

// func (connection *MySQLConnection) SetMaxOpenConns(n int) {
// 	connection.db.SetMaxOpenConns(n)
// }

// func (connection *MySQLConnection) SetMaxIdleConns(n int) {
// 	connection.db.SetMaxIdleConns(n)
// }

// func (connection *MySQLConnection) SetConnMaxLifetime(d time.Duration) {
// 	connection.db.SetConnMaxLifetime(d)
// }

// func (connection *MySQLConnection) SetConnMaxIdleTime(d time.Duration) {
// 	connection.db.SetConnMaxIdleTime(d)
// }
