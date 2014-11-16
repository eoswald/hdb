#include<iostream>
#include<cmath>
int find_sqrt(unsigned int in){
  unsigned int low(1), high(in), guess(1);
  while (high-low>1){
    guess= floor((high-low)/2)+low;
    if (guess*guess>in)high=guess;
    else if(guess*guess<in)low=guess;
    else return guess;
  }
  return high;
}

int main(){
  unsigned int input;
  std::cin>>input;
  while(input!=0){
    std::cout<<find_sqrt(input)<<"\n"<<std::endl;
    std::cin>>input;
  }
  return 0;
}
