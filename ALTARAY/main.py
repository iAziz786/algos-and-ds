# Link to the problem: https://www.codechef.com/problems/ALTARAY


def main():
    T = int(input())
    while T:
        T -= 1
        N = int(input())
        A = [int(i) for i in input().split()]
        ans = [1] * N
        for i in range(N - 1, 0, -1):
            if A[i] * A[i - 1] < 0:
                ans[i - 1] += ans[i]
        print(' '.join(map(str, ans)))


if __name__ == "__main__":
    main()
