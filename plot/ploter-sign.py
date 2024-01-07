
import faulthandler;
faulthandler.enable()

import matplotlib

matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib import font_manager
import numpy as np
import pandas as pd

def appendArr(arr, num, a, base):
    # cost = ax + b
    for i in range(num):
        if(i % 2 == 1):
            element = i * a + base
            arr.append(element)

sig_order = [1, 2, 3, 4, 5, 6, 7, 8]
sakai = [18.32, 41.0, 41.0, 41.0, 41.0, 41.0, 41.0, 41.0]
ibsas = [13.1, 70.7, 94.1, 117.1, 140.8, 163.3, 188.2, 211.4]
OMS = [38.0, 78.12, 79.20, 79.21, 78.93, 78.71, 79.85, 79.58]
WSA = [28.37, 66.38, 89.44, 112.99, 137.11, 160.83, 184.79, 207.44]

fig, ax = plt.subplots()

plt.plot(sig_order, sakai,  color='green', label="Sakai", marker = "*")
plt.plot(sig_order, ibsas,  color='orange', label="IBSAS", marker = "^")
plt.plot(sig_order, OMS,  color='red', label="CDH OMS", marker = ".")
plt.plot(sig_order, WSA,  color='blue', label="WSA", marker = "x")
ax.get_yaxis().get_major_formatter().set_scientific(False)
plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
plt.xlabel("Index of signers")  # 横坐标名字
plt.ylabel("Sign and verify time (ms)")  # 纵坐标名字
plt.legend()
my_x_ticks = np.arange(0, 9, 1)
my_y_ticks = np.arange(0, 241, 40)
plt.xticks(my_x_ticks)
plt.yticks(my_y_ticks)
plt.grid()
fig.savefig('./figures/链下签名及验证时间.svg', dpi=3200, format='svg')





