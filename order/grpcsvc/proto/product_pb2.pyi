from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Product(_message.Message):
    __slots__ = ("id", "name", "price")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    price: float
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., price: _Optional[float] = ...) -> None: ...

class GetProductByIDRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetProductResponse(_message.Message):
    __slots__ = ("product",)
    PRODUCT_FIELD_NUMBER: _ClassVar[int]
    product: Product
    def __init__(self, product: _Optional[_Union[Product, _Mapping]] = ...) -> None: ...

class GetProductsRequest(_message.Message):
    __slots__ = ("name", "price")
    NAME_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    name: str
    price: float
    def __init__(self, name: _Optional[str] = ..., price: _Optional[float] = ...) -> None: ...

class GetProductsResponse(_message.Message):
    __slots__ = ("products",)
    PRODUCTS_FIELD_NUMBER: _ClassVar[int]
    products: _containers.RepeatedCompositeFieldContainer[GetProductResponse]
    def __init__(self, products: _Optional[_Iterable[_Union[GetProductResponse, _Mapping]]] = ...) -> None: ...
