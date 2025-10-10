from pydantic import BaseModel


class ProductBase(BaseModel):
    id :str 
    price:float
    name:str


