import asyncio
import base64
import json

import websockets

from models import DenormDTO
from denormalizer import denormalization
from utils.env_utils import get_address_and_port


async def handler(websocket, path):
    counter = 0
    try:
        async for buffer in websocket:
            norm_dto = base64.b64decode(buffer)
            json_data = json.loads(norm_dto.decode('UTF-8'))
            result = DenormDTO(json_data['tran_id'], denormalization(json_data['data']))
            print(result.to_json())
            await websocket.send(result.to_json())
            counter += 1
            print(counter)
    except websockets.exceptions.ConnectionClosedError:
        print("error")
        return


if __name__ == '__main__':
    address, port = get_address_and_port()
    event_loop = asyncio.get_event_loop()
    start_server = websockets.serve(handler, address, port, ping_interval=None)
    event_loop.run_until_complete(start_server)
    event_loop.run_forever()
