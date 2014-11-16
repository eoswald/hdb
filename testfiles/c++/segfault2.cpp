#include <iostream>

void increment_pointer(int* &ptr){
  *ptr++;
  std::cout<<"ptr is now "<<*ptr<<"!"<<std::endl;
  ptr+=1;
  ptr=ptr-ptr;
}


int main(){
  int * ptr;
  ptr = new int(5);
  increment_pointer(ptr);
  std::cout<<*ptr<<std::endl;
  return 0;
}
