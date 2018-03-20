# Link to the problem: https://www.codechef.com/problems/CHEFSOCK


def main():
    jacketCost, sockCost, money = [int(x) for x in input().split()]
    money -= jacketCost
    if (money // sockCost) % 2 == 1:
        print("Unlucky Chef")
    else:
        print("Lucky Chef")


if __name__ == "__main__":
    main()
