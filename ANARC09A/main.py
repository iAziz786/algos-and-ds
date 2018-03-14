# Link to the problem: http://www.spoj.com/problems/ANARC09A/


def main():
    n = input()
    count = 1
    while n[0] != '-':
        o, c, ans = 0, 0, 0
        for i in n:
            if i == '{':
                o += 1
            elif i == '}':
                if o > 0:
                    o -= 1
                else:
                    c += 1
        # if (o + c) % 2 == 0:
        ans += (o // 2) + (c // 2)
        if o % 2 == 1 and c % 2 == 1:
            ans += 2
        # else:
        #     pass

        print('{}. {}'.format(count, ans))
        count += 1
        n = input()


if __name__ == '__main__':
    main()
