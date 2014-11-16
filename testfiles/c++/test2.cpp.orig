#include <iostream>
#include <math.h>

void find_missing(int array[], int n){
  //takes O(2^n)time and takes the  memory of two of the arrays
  int size = pow(2,n);
  bool found[size];
  for (int x=0;x<size;x++)found[x]=0;
  for (int x=0;x<size-2;x++)found[array[x]]=1;
  for (int x=0;x<size;x++) if(!found[x])std::cout<<x<<std::endl;
}

int main(){
  int v[]={0,1,3,4,5,6};
  find_missing( v,3);
  return 0;
}
