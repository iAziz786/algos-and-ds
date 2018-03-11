# Link to the problem: http://www.spoj.com/problems/FASHION/


def main():
    T = int(input())
    while T:
        T -= 1
        N = int(input())
        m = list(map(int, input().split()))
        w = list(map(int, input().split()))
        h = 0
        list.sort(m)
        list.sort(w)
        for i in range(N):
            h += m[i] * w[i]
        print(h)


if __name__ == '__main__':
    main()
