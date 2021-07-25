import asyncio
import base64
import json
import pickle
import socket
import sys

from common.env import get_regressor_file
from common.prediction import predict
from utils.models import ResultDTO




if __name__ == '__main__':
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('localhost', 8083)
    print('Starting up on port 8083')
    sock.bind(server_address)
    sock.listen(1)
    counter = 0
    while True:
        socket, client_addr = sock.accept()
        print('connection from', str(client_addr))
        try:
            while True:
                data = socket.recv(1024).decode('utf-8')
                counter += 1
                print('received ', str(data))
                print(counter)
        finally:
            socket.close()