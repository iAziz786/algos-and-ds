# Link to the problem: https://www.codechef.com/OCT18B/problems/CHSERVE
def CHSERVE():
  p1, p2, k = list(map(int, input().split()))
  switches = (p1 + p2) // k
  if (switches % 2 == 0):
    return "CHEF"
  return "COOK"

def main():
  T = int(input())
  while T:
    T -= 1
    print(CHSERVE())

if __name__ == '__main__':
  main()