# Link to the problem: https://www.codechef.com/problems/STACKS


def binary_search(arr, x):
    l = 0
    r = len(arr)

    while l < r:
        mid = (r + l) // 2
        if x < arr[mid]:
            r = mid
        else:
            l = mid + 1

    return l


def main():
    T = int(input())
    while T:
        T -= 1
        _ = int(input())
        given_stack = list(map(int, input().split()))
        top_stacks = []
        for i in given_stack:
            to_push = binary_search(top_stacks, i)
            if to_push == len(top_stacks):
                top_stacks.append(i)
            else:
                top_stacks[to_push] = i
        print(len(top_stacks), end=" ")
        print(" ".join([str(i) for i in top_stacks]))


if __name__ == "__main__":
    main()
