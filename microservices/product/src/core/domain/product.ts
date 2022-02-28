/* eslint-disable */
import Long from "long";
import {
  makeGenericClientConstructor,
  ChannelCredentials,
  ChannelOptions,
  UntypedServiceImplementation,
  handleUnaryCall,
  Client,
  ClientUnaryCall,
  Metadata,
  CallOptions,
  ServiceError,
} from "@grpc/grpc-js";
import _m0 from "protobufjs/minimal";

export interface FetchRequest {}

export interface FetchResponse {
  Products: Product[];
}

export interface Product {
  id: number;
  name: string;
  price: number;
}

function createBaseFetchRequest(): FetchRequest {
  return {};
}

export const FetchRequest = {
  encode(
    _: FetchRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FetchRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFetchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): FetchRequest {
    return {};
  },

  toJSON(_: FetchRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FetchRequest>, I>>(
    _: I
  ): FetchRequest {
    const message = createBaseFetchRequest();
    return message;
  },
};

function createBaseFetchResponse(): FetchResponse {
  return { Products: [] };
}

export const FetchResponse = {
  encode(
    message: FetchResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.Products) {
      Product.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FetchResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFetchResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Products.push(Product.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FetchResponse {
    return {
      Products: Array.isArray(object?.Products)
        ? object.Products.map((e: any) => Product.fromJSON(e))
        : [],
    };
  },

  toJSON(message: FetchResponse): unknown {
    const obj: any = {};
    if (message.Products) {
      obj.Products = message.Products.map((e) =>
        e ? Product.toJSON(e) : undefined
      );
    } else {
      obj.Products = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FetchResponse>, I>>(
    object: I
  ): FetchResponse {
    const message = createBaseFetchResponse();
    message.Products =
      object.Products?.map((e) => Product.fromPartial(e)) || [];
    return message;
  },
};

function createBaseProduct(): Product {
  return { id: 0, name: "", price: 0 };
}

export const Product = {
  encode(
    message: Product,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.price !== 0) {
      writer.uint32(29).float(message.price);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Product {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProduct();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.price = reader.float();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Product {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      price: isSet(object.price) ? Number(object.price) : 0,
    };
  },

  toJSON(message: Product): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Product>, I>>(object: I): Product {
    const message = createBaseProduct();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.price = object.price ?? 0;
    return message;
  },
};

export const ProductServiceService = {
  fetch: {
    path: "/productservice.ProductService/Fetch",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: FetchRequest) =>
      Buffer.from(FetchRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => FetchRequest.decode(value),
    responseSerialize: (value: FetchResponse) =>
      Buffer.from(FetchResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => FetchResponse.decode(value),
  },
} as const;

export interface ProductServiceServer  {
  fetch: handleUnaryCall<FetchRequest, FetchResponse>;
}

export interface ProductServiceClient extends Client {
  fetch(
    request: FetchRequest,
    callback: (error: ServiceError | null, response: FetchResponse) => void
  ): ClientUnaryCall;
  fetch(
    request: FetchRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: FetchResponse) => void
  ): ClientUnaryCall;
  fetch(
    request: FetchRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: FetchResponse) => void
  ): ClientUnaryCall;
}

export const ProductServiceClient = makeGenericClientConstructor(
  ProductServiceService,
  "productservice.ProductService"
) as unknown as {
  new (
    address: string,
    credentials: ChannelCredentials,
    options?: Partial<ChannelOptions>
  ): ProductServiceClient;
  service: typeof ProductServiceService;
};

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
