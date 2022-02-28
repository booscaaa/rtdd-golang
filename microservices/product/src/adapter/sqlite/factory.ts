import sqlite3 from 'sqlite3'

const DBSOURCE = 'product.db'
const SQL_ITENS_CREATE = `
    CREATE TABLE product (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        price float
    );`

const INSERTED = "INSERT INTO product(name, price) VALUES ('Arroz', 10.25);"
const database = new sqlite3.Database(DBSOURCE, (err) => {
    if (err) {
        console.error(err.message)
        throw err
    } else {
        database.run(SQL_ITENS_CREATE, (_) => { })

        database.run(INSERTED, (_) => { })
    }
})

export { database }