# Link to the problem: https://www.codechef.com/LTIME63B/problems/EID

def EID():
  _ = input()
  A = list(map(int, input().split()))
  A.sort()
  size = len(A)
  diff = A[size - 1]
  for i in range(size - 1):
    tempDiff = A[i + 1] - A[i]
    if tempDiff < diff:
      diff = tempDiff
  return diff

def main():
  T = int(input())
  while T:
    T -= 1
    res = EID()
    print(res)

if __name__ == "__main__":
  main()