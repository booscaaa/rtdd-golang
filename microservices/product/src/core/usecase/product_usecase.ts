
import { Product } from "../domain/product";
import { ProductRepository, ProductUseCase } from "../entity/product";

class ProductUseCaseImpl implements ProductUseCase {
    constructor(
        private productRepository: ProductRepository
    ) { }

    async fetch(): Promise<Product[]> {
        try {
            const products = await this.productRepository.fetch()
            
            return products
        } catch (error) {
            throw error
        }
    }
}

export { ProductUseCaseImpl }