package connection

// import (
// 	"database/sql"
// 	"fmt"
// 	"github.com/thuongtruong1009/gouse/types"
// 	"log"
// 	"reflect"
// 	"strings"
// )

// type Postgres struct {
// 	Host     string
// 	Port     int
// 	Username string
// 	Password string
// 	Database string
// }

// func (p Postgres) Connect() *sql.DB {
// 	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// func (p Postgres) Insert(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var fields []string
// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldName := t.Field(i).Name
// 		fieldValue := v.Field(i).Interface()
// 		fields = append(fields, fieldName)
// 		values = append(values, types.InterfaceToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s')", t.Name(), strings.Join(fields, ", "), strings.Join(values, "', '"))
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (p Postgres) Update(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var fields []string
// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldName := t.Field(i).Name
// 		fieldValue := v.Field(i).Interface()
// 		fields = append(fields, fieldName)
// 		values = append(values, types.InterfaceToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %s", t.Name(), strings.Join(fields, ", "), strings.Join(values, "', '"))
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (p Postgres) Delete(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldValue := v.Field(i).Interface()
// 		values = append(values, types.InterfaceToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("DELETE FROM %s WHERE id = %s", t.Name(), strings.Join(values, "', '"))
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (p Postgres) Select(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	var values []string

// 	for i := 0; i < v.NumField(); i++ {
// 		fieldValue := v.Field(i).Interface()
// 		values = append(values, types.InterfaceToString(fieldValue))
// 	}

// 	query := fmt.Sprintf("SELECT * FROM %s WHERE id = %s", t.Name(), strings.Join(values, "', '"))
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func (p Postgres) SelectAll(data interface{}) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	query := fmt.Sprintf("SELECT * FROM %s", t.Name())
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func (p Postgres) SelectWhere(data interface{}, where string) {
// 	db := p.Connect()
// 	defer db.Close()

// 	v := reflect.ValueOf(data)
// 	t := v.Type()

// 	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", t.Name(), where)
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }