import asyncio
import base64
import json
import pickle
import websockets
import nest_asyncio

from common.env import get_regressor_file
from common.prediction import predict2
from utils.models import ResultDTO


async def receive_handler(websocket, path):
    print("here")
    counter = 0
    lr_file = get_regressor_file()
    with open(lr_file, 'rb') as handle:
        lr = pickle.load(handle)
    try:
        async for buffer in websocket:
            norm_dto = base64.b64decode(buffer)
            json_data = json.loads(norm_dto.decode('UTF-8'))
            result = predict2(lr, json_data['data'])[0]
            result_dto = ResultDTO(json_data['tran_id'], json_data['data'], True if result == 1 else False)
            print(result_dto)
            await send_handler(result_dto)
            counter += 1
            print(counter)
    except websockets.exceptions.ConnectionClosedError:
        print("error")
        return


async def send_handler(result_dto):
    await send_socket.send(result_dto.to_json())


async def main():
    nest_asyncio.apply()
    address = 'localhost'
    port = 8083
    uri = "ws://localhost:8084/"
    global send_socket
    send_socket = await websockets.connect(uri)
    event_loop = asyncio.get_event_loop()
    start_server = websockets.serve(receive_handler, address, port, ping_interval=None)
    event_loop.run_until_complete(start_server)
    event_loop.run_forever()


asyncio.run(main())
