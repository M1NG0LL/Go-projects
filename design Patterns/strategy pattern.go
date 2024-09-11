package main

import "fmt"

type IDBconnection interface {
	connect()
}

type DBconnection struct {
	Db IDBconnection
}

func (con DBconnection) Dbconnect() {
	con.Db.connect()
}

type MYSQLConnection struct {
	Connectionstring string
}

func (con MYSQLConnection) connect() {
	fmt.Println(("MySql " + con.Connectionstring))
}

// ================

type PostgressConnection struct {
	ConnectionString string
}

func (con PostgressConnection) connect() {
	fmt.Println("Postgress " + con.ConnectionString)
}

// ================
type MongoDbConnection struct {
	ConnectionString string
}

func (con MongoDbConnection) connect() {
	fmt.Println("MongoDb " + con.ConnectionString)
}


func main() {
	MySqlConnection := MYSQLConnection{Connectionstring: "MySql DB is connected"}
	con := DBconnection{Db: MySqlConnection}
	con.Dbconnect()

	pgConnection := PostgressConnection{ConnectionString: "Postgress DB is connected"}
	con2 := DBconnection{Db: pgConnection}
	con2.Dbconnect()

	mgConnection := MongoDbConnection{ConnectionString: "Mongo DB is connected"}
	con3 := DBconnection{Db: mgConnection}
	con3.Dbconnect()
}