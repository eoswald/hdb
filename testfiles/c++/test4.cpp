#include<vector>
#include<iostream>

int findprime(int n){
	std::vector<bool> v(n,1);
	int count = 0;
	for (int x=1;x<n;x++){
		if (v[x]){
			for (int y=2*x+1;y<n;y+=x+1){
				v[y]=0;
			}
			count+=1;
		}
	}
	count-=2;
	return count;
}
int main(){
	std::cout<<findprime(1000)<<std::endl;
	return 0;
}