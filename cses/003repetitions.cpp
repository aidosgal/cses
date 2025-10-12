#include <bits/stdc++.h>

using namespace std;

int main(void) {
    string n;
    cin >> n;
    int ans = 0;
    int c = 0;
    for (int i = 0; i < (int) n.size(); i++) {
      if (i != 0 && n[i] == n[i-1]) {
        c++;
      } else {
        ans = max(ans, c);
        c = 1;
      }
    }
    ans = max(ans, c);
    cout << ans;
    return 0;
}
