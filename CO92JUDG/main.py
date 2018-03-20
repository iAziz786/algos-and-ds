# Link to the problem: https://www.codechef.com/problems/CO92JUDG


def main():
    T = int(input())
    while T:
        T -= 1
        _ = int(input())
        A = [int(x) for x in input().split()]
        B = [int(x) for x in input().split()]
        asum = sum(A) - max(A)
        bsum = sum(B) - max(B)
        if asum < bsum:
            print("Alice")
        elif bsum < asum:
            print("Bob")
        else:
            print("Draw")


if __name__ == "__main__":
    main()
