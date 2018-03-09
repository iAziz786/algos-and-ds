# Link to the problem: https://www.codechef.com/problems/CHEFST


def main():
    T = int(input())
    while T:
        T -= 1
        n1, n2, m = list(map(int, input().split()))
        print(max(n1 + n2 - m * (m + 1), n1 - n2, n2 - n1))


if __name__ == "__main__":
    main()
