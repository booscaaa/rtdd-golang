import sqlite3 from "sqlite3";
import { Product } from "../../core/domain/product";
import { ProductRepository } from "../../core/entity/product";

class ProductRepositoryImpl implements ProductRepository {
    constructor(
        private db: sqlite3.Database
    ) { }

    async fetch(): Promise<Product[]> {
        try {
            const sql = 'SELECT * FROM product'

            const products: any = await new Promise((resolve, reject) => {
                this.db.all(sql, [], (err, row) => {
                    if (err)
                        reject(err)

                    resolve(row)
                })
            })
            
            return products.map((product: any) => Product.fromJSON(product))
        } catch (error) {
            throw error
        }
    }
}

export { ProductRepositoryImpl }