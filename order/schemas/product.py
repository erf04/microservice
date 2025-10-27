from pydantic import BaseModel


class ProductBase(BaseModel):
    id :str 
    price:float
    name:str


class ProductFilter(BaseModel):
    price: float | None = None
    name : str | None = None

