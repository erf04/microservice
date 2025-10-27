from fastapi.routing import APIRouter
from schemas.order import OrderCreate,OrderOut,OrderUpdate
from schemas.product import ProductFilter
from models.order import OrderModel
from services.product import ProductService
from nats_client.client import nats_client

router = APIRouter(prefix="/order",tags=["orders"])

product_service = ProductService(nats_client)

@router.get("/list", response_model=list[OrderOut])
async def list_orders():
    orders = await OrderModel.get_all()
    return orders

@router.post("/create",response_model=OrderOut)
async def create_order(order:OrderCreate):
    order = order.model_dump()
    order_id = await OrderModel.create(order)
    return OrderOut(id=order_id,**order)





@router.put("/update")
async def update_order(order_detail:OrderUpdate):
    order_updated = await OrderModel.update(order_detail)
    return order_updated


@router.get("/product/{product_id}")
async def get_product(product_id:str):
    response = await product_service.get_product_by_id(product_id=product_id)
    return response

@router.get('/products')
async def get_products(name : str | None = None , price : int | None = None):
    response = await product_service.get_products(body={"name":name,"price":price})
    return response


# @router.get("/user/{user_id}",response_model=list[OrderOut])
# async def get_orders_by_user_id(user_id:int):
#     orders = await OrderModel.get_by_user_id(user_id)
#     return orders