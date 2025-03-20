#ifndef ARRAY_HPP
#define ARRAY_HPP

#include <vector>
#include <bits/stdc++.h>

class Solution {
public:
    bool containsDuplicate(std::vector<int>& nums);
    bool isAnagram(std::string s, std::string t);
    std::vector<int> twoSum(std::vector<int>& nums, int target);
    std::vector<std::vector<std::string>> groupAnagrams(std::vector<std::string>& strs);
    std::vector<int> topKFrequent(std::vector<int>& nums, int k);
    std::vector<int> productExceptSelf(std::vector<int>& nums);
    bool isValidSudoku(std::vector<std::vector<char>>& board);
};
#endif
