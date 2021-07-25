import asyncio
import base64
import time

import websockets
import json

from models import NormDTO
from normalizer import normalization
from utils.env_utils import get_address_and_port, get_ws_url


async def handler(websocket, path):
    counter = 0
    try:
        async for buffer in websocket:
            norm_dto = base64.b64decode(buffer)
            json_data = json.loads(norm_dto.decode('UTF-8'))
            result = NormDTO(json_data['tran_id'], normalization(json_data['data']))
            print(result.to_json())
            await websocket.send(result.to_json())
            counter += 1
            print(counter)
    except asyncio.exceptions.CancelledError:
        print("error")
        return

if __name__ == '__main__':
    address, port = get_address_and_port()
    start_server = websockets.serve(handler, address, port)
    event_loop = asyncio.get_event_loop()
    event_loop.run_until_complete(start_server)
    event_loop.run_forever()
