#include <algorithm>
#include <iostream>
#include <string>

using namespace std;

string solution(string enemypath){
	string path = "";
	for(int i=0;i<enemypath.size();i++){
		if(enemypath[i] == 'E'){
			path += "S";
		} else {
			path += "E";
		}
	}
	return path;
}

int main() {

	int test_cases;
	cin >> test_cases;

	for (int test_case = 1; test_case <= test_cases; test_case++) {
		string N, enemypath;
		cin >> N >> enemypath;
		string answer = solution(enemypath);
		cout << "Case #" << test_case << ": " << answer << endl;
		cout.flush();
	}
	return 0;
}