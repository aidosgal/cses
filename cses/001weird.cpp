#include <iostream>

int main(void) {
    long long n; std::cin >> n;
    while(n != 1) {
        if (n % 2 == 0) {
            n = n / 2;
        } else {
            n = (n * 3) + 1;
        }
    }
    std::cout << n;
}
