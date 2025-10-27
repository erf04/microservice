from nats_client.client import NatsClient
from schemas import product

class ProductService:

    def __init__(self,nats_client:NatsClient):
        self.client = nats_client

    async def get_product_by_id(self,product_id:str) -> product.ProductBase:
        # channel = grpc.insecure_channel(f"{os.environ.get('PRODUCTSVC')}:50051")
        # stub = product_pb2_grpc.ProductServiceStub(channel=channel)
        # response = stub.GetProductByID(product_pb2.GetProductByIDRequest(id=product_id))
        # return MessageToDict(response.product, preserving_proto_field_name=True)
        product = await self.client.request("product.getbyid",{"id":product_id})
        return product
        


    async def get_products(self,body:dict):
        products = await self.client.request("product.all",body)
        return products

