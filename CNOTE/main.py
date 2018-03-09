def main():
    T = int(input())
    while T:
        T -= 1
        X, Y, K, N = list(map(int, input().split()))
        lucky = False
        for _ in range(N):
            P, C = list(map(int, input().split()))
            if lucky == False and P >= (X - Y) and C <= K:
                lucky = True
        if lucky:
            print("LuckyChef")
        else:
            print("UnluckyChef")


if __name__ == "__main__":
    main()
