# Link to the problem: https://www.codechef.com/problems/INTEST


def main():
    n, k = list(map(int, input().split()))
    ans = 0
    for _ in range(n):
        t = int(input())
        if t % k == 0:
            ans += 1
    print(ans)


if __name__ == "__main__":
    main()
