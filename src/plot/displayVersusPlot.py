import numpy as np
import matplotlib.pyplot as plt
import sys

colors = [
    "red",
    "blue",
    "green",
    "purple",
    "aqua",
    "brown",
    "black",
    "yellow",
    "gray",
]

arrays = [
    "Selection sort",
    "Bubble sort",
    "Insertion sort",
    "Merge sort",
    "Quick sort",
    "Shell sort",
    "Shell-pratt sort",
    "Shell-hibbard sort",
    "Heap sort"
]

if __name__ == "__main__":
    data = sys.argv[1]

    AllSortedXY = data.split(";")

    for i in range(0, 18, 2):
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
        plt.title("Сравнение средних случаев сортировок")
        plt.legend()
        plt.grid(True)

        plt.savefig(f"image/versus/theory.png")

