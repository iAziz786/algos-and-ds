# Link to the problem: http://www.spoj.com/problems/BAISED/


def main():
    T = int(input())
    while T:
        T -= 1
        _ = input()
        n = int(input())
        ranks = []
        for i in range(n):
            _, num = list(input().split())
            ranks.append(int(num))
        list.sort(ranks)
        badness = 0
        for i in range(n):
            rank = i + 1
            guess = ranks[i]
            if rank < guess:
                badness += guess - rank
            else:
                badness += rank - guess
        print(badness)


if __name__ == '__main__':
    main()
