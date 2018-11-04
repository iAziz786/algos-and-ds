# Link to the problem: https://www.codechef.com/NOV18B/problems/CHHAPPY


def main():
  T = int(input())
  while T:
    T -= 1
    _ = input()
    A = list(map(int, input().split()))
    ans = "Poor Chef"
    previousVisits = dict()
    for i in range(len(A)):
      # Do you have A[A[i]] ?
      # if yes ? what is the value it points to : put this there
      # Indication: points_to = index of the value
      # if A[points_to] != A[i] and A[A[points_to]] == A[A[i]]
      # Hurray result!
      current = A[A[i] - 1]
      previousValue = previousVisits.get(current)
      if (previousValue == None):
        previousVisits.update({current: i})
      else:
        if (A[previousValue] != A[i]):
          ans = "Truly Happy"
    print(ans)

if __name__ == '__main__':
  main()
