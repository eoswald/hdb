#include<vector>
#include<iostream>
#include<cmath>

int lcm(int x, int y){
	float lower, higher;
	if (x<y){
		lower=x;
		higher=y;
	}
	else{
		higher=x;
		lower=y;
	}
	std::cout<<higher<<"    "<<lower<<std::endl;

	for (int a=ceil(higher/lower)*lower;a<higher*lower;a+=lower){
		if (floor(a/higher)*higher==a)return a;
	}
}
int main(){

	std::cout<<lcm(15,30)<<std::endl;
	return 0;
}