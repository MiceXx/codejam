#include <iostream>
#include <vector>
#include <fstream>
using namespace std;

int main() {
	int i, T, N, result, d;
	int count = 0;
	vector<int> seenDigits;
	ifstream infile("A-large-practice.in");
	ofstream outfile("output.txt");

	infile >> T;

	for (int t = 0; t < T; t++) {	//example number
		infile >> N;
		count = 0;
		seenDigits = {};
		for (i = 1; i < 101; i++) {  //counting sheep
			result = i*N;
			while (result >= 10) { //iterate through each digit
				d = result % 10; //last digit
				if (find(seenDigits.begin(), seenDigits.end(), d) != seenDigits.end()) {}
				else {
					seenDigits.push_back(d);
					count++;
				}
				result = result / 10;
			}
			if (find(seenDigits.begin(), seenDigits.end(), result) != seenDigits.end()) {}
			else {
				seenDigits.push_back(result);
				count++;
			}
			if (count == 10) break;
		}
		outfile << "Case #" << t + 1 << ": ";
		if (count == 10) outfile << i*N << "\n";
		else outfile << "INSOMNIA" << "\n";
	}
	return 0;
}