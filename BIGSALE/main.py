# Link to the problem: https://www.codechef.com/MARCH18B/problems/BIGSALE


def main():
    T = int(input())
    while T:
        total_loss = float(0)
        T -= 1
        N = int(input())
        while N:
            N -= 1
            pqd = []
            pqd = list(map(int, input().split()))
            inc_price = float(pqd[0] + pqd[0] * pqd[2] / 100)
            dis_price = float(inc_price - (inc_price * pqd[2]) / 100)
            total_loss += (pqd[1] * (pqd[0] - dis_price))
        print(total_loss)


if __name__ == "__main__":
    main()
