import asyncio
import base64
import pickle

import websockets
import json

from models import NormDTO
from normalizer import normalization


async def hello(websocket, path):
    buffer = await websocket.recv()
    norm_dto = base64.b64decode(buffer)
    json_data = json.loads(norm_dto.decode('UTF-8'))
    print(NormDTO(json_data['tran_id'], normalization(json_data['data'])))
    base64.b64encode(pickle.dumps(NormDTO(json_data["tran_id"], normalization(json_data['data']))), buffer)
    await websocket.send(bytes(buffer))


def main():
    start_server = websockets.serve(hello, "localhost", 8082)
    asyncio.get_event_loop().run_until_complete(start_server)
    asyncio.get_event_loop().run_forever()


main()
