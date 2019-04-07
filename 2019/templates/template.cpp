#include <algorithm>
#include <iostream>
#include <string>
#include <vector>
#include <queue>
#include <set>
#include <map>

using namespace std;

string solution(int N, vector<int> ciper){
    return "Hello";
}

int main() {
	int test_cases;
	cin >> test_cases;

	for (int test_case = 1; test_case <= test_cases; test_case++) {
		int N, L;
		cin >> N >> L;
		vector<int> ciper(L, 0);
		for(int i=0;i<L;i++){
		    cin >> ciper[i];
		}
		string answer = solution(N, ciper);
		cout << "Case #" << test_case << ": " << answer << endl;
		cout.flush();
	}
	return 0;
}