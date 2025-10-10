import grpc
import grpcsvc.proto.product_pb2 as product_pb2
import grpcsvc.proto.product_pb2_grpc as product_pb2_grpc
import os
from google.protobuf.json_format import MessageToDict

async def get_product_by_id(product_id:str):
    channel = grpc.insecure_channel(f"{os.environ.get('PRODUCTSVC')}:50051")
    stub = product_pb2_grpc.ProductServiceStub(channel=channel)
    response = stub.GetProductByID(product_pb2.GetProductByIDRequest(id=product_id))
    return MessageToDict(response.product, preserving_proto_field_name=True)


async def get_products(body:dict):
    channel = grpc.insecure_channel(f"{os.environ.get('PRODUCTSVC')}:50051")
    stub = product_pb2_grpc.ProductServiceStub(channel=channel)
    response = stub.GetProducts(product_pb2.GetProductsRequest(**body))
    return MessageToDict(response, preserving_proto_field_name=True)
