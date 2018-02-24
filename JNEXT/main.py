# Link to the problem: http://www.spoj.com/problems/JNEXT/


def main():
    test = int(input())
    while test:
        test = test - 1
        size = int(input())
        challenge = list(map(int, input().split()))
        left = size - 1
        while left > 0 and challenge[left] <= challenge[left - 1]:
            left = left - 1

        if left <= 0:
            print(-1)
            continue

        right = size - 1
        while challenge[right] <= challenge[left - 1]:
            right = right - 1

        temp = challenge[left - 1]
        challenge[left - 1] = challenge[right]
        challenge[right] = temp

        right = size - 1
        while right > left:
            temp = challenge[left]
            challenge[left] = challenge[right]
            challenge[right] = temp
            right = right - 1
            left = left + 1

        for i in challenge:
            print(i, end='')
        print()


if __name__ == "__main__":
    main()
