import json
import json
import pickle
import socket

from denormalizer import denormalization
from models import DenormDTO
from utils.env_utils import get_scaler_file

if __name__ == '__main__':
    scaler_file = get_scaler_file()
    with open(scaler_file, 'rb') as handle:
        lr = pickle.load(handle)

    recv_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('localhost', 8085)
    print('Starting up on port 8085')
    recv_sock.bind(server_address)
    recv_sock.listen(5)
    counter = 0

    send_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    #send_sock.connect(('localhost', 8084))

    while True:
        recv_conn, client_addr = recv_sock.accept()
        print('connection from', str(client_addr))
        try:
            while True:
                json_data = json.loads(recv_conn.recv(1024).decode('UTF-8'))
                print(json_data)
                if not json_data:
                    break
                counter += 1
                denorm_data = denormalization(lr, json_data['data'])
                result_dto = DenormDTO(json_data['tran_id'], denorm_data)
                print(result_dto)
                result_bytes = bytes(result_dto.to_json(), encoding="UTF-8")
                send_sock.send(result_bytes)
        finally:
            print("in finally")
            recv_conn.close()
            send_sock.close()
