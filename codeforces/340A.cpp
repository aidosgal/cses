#include <bits/stdc++.h>

int main(void) {
    int n; std::cin >> n;

    int ans = 0;

    while (n > 0) {
        if (n >= 5) {
            n = n % 5;
        } else if (n >= 4) {
            n = n % 4;
        } else if (n >= 3) {
            n = n % 3;
        } else if (n >= 2) {
            n = n % 2;
        } else if (n >= 1) {
            n = n % 1;
        }

        ans++;
    }

    std::cout << ans;
}
