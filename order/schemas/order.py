from pydantic import BaseModel, EmailStr
from typing import Optional

class OrderBase(BaseModel):
    user_id : int
    product_ids : list[int]

class OrderCreate(OrderBase):
    pass

class OrderUpdate(OrderBase):
    id:str

class OrderOut(OrderBase):
    id: str


# class OrderWithProduct(OrderOut):
#     products: list[ProductOut]
