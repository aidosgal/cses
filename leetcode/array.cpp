#include "array.hpp"

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
