#include <algorithm>
#include <iostream>
#include <string>
#include <vector>
#include <queue>
#include <set>
#include <map>

using namespace std;


vector<bool> sieveEratosthenes(int max) { //Helper for getPrimes
	//initialize
	vector<bool> flags(max + 1);
	flags[0] = false, flags[1] = false;
	for (int i = 2; i < flags.size(); i++)	flags[i] = true;

	int prime = 2;
	while (prime < sqrt(max)) {
		for (int i = prime*prime; i < flags.size(); i += prime)	flags[i] = false;

		int next = prime + 1;
		while (next < flags.size() && !flags[next]) next++;
		prime =  next;
	}
	return flags;
}

vector<int> getPrimes(int n){ //Finds all prime numbers up to max
	vector<bool> primeTable;
}


string solution(int N, vector<int> ciper){
	vector<int>
    return "Hello";
}

int main() {
	int test_cases;
	cin >> test_cases;

	for (int test_case = 1; test_case <= test_cases; test_case++) {
		int N, L;
		cin >> N >> L;
		vector<int> cipher(L, 0);
		for(int i=0;i<L;i++){
		    cin >> cipher[i];
		}
		string answer = solution(N, cipher);
		cout << "Case #" << test_case << ": " << answer << endl;
		cout.flush();
	}
	return 0;
}