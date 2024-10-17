import numpy as np
import matplotlib.pyplot as plt
from sympy import symbols, Poly
import sys

colors = ["red", "blue", "green", "yellow"]
arrays = ["Случайный массив", "Отсортированный массив", "Обратно отсортированный массив", "Частично отсортированный массив"]

if __name__ == "__main__":
    data = sys.argv[1]
    name = sys.argv[2]

    AllSortedXY = data.split(";")

    for i in range(0, 8, 2):
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

        # Степень полинома для регрессии
        degree = 2

        # Выполняем полиномиальную регрессию с использованием numpy
        coefficients = np.polyfit(x_data, y_data, degree)

        # Получаем полином
        poly = np.poly1d(coefficients)

        # Создаем символ для x
        x = symbols('x')

        # Преобразуем полиномиальные коэффициенты в polynom SymPy
        sympy_poly = Poly(coefficients[::-1], x)

        # Генерируем значения для графика
        x_values = np.linspace(np.min(x_data), np.max(x_data), 100)
        y_values = poly(x_values)

        # Строим график
        plt.scatter(x_data, y_data, color=colors[i//2], label=arrays[i//2])
        plt.plot(x_values, y_values, label=f'Полиномиальная регрессия', color=colors[i//2])
        plt.xlabel('N')
        plt.ylabel('Мс')
        plt.title(name)
        plt.legend()
        plt.grid(True)

        plt.savefig(f"image/{name}.png")

