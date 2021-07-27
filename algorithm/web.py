import socket

if __name__ == '__main__':
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('localhost', 8083)
    print('Starting up on port 8083')
    sock.bind(server_address)
    sock.listen(5)
    counter = 0
    while True:
        socket, client_addr = sock.accept()
        print('connection from', str(client_addr))
        try:
            while True:
                data = socket.recv(800).decode('UTF-8')
                if not data:
                    break
                counter += 1
                print('received ', str(data))
                print(counter)
        finally:
            print("in finally")
            socket.close()
