import time
import socket

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect(("localhost", 4510))
    for i in range(10):
        if i == 9:
            s.sendall(b"Quit\n")
        else:
            s.sendall(b"Hi mom\n")
        time.sleep(1)
