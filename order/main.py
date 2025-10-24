from typing import Union

from fastapi import FastAPI
from routes import order
import asyncio
import nats
import os
from nats_client.client import nats_client

app = FastAPI()
app.include_router(order.router)



@app.on_event("startup")
async def startup():
    await nats_client.connect()

@app.on_event("shutdown")
async def shutdown():
    await nats_client.close()

# async def main():
#     # Connect to NATS!
#     nc = await nats.connect(f"{os.environ.get('NATS_SERVER')}:{os.environ.get('NATS_PORT')}")

#     # Receive messages on 'foo'
#     sub = await nc.subscribe("foo")

#     # Publish a message to 'foo'
#     await nc.publish("foo", b'Hello from Python!')

#     # Process a message
#     msg = await sub.next_msg()
#     print("Received:", msg)

#     # Make sure all published messages have reached the server
#     await nc.flush()

#     # Close NATS connection
#     await nc.close()


@app.get("/")
def read_root():
    # asyncio.run(main())
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}