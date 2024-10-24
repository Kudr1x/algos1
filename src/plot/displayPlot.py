import numpy as np
import matplotlib.pyplot as plt
import sys

colors = ["red", "blue", "green"]
arrays = ["Best case", "Average case", "Wurst case"]

if __name__ == "__main__":
    data = sys.argv[1]
    name = sys.argv[2]

    AllSortedXY = data.split(";")

    for i in range(0, 6, 2):
        arrXStr = AllSortedXY[i].split(",")
        arrYStr = AllSortedXY[i+1].split(",")

        arrXInt = []
        arrYInt = []

        for j in arrXStr:
            arrXInt.append(int(j))

        for j in arrYStr:
            arrYInt.append(int(j))

        x_data = np.array(arrXInt)
        y_data = np.array(arrYInt)

        # Строим график
        plt.plot(arrXInt, arrYInt, marker='o', color=colors[i//2], label=arrays[i//2])
        plt.xlabel('N')
        plt.ylabel('t')
        plt.title(name)
        plt.legend()
        plt.grid(True)

        plt.savefig(f"image/theory/{name}.png")

