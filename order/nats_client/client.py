# order_service/app/nats_client.py
import asyncio, json, uuid
from nats.aio.client import Client as NATS
from nats.aio.errors import ErrConnectionClosed, ErrTimeout
import os


class NatsClient:
    def __init__(self, urls):
        self._nc = NATS()
        self._urls = urls

    async def connect(self):
        await self._nc.connect(servers=self._urls, allow_reconnect=True,
                               max_reconnect_attempts=-1, reconnect_time_wait=2)

    async def close(self):
        await self._nc.drain()  # flush & close

    async def request(self, subject, payload, timeout=2.0, headers=None):
        data = json.dumps(payload).encode()
        hdr = {}
        if headers:
            from nats.aio.msg import Msg
            # nats-py supports Msg, but headers are passed differently; pattern shown for concept
            # attach headers using nc.request(..., headers=hdr) if lib supports it
        try:
            msg = await self._nc.request(subject, data, timeout=timeout)
            return json.loads(msg.data.decode())
        except asyncio.TimeoutError:
            raise ErrTimeout()
        except Exception as e:
            raise

# usage: instantiate once at startup and attach to app.state.nats
nats_client = NatsClient([f"{os.environ.get('NATS_SERVER')}:{os.environ.get('NATS_PORT')}"])

