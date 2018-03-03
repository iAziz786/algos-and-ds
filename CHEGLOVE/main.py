# Link to the problem: https://www.codechef.com/MARCH18B/problems/CHEGLOVE


def main():
    T = int(input())
    while T:
        T -= 1
        isFront = True
        isBack = True
        N = int(input())
        L = list(map(int, input().split()))
        G = list(map(int, input().split()))
        i, j = 0, N - 1
        while i < N and j >= 0:
            if L[i] > G[i]:
                isFront = False

            if L[i] > G[j]:
                isBack = False

            i, j = i + 1, j - 1

        if isFront and isBack:
            print("both")
        elif isFront and not isBack:
            print("front")
        elif isBack and not isFront:
            print("back")
        else:
            print("none")


if __name__ == "__main__":
    main()
