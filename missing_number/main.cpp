#include <iostream>

int main() {
    int n;
    std::cin >> n;
    int sum = n;
    int s2 = 0;
    for(int i = 1; i < n; ++i) {
       sum += i;
       int n2;
       std::cin >> n2;
       s2 += n2;
    }

    std::cout << sum - s2; 
}
