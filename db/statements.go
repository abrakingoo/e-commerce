package db

const create_user_table = `CREATE TABLE IF NOT EXISTS users(
id VARCHAR(100) PRIMARY KEY NOT NULL,
firstname VARCHAR(50) NOT NULL,
lastname VARCHAR(50) NOT NULL,
phonenumber VARCHAR(10) NOT NULL UNIQUE,
email VARCHAR(50) NOT NULL UNIQUE,
role VARCHAR,
password VARCHAR(50) NOT NULL);`


const create_products_table = `CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    price VARCHAR(50) NOT NULL,
    description VARCHAR,
    category VARCHAR(50),
    image VARCHAR,
    total VARCHAR,
    rating INTEGER DEFAULT 0
);`

const create_orders_table = `CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    userid VARCHAR(50) NOT NULL,
    productid VARCHAR(50) NOT NULL,
	total VARCHAR NOT NULL,
    status BOOLEAN DEFAULT false,
    FOREIGN KEY (userid) REFERENCES users(id)
);`

var Statements = []string{create_user_table, create_orders_table, create_products_table}