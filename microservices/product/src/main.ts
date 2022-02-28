import { ProductRepositoryImpl } from "./adapter/sqlite/product_repository";
import { ProductServiceImpl } from "./adapter/grpc_service/product_service";
import { ProductUseCaseImpl } from "./core/usecase/product_usecase";
import { ProductServiceService } from "./core/domain/product";
import { Server, ServerCredentials, ServiceDefinition } from "@grpc/grpc-js";
import { database } from "./adapter/sqlite/factory";
import { IProductServiceServer } from "./core/entity/product";

class TypedServerOverride extends Server {
    addServiceTyped<TypedServiceImplementation extends Record<any, any>>(service: ServiceDefinition, implementation: TypedServiceImplementation): void {
        this.addService(service, implementation)
    }
}

const server = new TypedServerOverride({
    'grpc.max_receive_message_length': -1,
    'grpc.max_send_message_length': -1,
});

const productRepository = new ProductRepositoryImpl(database)
const productUseCase = new ProductUseCaseImpl(productRepository)
const productServer = new ProductServiceImpl(productUseCase)

server.addServiceTyped<IProductServiceServer>(ProductServiceService, productServer);

server.bindAsync('localhost:2222', ServerCredentials.createInsecure(), (err: Error | null, bindPort: number) => {
    if (err) {
        throw err;
    }

    server.start();
});