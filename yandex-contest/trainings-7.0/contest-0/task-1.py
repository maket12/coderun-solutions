n = int(input())

if n == 1:
    print(2)
elif n == 2:
    print(4)
elif n == 3:
    print(7)
else:
    f1, f2, f3 = 2, 4, 7
    for i in range(4, n + 1):
        current = f1 + f2 + f3
        f1, f2, f3 = f2, f3, current

    print(f3)
