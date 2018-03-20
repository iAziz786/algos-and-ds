# Link to the problem: https://www.codechef.com/problems/SIMPSTAT


def main():
    T = int(input())
    while T:
        T -= 1
        N, K = [int(x) for x in input().split()]
        A = [int(x) for x in input().split()]
        A.sort()
        A = A[K: N - K]
        print(float(sum(A) / len(A)))


if __name__ == "__main__":
    main()
