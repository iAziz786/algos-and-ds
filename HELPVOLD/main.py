# Link to the problem: https://www.codechef.com/problems/HELPVOLD


def main():
    N = int(input())
    r = [int(x) for x in input().split()]
    r.sort()
    ee = 0
    for i in range(1, N):
        ee += (r[i] * (r[i - 1]))
    print(ee)


if __name__ == "__main__":
    main()
