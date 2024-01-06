
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

sig_nums = [2, 3, 4, 5, 6, 7, 8]
sakai = [542775, 691222, 839647, 988096, 1136555, 1284981, 1433428]
ibsas = [752595, 970787, 1181384, 1397028, 1618147, 1845781, 2078198]
notbatch = [688414, 1019556, 1350712, 1681853, 2013033, 2344198, 2675340]

fig, ax = plt.subplots()

plt.plot(sig_nums, sakai,  color='green', label="Sakai", marker = "*")
plt.plot(sig_nums, ibsas,  color='orange', label="IBSAS", marker = "^")
plt.plot(sig_nums, notbatch,  color='red', label="Sakai without batch verification", marker = "^")

ax.get_yaxis().get_major_formatter().set_scientific(False)
plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
plt.xlabel("Number of signers")  # 横坐标名字
plt.ylabel("Gas consumption")  # 纵坐标名字
plt.legend()
my_x_ticks = np.arange(0, 10, 1)
my_y_ticks = np.arange(0, 3000001, 500000)
plt.xticks(my_x_ticks)
plt.yticks(my_y_ticks)
plt.grid()
fig.savefig('./figures/签名验证gas消耗对比.svg', dpi=3200, format='svg')





