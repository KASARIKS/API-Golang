package db

import "database/sql"

const usersTable = "CREATE TABLE IF NOT EXISTS users (" +
	"id INT NOT NULL PRIMARY KEY," +
	"firstName VARCHAR(255) NOT NULL," +
	"lastName VARCHAR(255) NOT NULL," +
	"email VARCHAR(255) NOT NULL UNIQUE," +
	"password VARCHAR(255) NOT NULL," +
	"createdAt TEXT NOT NULL" +
	");"

const productsTable = "CREATE TABLE IF NOT EXISTS products (" +
	"id INT NOT NULL PRIMARY KEY," +
	"name VARCHAR(255) NOT NULL," +
	"description TEXT NOT NULL," +
	"image VARCHAR(255) NOT NULL," +
	"price DECIMAL(10, 2) NOT NULL," +
	"quantity INT NOT NULL," +
	"createdAt TEXT NOT NULL" +
	");"

const ordersTable = "CREATE TABLE IF NOT EXISTS orders (" +
	"id INT NOT NULL PRIMARY KEY," +
	"userId INT NOT NULL," +
	"total DECIMAL(10, 2) NOT NULL," +
	"status TEXT NOT NULL," +
	"address TEXT NOT NULL," +
	"createdAt TEXT NOT NULL," +
	"FOREIGN KEY (userId) REFERENCES users(id)" +
	");"

const ordersItemsTable = "CREATE TABLE IF NOT EXISTS order_items (" +
	"id INT NOT NULL PRIMARY KEY," +
	"orderId INT NOT NULL," +
	"productId INT NOT NULL," +
	"quantity INT NOT NULL," +
	"price DECIMAL(10, 2) NOT NULL," +
	"FOREIGN KEY (orderId) REFERENCES orders(id)," +
	"FOREIGN KEY (productId) REFERENCES products(id)" +
	");"

func createTables(db *sql.DB) error {
	if _, err := db.Exec(usersTable); err != nil {
		return err
	}

	if _, err := db.Exec(productsTable); err != nil {
		return err
	}

	if _, err := db.Exec(ordersTable); err != nil {
		return err
	}

	if _, err := db.Exec(ordersItemsTable); err != nil {
		return err
	}

	return nil
}
