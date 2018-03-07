# Link to the problem: https://www.codechef.com/problems/COMPILER


def main():
    test = int(input())
    while test:
        test = test - 1
        exp = input()
        count, res = 0, 0
        for i, val in enumerate(exp):

            if val == "<":
                count = count + 1
            else:
                count = count - 1

            if count == 0:
                res = i + 1
            elif count < 0:
                break
        print(res)


if __name__ == "__main__":
    main()
