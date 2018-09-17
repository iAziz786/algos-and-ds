# Link to the problem: https://www.codechef.com/SEPT18B/problems/CHEFADV

def CHEFADV():
    knowladge, power = 1, 1
    N, M, X, Y = list(map(int, input().split()))
    N -= knowladge
    M -= power
    if (N % X == 0) and (M % Y == 0):
        return "Chefirnemo"
    elif ((N >= 1) and (M >= 1)) and (((N-1) % X == 0) and ((M-1) % Y == 0)):
        return "Chefirnemo"
    return "Pofik"


def main():
    T = int(input())
    while T:
        T -= 1
        ans = CHEFADV()
        print(ans)


if __name__ == "__main__":
    main()
