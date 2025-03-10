#include <iostream>

int main(void) {
    std::string n;
    std::cin >> n;
    
    int ans = 0;
    int c = 1;
    
    for (int i = 0; i < n.length() - 1; i++){
        if (n[i] == n[i+1]) {
            c++;
        } else {
            c = 1;
        }
        if (c > ans) {
            ans = c;
        }
    }

    std::cout << ans;
    return 0;
}
