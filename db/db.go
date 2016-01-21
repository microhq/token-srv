package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	token "github.com/micro/token-srv/proto/record"
)

var (
	db       *sql.DB
	Url      = "root@tcp(127.0.0.1:3306)/token"
	database string

	tokenSchema = `CREATE TABLE IF NOT EXISTS tokens (
id varchar(36) primary key,
namespace varchar(64),
name varchar(64),
created integer,
updated integer,
unique (namespace, name));`

	q = map[string]string{
		"delete": "DELETE from %s.%s where id = ?",
		"create": `INSERT into %s.%s (
				id, namespace, name, created, updated) 
				values (?, ?, ?, ?, ?)`,
		"update":                 "UPDATE %s.%s set namespace = ?, name = ?, updated = ? where id = ?",
		"read":                   "SELECT * from %s.%s where id = ?",
		"list":                   "SELECT * from %s.%s limit ? offset ?",
		"searchNamespace":        "SELECT * from %s.%s where namespace = ? limit ? offset ?",
		"searchName":             "SELECT * from %s.%s where name = ? limit ? offset ?",
		"searchNamespaceAndName": "SELECT * from %s.%s where namespace = ? and name = ? limit ? offset ?",
	}
	st = map[string]*sql.Stmt{}
)

func Init() {
	var d *sql.DB
	var err error

	parts := strings.Split(Url, "/")
	if len(parts) != 2 {
		panic("Invalid database url")
	}

	if len(parts[1]) == 0 {
		panic("Invalid database name")
	}

	url := parts[0]
	database = parts[1]

	if d, err = sql.Open("mysql", url+"/"); err != nil {
		log.Fatal(err)
	}
	if _, err := d.Exec("CREATE DATABASE IF NOT EXISTS " + database); err != nil {
		log.Fatal(err)
	}
	d.Close()
	if d, err = sql.Open("mysql", Url); err != nil {
		log.Fatal(err)
	}
	if _, err = d.Exec(tokenSchema); err != nil {
		log.Fatal(err)
	}
	db = d

	for query, statement := range q {
		prepared, err := db.Prepare(fmt.Sprintf(statement, database, "tokens"))
		if err != nil {
			log.Fatal(err)
		}
		st[query] = prepared
	}
}

func Create(token *token.Token) error {
	token.Created = time.Now().Unix()
	token.Updated = time.Now().Unix()
	_, err := st["create"].Exec(token.Id, token.Namespace, token.Name, token.Created, token.Updated)
	return err
}

func Delete(id string) error {
	_, err := st["delete"].Exec(id)
	return err
}

func Update(token *token.Token) error {
	token.Updated = time.Now().Unix()
	_, err := st["update"].Exec(token.Namespace, token.Name, token.Updated, token.Id)
	return err
}

func Read(id string) (*token.Token, error) {
	token := &token.Token{}

	r := st["read"].QueryRow(id)
	if err := r.Scan(&token.Id, &token.Namespace, &token.Name, &token.Created, &token.Updated); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return token, nil
}

func Search(namespace, name string, limit, offset int64) ([]*token.Token, error) {
	var r *sql.Rows
	var err error

	if len(namespace) > 0 && len(name) > 0 {
		r, err = st["searchNamespaceAndName"].Query(namespace, name, limit, offset)
	} else if len(namespace) > 0 {
		r, err = st["searchNamespace"].Query(namespace, limit, offset)
	} else if len(name) > 0 {
		r, err = st["searchName"].Query(name, limit, offset)
	} else {
		r, err = st["list"].Query(limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer r.Close()

	var tokens []*token.Token

	for r.Next() {
		token := &token.Token{}
		if err := r.Scan(&token.Id, &token.Namespace, &token.Name, &token.Created, &token.Updated); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("not found")
			}
			return nil, err
		}
		tokens = append(tokens, token)

	}
	if r.Err() != nil {
		return nil, err
	}

	return tokens, nil
}
