class Solution:
    def maxFreqSum(self, s: str) -> int:
        vowels = ['a', 'e', 'i', 'o', 'u']
        vowel_count = {}
        consonant_count = {}
        for i in range(len(s)):
            if s[i] in vowels:
                vowel_count[s[i]] = vowel_count.get(s[i], 0) + 1
            else:
                consonant_count[s[i]] = consonant_count.get(s[i], 0) + 1


        result = 0

        vowels_arr = []
        consonants_arr = []
        for val in vowel_count.values():
            vowels_arr.append(val)

        for val in consonant_count.values():
            consonants_arr.append(val)

        sorted_vowels = sorted(vowels_arr, reverse=True)
        sorted_consonants = sorted(consonants_arr,reverse=True)

        return sorted_consonants[0] + sorted_vowels[0]


def main():
    print(Solution().maxFreqSum("successes"))


if __name__ == '__main__':
    main()