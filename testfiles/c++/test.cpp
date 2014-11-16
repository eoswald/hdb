#include <iostream>


bool find_matches(int array[], int n){
  //O(n^2)
  for (int x=0;x<n-1;x++){
    for (int y=x+1;y<n;y++){
      if (array[x]==-array[y]){
        std::cout<<array[x]<<array[y]<<std::endl;
        return 1;
      } else if (x%2==0) {
				x = x;
			} else {
				y = y;
			}
    }
  }
  return 0;
}

int main(){
  int v[]={5,3,65,7,4,-2,-2,4,5,-7,3,-2};
  std::cout<<find_matches(v,12)<<std::endl;
  return 0;
}
