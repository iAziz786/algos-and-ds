# Link to the problem: https://www.codechef.com/problems/ALTARAY


def main():
    T = int(input())
    while T:
        T -= 1
        N = int(input())
        A = list(map(int, input().split()))
        pc, mpc = 0, 0
        s = -1
        ss = -1
        for i in range(N - 1):
            if (A[i] > 0 and A[i + 1] < 0) or (A[i] < 0 and A[i + 1] > 0):
                if pc == 0:
                    s = i
                pc += 1
                if pc > mpc:
                    mpc = pc
                    ss = s
            else:
                pc = 0
        ans = [1] * (N)
        if mpc < 1:
            print(' '.join(map(str, ans)))
        else:
            mpc += 1
            for i in range(ss, (ss + mpc)):
                ans[i] = mpc
                mpc -= 1
            print(' '.join(map(str, ans)))


if __name__ == '__main__':
    main()
