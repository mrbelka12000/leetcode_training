class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        product =1
        summ= 0

        while n > 0:
            last_digit = n % 10
            n = n//10
            product = product*last_digit
            summ = summ+product

        return product-summ

def main():

    print(Solution().subtractProductAndSum(235))



if __name__ == '__main__':
    main()

