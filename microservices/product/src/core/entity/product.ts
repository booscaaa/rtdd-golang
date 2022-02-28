import { sendUnaryData, ServerUnaryCall } from "@grpc/grpc-js";
import { FetchRequest, FetchResponse, Product, ProductServiceServer } from "../domain/product";

interface ProductUseCase {
    fetch(): Promise<Product[]>
}

interface ProductRepository {
    fetch(): Promise<Product[]>
}

type KnownKeys<T> = {
    [K in keyof T]: string extends K ? never : number extends K ? never : K
} extends { [_ in keyof T]: infer U } ? U : never;

type KnownOnly<T extends Record<any, any>> = Pick<T, KnownKeys<T>>;

type IProductServiceServer = KnownOnly<ProductServiceServer>;

export { ProductUseCase, ProductRepository, IProductServiceServer }