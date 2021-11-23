package constants

import "fmt"

//Port is a
const Port = ":8080"

//postgresAddress is the address of the postgres
const postgresAddress = "192.168.53.62"

//postgresPort is the port of the postgres
const postgresPort = "6432"

//postgresDataBase is the name of the database
const postgresDatabaseMWallet = "mwallet"

//postgresDatabaseEWallet is the name of the database
const postgresDatabaseEWallet = "ewallet"

//postgresUsernameMWallet is the name of the user inside DBA
const postgresUsernameMWallet = "mwallet"

//postgresPasswordMWallet is the password of the user
const postgresPasswordMWallet = "k1}CS}xmDvpcsMr"

//postgresUsernameEWallet is the name of the user inside DBA
const postgresUsernameEWallet = "ewallet"

//postgresPasswordEWallet is the password of the user
const postgresPasswordEWallet = "XABA:gEN8T:REX8"

//PostgresConnectionMWallet is the connection string to the database
var PostgresConnectionMWallet = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	postgresAddress, postgresPort, postgresUsernameMWallet, postgresPasswordMWallet, postgresDatabaseMWallet)

//PostgresConnectionEWallet is the connection string to the database
var PostgresConnectionEWallet = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	postgresAddress, postgresPort, postgresUsernameEWallet, postgresPasswordEWallet, postgresDatabaseEWallet)

//MongoClientURI is a
var MongoClientURI = "mongodb://mongoRoot:mongoRoot@192.168.53.5:27017"
