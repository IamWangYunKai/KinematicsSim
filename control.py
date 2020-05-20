from informer import Informer
from time import sleep
import random
import json

# if __name__ == '__main__':
#     ifm = Informer(random.randint(100000,999999), block=False)
#     cnt = 0
#     while True:
#         ifm.send_message(json.dumps({"V":1.0, "W":0.0}), mtype='cmd')
#         ifm.send_message("", mtype='step')
#         sleep(1/30.)
#         cnt += 1

if __name__ == '__main__':
    N = 10
    ifms = []
    for i in range(10):
        ifm = Informer(random.randint(100000,999999), block=False)
        ifms.append(ifm)

    while True:
        for i in range(10):
            ifms[i].send_message(json.dumps({"V":1.0, "W":0.0}), mtype='cmd')

        ifms[0].send_message("", mtype='step')
        sleep(1/30.)
