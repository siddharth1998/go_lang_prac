import socket
import time


socket_client=socket.socket(socket.AF_INET,socket.SOCK_STREAM)

socket_client.connect(("127.0.0.1",2080))
count =0
while(True):
    socket_client.sendall(b"Hello there")
    count+=1
    print("Sent "+str(count))
    if count>10000:
        break
    time.sleep(0.01)
socket_client.close()