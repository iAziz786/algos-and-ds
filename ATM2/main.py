# Link to the problem: https://www.codechef.com/COOK98B/problems/ATM2

def main():
  T = int(input())
  while T:
    T -= 1
    N, K = list(map(int, input().split()))
    A = list(map(int, input().split()))
    ans = ''
    for i in range(N):
      if (K - A[i] >= 0):
        ans += '1'
        K -= A[i]
      else:
        ans += '0'
    print(ans)

if __name__ == '__main__':
  main()