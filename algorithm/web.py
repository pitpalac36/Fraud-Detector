import asyncio
import base64
import json
import pickle

import websockets

from common.env import get_regressor_file
from common.prediction import predict
from utils.models import ResultDTO


async def handler(websocket, path):
    buffer = await websocket.recv()
    norm_dto = base64.b64decode(buffer)
    json_data = json.loads(norm_dto.decode('UTF-8'))

    lr_file = get_regressor_file()
    with open(lr_file, 'rb') as handle:
        lr = pickle.load(handle)

    prediction = True if predict(lr, json_data['data']) == [1] else False
    result = ResultDTO(json_data['tran_id'], prediction)
    print(result)
    bytes = pickle.dumps(result)
    buffer = base64.b64encode(bytes)
    await websocket.send(buffer)


if __name__ == '__main__':
    start_server = websockets.serve(handler, "localhost", 8082)
    asyncio.get_event_loop().run_until_complete(start_server)
    asyncio.get_event_loop().run_forever()
