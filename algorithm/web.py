import asyncio
import base64
import json
import pickle
import websockets

from common.env import get_regressor_file
from common.prediction import predict2
from utils.models import ResultDTO


async def handler(websocket, path):
    counter = 0
    lr_file = get_regressor_file()
    with open(lr_file, 'rb') as handle:
        lr = pickle.load(handle)
    try:
        async for buffer in websocket:
            norm_dto = base64.b64decode(buffer)
            json_data = json.loads(norm_dto.decode('UTF-8'))
            result = predict2(lr, json_data['data'])[0]
            result_dto = ResultDTO(json_data['tran_id'], True if result == 1 else False)
            print(result_dto)
            await websocket.send(result_dto.to_json())
            counter += 1
            print(counter)
    except websockets.exceptions.ConnectionClosedError:
        print("error")
        return


if __name__ == '__main__':
    address = 'localhost'
    port = 8083
    event_loop = asyncio.get_event_loop()
    start_server = websockets.serve(handler, address, port, ping_interval=None)
    event_loop.run_until_complete(start_server)
    event_loop.run_forever()
