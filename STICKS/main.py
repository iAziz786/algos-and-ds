# Link to the problem: https://www.codechef.com/problems/STICKS


def main():
    T = int(input())
    while T:
        T -= 1
        N = int(input())
        s = list(map(int, input().split()))
        list.sort(s)
        if N < 4:
            return -1
        width = 0
        height = 0
        size = N - 1
        while size:
            if s[size] == s[size - 1] and width == 0:
                width = s[size]
                size -= 2
                continue
            if s[size] == s[size - 1] and width > 0 and height == 0:
                height = s[size]
                break
            size -= 1
        if width > 0 and height > 0:
            print(width * height)
        else:
            print(-1)


if __name__ == '__main__':
    main()
