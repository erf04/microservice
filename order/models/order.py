from database import db
from schemas.order import OrderCreate,OrderUpdate,OrderOut
from bson import ObjectId

class OrderModel:
    collection = db["order"]

    @staticmethod
    async def create(order:OrderCreate):
        result = await OrderModel.collection.insert_one(order)
        return str(result.inserted_id)

    @staticmethod
    async def get_all():
        orders = []
        async for doc in OrderModel.collection.find({}):
            doc["id"] = str(doc["_id"])
            del doc["_id"]
            orders.append(doc)
        return orders
    

    @staticmethod
    async def get_by_user_id(user_id:int):
        orders = []
        async for doc in OrderModel.collection.find({"user_id":user_id}):
            doc["id"] = str(doc["_id"])
            del doc["_id"]
            orders.append(doc)
        return orders
    

    @staticmethod
    async def update(order:OrderUpdate):
        order_dict = order.model_dump()
        order_id = order_dict.pop('id')

        # Convert string to ObjectId
        if not ObjectId.is_valid(order_id):
            return 0
        

        order_updated = await OrderModel.collection.update_one({"_id":ObjectId(order_id)},{"$set":order_dict})
        if order_updated.matched_count == 0:
            return None  # no doc found
        updated_doc = await OrderModel.collection.find_one({"_id": ObjectId(order_id)})
        if updated_doc:
            updated_doc["id"] = str(updated_doc["_id"])
            del updated_doc["_id"]
        return updated_doc
    
    @staticmethod
    async def find(filter:dict) -> list[OrderOut]:
        orders = []
        async for doc in OrderModel.collection.find(filter):
            doc["id"] = str(doc["_id"])
            del doc["_id"]
            orders.append(doc)
        return orders

    
