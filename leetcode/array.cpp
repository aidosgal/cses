#include "array.hpp"
#include <string.h>

bool Solution::containsDuplicate(std::vector<int>& nums) {
    std::unordered_set<int> seen;
    for (int num : nums) {
        if(seen.count(num)) {
            return true;
        }

        seen.insert(num);
    }
    return false;
}

bool Solution::isAnagram(std::string s, std::string t) {
    std::unordered_map<char, int> count;
    if (s.size() != t.size()) {
        return false;
    }
    
    for (char c : s) {
        count[c]++;
    }

    for (char c : t ) {
        if (count.find(c) == count.end() || count[c] == 0) {
            return false;
        }
        count[c]--;
    }
    return true;
}

std::vector<int> Solution::twoSum(std::vector<int>& nums, int target) {
    std::unordered_map<int, int> map;

    for (int i = 0; i < nums.size(); i++) {
        int dif = target - nums[i];
        if (map.find(dif) != map.end()) {
            return {map[dif], i};
        }

        map.insert({nums[i], i});
    }

    return {};
}

std::vector<std::vector<std::string>> Solution::groupAnagrams(std::vector<std::string>& strs) {
    std::unordered_map<std::string, std::vector<std::string>> map;

    for (std::string s : strs) {
        std::vector<int> l(26, 0);
        for (char c : s) {
            l[c - 'a']++;
        }
        std::string key = std::to_string(l[0]);
        for (int i = 0; i < 26; i++) {
            key += ',' + std::to_string(l[i]);
        }
        map[key].push_back(s);
    }
    
    std::vector<std::vector<std::string>> ans;

    for (const auto& pair : map) {
        ans.push_back(pair.second);
    }

    return ans;
}

std::vector<int> Solution::topKFrequent(std::vector<int>& nums, int k) {
    std::unordered_map<int, int> count;
    std::vector<std::vector<int>> freq(nums.size() + 1);
    for (int num : nums) {
        count[num] = 1 + count[num];
    }
    for (const auto& pair : count) {
        freq[pair.second].push_back(pair.first);
    }

    std::vector<int> res;
    for (int i = freq.size() - 1; i > 0; --i) {
        for (int n : freq[i]) {
            res.push_back(n);
            if (res.size() == k) {
                return res;
            }
        }
    }

    return res;
}

std::vector<int> Solution::productExceptSelf(std::vector<int>& nums) {
    std::vector<int> res(nums.size(), 1);

    for(int i = 1; i < nums.size(); i++) {
        res[i] = res[i-1] * nums[i-1];
    }

    int p = 1;
    for (int i = nums.size() - 1; i >= 0; --i) {
        res[i] *= p;
        p *= nums[i];
    }

    return res;
}

bool Solution::isValidSudoku(std::vector<std::vector<char>>& board) {
    std::map<std::pair<int, int>, std::unordered_set<char>> sq;
    std::unordered_map<int, std::unordered_set<char>> rows, cols;
    
    for (int i = 0; i < board.size(); i++) {
        for (int j = 0; j < board[i].size(); j++) {
            if (board[i][j] == '.') {
                continue;
            }

            std::pair<int, int> sqk = {i/3, j/3};

            if (rows[i].count(board[i][j]) || cols[j].count(board[i][j]) || sq[sqk].count(board[i][j])) {
                return false;
            }
            
            rows[i].insert(board[i][j]);
            cols[j].insert(board[i][j]);
            sq[sqk].insert(board[i][j]);
        }
    }
    return true;
}
