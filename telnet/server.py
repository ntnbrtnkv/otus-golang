#!/usr/bin/env python3

import socket

# Read "distributed echo server" as "(distributed echo) server". The "server"
# is not "distributed" but the echos are "distributed" to every connected
# client.

# Connect to the server with `telnet localhost 5000`.

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.setblocking(False)
server.bind(('localhost', 8001))
server.listen(5)

connections = []

while True:
    try:
        connection, address = server.accept()
        print("New connection {}".format(address))
        connection.setblocking(False)
        connections.append(connection)
    except BlockingIOError:
        pass

    for connection in connections:
        try:
            message = connection.recv(4096)
        except BlockingIOError:
            continue
        print("New message from {}: {}".format(connection, message))
        for connection in connections:
            connection.send(bytes("Server response: {}".format(message), 'ascii'))