package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type Key struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
	Times  int    `json:"times"`
	Used   int    `json:"used"`
}

var Conn *pgx.Conn

func init() {
	var err error
	Conn, err = pgx.Connect(context.Background(), Conf.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func GetKey(key string) *Key {
	rows, err := Conn.Query(context.Background(), "SELECT key,secret,times,used FROM key WHERE key = $1", key)
	if err != nil {
		return nil
	}
	k, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Key])
	if err != nil {
		return nil
	}
	return &k
}

func Update(key string, use int) {
	_, err := Conn.Exec(context.Background(), "UPDATE key SET used = used + $1 WHERE key=$2", use, key)
	if err != nil {
		return
	}
}
