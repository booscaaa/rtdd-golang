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
  People: Person[];
}

export interface Person {
  id: number;
  name: string;
  age: number;
}

export interface GetByIDRequest {
  id: number;
}

export interface GetByIDResponse {
  Person?: Person;
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
  return { People: [] };
}

export const FetchResponse = {
  encode(
    message: FetchResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.People) {
      Person.encode(v!, writer.uint32(10).fork()).ldelim();
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
          message.People.push(Person.decode(reader, reader.uint32()));
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
      People: Array.isArray(object?.People)
        ? object.People.map((e: any) => Person.fromJSON(e))
        : [],
    };
  },

  toJSON(message: FetchResponse): unknown {
    const obj: any = {};
    if (message.People) {
      obj.People = message.People.map((e) =>
        e ? Person.toJSON(e) : undefined
      );
    } else {
      obj.People = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FetchResponse>, I>>(
    object: I
  ): FetchResponse {
    const message = createBaseFetchResponse();
    message.People = object.People?.map((e) => Person.fromPartial(e)) || [];
    return message;
  },
};

function createBasePerson(): Person {
  return { id: 0, name: "", age: 0 };
}

export const Person = {
  encode(
    message: Person,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.age !== 0) {
      writer.uint32(24).int32(message.age);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Person {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePerson();
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
          message.age = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Person {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      age: isSet(object.age) ? Number(object.age) : 0,
    };
  },

  toJSON(message: Person): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.age !== undefined && (obj.age = Math.round(message.age));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Person>, I>>(object: I): Person {
    const message = createBasePerson();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.age = object.age ?? 0;
    return message;
  },
};

function createBaseGetByIDRequest(): GetByIDRequest {
  return { id: 0 };
}

export const GetByIDRequest = {
  encode(
    message: GetByIDRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetByIDRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetByIDRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetByIDRequest {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: GetByIDRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetByIDRequest>, I>>(
    object: I
  ): GetByIDRequest {
    const message = createBaseGetByIDRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseGetByIDResponse(): GetByIDResponse {
  return { Person: undefined };
}

export const GetByIDResponse = {
  encode(
    message: GetByIDResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.Person !== undefined) {
      Person.encode(message.Person, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetByIDResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetByIDResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Person = Person.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetByIDResponse {
    return {
      Person: isSet(object.Person) ? Person.fromJSON(object.Person) : undefined,
    };
  },

  toJSON(message: GetByIDResponse): unknown {
    const obj: any = {};
    message.Person !== undefined &&
      (obj.Person = message.Person ? Person.toJSON(message.Person) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetByIDResponse>, I>>(
    object: I
  ): GetByIDResponse {
    const message = createBaseGetByIDResponse();
    message.Person =
      object.Person !== undefined && object.Person !== null
        ? Person.fromPartial(object.Person)
        : undefined;
    return message;
  },
};

export const PersonServiceService = {
  fetch: {
    path: "/personservice.PersonService/Fetch",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: FetchRequest) =>
      Buffer.from(FetchRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => FetchRequest.decode(value),
    responseSerialize: (value: FetchResponse) =>
      Buffer.from(FetchResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => FetchResponse.decode(value),
  },
  getByID: {
    path: "/personservice.PersonService/GetByID",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: GetByIDRequest) =>
      Buffer.from(GetByIDRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => GetByIDRequest.decode(value),
    responseSerialize: (value: GetByIDResponse) =>
      Buffer.from(GetByIDResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => GetByIDResponse.decode(value),
  },
} as const;

export interface PersonServiceServer extends UntypedServiceImplementation {
  fetch: handleUnaryCall<FetchRequest, FetchResponse>;
  getByID: handleUnaryCall<GetByIDRequest, GetByIDResponse>;
}

export interface PersonServiceClient extends Client {
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
  getByID(
    request: GetByIDRequest,
    callback: (error: ServiceError | null, response: GetByIDResponse) => void
  ): ClientUnaryCall;
  getByID(
    request: GetByIDRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: GetByIDResponse) => void
  ): ClientUnaryCall;
  getByID(
    request: GetByIDRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: GetByIDResponse) => void
  ): ClientUnaryCall;
}

export const PersonServiceClient = makeGenericClientConstructor(
  PersonServiceService,
  "personservice.PersonService"
) as unknown as {
  new (
    address: string,
    credentials: ChannelCredentials,
    options?: Partial<ChannelOptions>
  ): PersonServiceClient;
  service: typeof PersonServiceService;
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
