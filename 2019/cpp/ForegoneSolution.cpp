#include <algorithm>
#include <iostream>
#include <string>
#include <vector>
#include <queue>
#include <set>
#include <map>

using namespace std;

string trimZero(string s){
	int i;
	for(i=0;i<s.size() && s[i] == '0';i++){
	}
	return s.substr(i);
}

vector<string> solution(string N){
	string right = "";
	string left = "";
	for(int i=0;i<N.size();i++){
		if(N[i] == '4'){
			left += "2";
			right += "2";
		} else {
			left += N[i];
			right += "0";
		}
	}
	vector<string>out = {trimZero(left), trimZero(right)};
	return out;
}

int main() {

	int test_cases;
	cin >> test_cases;

	for (int test_case = 1; test_case <= test_cases; test_case++) {
		string N;
		cin >> N;
		vector<string> answer = solution(N);
		cout << "Case #" << test_case << ": " << answer[0] << " " << answer[1] << endl;
		cout.flush();
	}
	return 0;
}