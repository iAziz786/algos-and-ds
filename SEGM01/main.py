# Link to the problem: https://www.codechef.com/problems/SEGM01


def main():
    T = int(input())
    while T:
        T -= 1
        s = input()
        findAny, getZero = False, False
        for i in s:
            if getZero and findAny and i == "1":
                findAny = False
                break
            elif findAny and i == "0":
                getZero = True
            elif i == "1":
                findAny = True

        if findAny:
            print("YES")
        else:
            print("NO")


if __name__ == "__main__":
    main()
