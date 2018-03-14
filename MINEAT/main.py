# Link to the problem: https://www.codechef.com/problems/MINEAT

import math


def main():
    T = int(input())
    while T:
        T -= 1
        _, H = list(map(int, input().split()))
        A = list(map(int, input().split()))
        l, r = 1, max(A)
        while l < r:
            mid = (l + r) // 2
            hours_taken = 0
            for i in A:
                hours_taken += math.ceil(i / mid)
            if hours_taken <= H:
                r = mid
            elif hours_taken > H:
                l = mid + 1
        print(l)


if __name__ == "__main__":
    main()
