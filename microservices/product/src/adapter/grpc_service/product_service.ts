import { FetchRequest, FetchResponse, ProductServiceServer } from "../../core/domain/product";
import { sendUnaryData, ServerUnaryCall, UntypedHandleCall } from "@grpc/grpc-js";
import { IProductServiceServer, ProductUseCase } from "../../core/entity/product";

class ProductServiceImpl implements IProductServiceServer {
    constructor(
        private productUseCase: ProductUseCase
    ) { }

    async fetch(call: ServerUnaryCall<FetchRequest, FetchResponse>, callback: sendUnaryData<FetchResponse>) {
        try {
            const products = await this.productUseCase.fetch()
            callback(null, FetchResponse.fromJSON({ Products: products }));
        } catch (error: any) {
            return callback(error.toString(), null);
        }
    }
}

export { ProductServiceImpl }