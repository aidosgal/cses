#include <iostream>

int main(void) {
    int n; std::cin >> n;
    int sum1 = 0;
    int sum2 = 0;
    for(int i = 0; i < n - 1; i++) {
        sum1 = sum1 + i;
        int s; std::cin >> s;
        sum2 = sum2 + s;
    }
    sum1 = sum1 + n;

    std::cout << sum1 - sum2;
}
