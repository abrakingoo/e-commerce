package db

const create_user_table = `CREATE TABLE IF NOT EXISTS users(
id VARCHAR(100) PRIMARY KEY NOT NULL,
firstname VARCHAR(50) NOT NULL,
lastname VARCHAR(50) NOT NULL,
phonenumber VARCHAR(10) NOT NULL UNIQUE,
email VARCHAR(50) NOT NULL UNIQUE,
password VARCHAR(50) NOT NULL);`

var Statements = []string{create_user_table,}