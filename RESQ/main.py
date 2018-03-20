# Link to the problem: https://www.codechef.com/problems/RESQ

import math


def main():
    T = int(input())
    while T:
        T -= 1
        N = int(input())
        diff = 0
        end = math.floor(math.sqrt(N))
        for i in range(1, end + 1):
            reminder = N % i
            if reminder > 0:
                continue
            div1 = N / i
            absolute = math.floor(abs(div1 - i))
            if i == 1:
                diff = absolute
            else:
                if absolute < diff:
                    diff = absolute
        print(diff)


if __name__ == "__main__":
    main()
