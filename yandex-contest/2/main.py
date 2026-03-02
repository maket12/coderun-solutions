import sys


def main():
    """
    Пример ввода и вывода числа n, где -10^9 < n < 10^9:
    n = int(input())
    print(n)
    """
    n = int(input())
    arr = list(map(int, input().split()))

    prefix = [0] * n
    for i in range(n):
        if i == 0:
            prefix[i] = arr[i]
        else:
            prefix[i] = prefix[i-1] + arr[i]
    
    for i in range(len(arr) - 1):
        can_win = True
        capital = prefix[i]
        for j in range(i+1, len(arr)):
            if capital > arr[j]:
                capital += arr[j]
            else:
                can_win = False
                break
        print(int(can_win))
    print(1)


if __name__ == '__main__':
    main()
