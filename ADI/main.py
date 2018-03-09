# Link to the problem: https://www.codechef.com/COUT2018/problems/ADI


def main():
    N = int(input())
    ans = []
    while N:
        N -= 1
        string = input()
        if "a" in string and "e" in string and "i" in string and "o" in string and "u" in string:
            ans.append(string)
    if ans:
        for i in ans:
            print(i)


if __name__ == "__main__":
    main()
