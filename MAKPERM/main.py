# Link to the problem: https://www.codechef.com/COOK98B/problems/MAKPERM

def main():
  T = int(input())
  while T:
    T -= 1
    N = int(input())
    A = list(map(int, input().split()))
    mask = [False] * N
    for i in A:
      if i <= N and not mask[i - 1]:
        mask[i - 1] = True
    print(mask.count(False))


if __name__ == '__main__':
  main()