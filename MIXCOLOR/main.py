# Link to the problem: https://www.codechef.com/MARCH18B/problems/MIXCOLOR


def main():
    T = int(input())
    while T:
        T -= 1
        _ = int(input())
        occr = {}
        colors = list(map(int, input().split()))
        for i in colors:
            occr[i] = 0
        for val in colors:
            occr[val] += 1

        ans = 0
        for _, v in occr.items():
            ans += (v - 1)

        print(ans)


if __name__ == "__main__":
    main()
