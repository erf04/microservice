import os
from dotenv import load_dotenv, find_dotenv
from pymongo import MongoClient
from motor.motor_asyncio import AsyncIOMotorClient
from urllib.parse import quote_plus


# Load .env file only if it exists
load_dotenv(find_dotenv())  # automatically finds .env up the tree

DB_HOST = os.getenv("DB_HOST", "localhost")
DB_PORT = int(os.getenv("DB_PORT", 27017))
DB_USER = os.getenv("DB_USER")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_NAME = os.getenv("DB_NAME", "mydb")


if DB_USER and DB_PASSWORD:
    user = quote_plus(DB_USER)
    password = quote_plus(DB_PASSWORD)
    MONGO_URI = f"mongodb://{user}:{password}@{DB_HOST}:{DB_PORT}/{DB_NAME}?authSource=admin"
else:
    MONGO_URI = f"mongodb://{DB_HOST}:{DB_PORT}/{DB_NAME}"

client = AsyncIOMotorClient(MONGO_URI)
db = client[DB_NAME]


