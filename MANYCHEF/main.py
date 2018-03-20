# Link to the problem: https://www.codechef.com/problems/MANYCHEF


def main():
    T = int(input())
    while T:
        T -= 1
        ss = input()
        s = [i for i in ss]
        size = len(s)
        i = size - 4
        while i >= 0:
            if (s[i] == "?" or s[i] == "C") and (s[i + 1] == "?" or s[i + 1] == "H") and (s[i + 2] == "?" or s[i + 2] == "E") and (s[i + 3] == "?" or s[i + 3] == "F"):
                s[i] = "C"
                s[i + 1] = "H"
                s[i + 2] = "E"
                s[i + 3] = "F"
                i -= 4
            else:
                i -= 1
        for i in range(len(s)):
            if s[i] == "?":
                s[i] = "A"
        print(''.join(map(str, s)))


if __name__ == "__main__":
    main()
