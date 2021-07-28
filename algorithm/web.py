import json
import pickle
import socket

from common.env import get_regressor_file
from common.prediction import predict2
from utils.models import ResultDTO

if __name__ == '__main__':
    lr_file = get_regressor_file()
    with open(lr_file, 'rb') as handle:
        lr = pickle.load(handle)

    recv_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('localhost', 8083)
    print('Starting up on port 8083')
    recv_sock.bind(server_address)
    recv_sock.listen(5)
    counter = 0

    send_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    send_sock.connect(('localhost', 8084))

    while True:
        recv_conn, client_addr = recv_sock.accept()
        print('connection from', str(client_addr))
        try:
            while True:
                json_data = json.loads(recv_conn.recv(642).decode('UTF-8'))
                print(json_data)
                if not json_data:
                    break
                counter += 1
                result = predict2(lr, json_data['data'])[0]
                result_dto = ResultDTO(json_data['tran_id'], True if result == 1 else False)
                print(result_dto)
                result_bytes = bytes(result_dto.to_json(), encoding="UTF-8")
                send_sock.send(result_bytes)
        finally:
            print("in finally")
            recv_conn.close()
            send_sock.close()
