#include <iostream>
/*
this is an example segfault file that displays the dopeness of HDB(Hillbilly DeBugger)
have fun!
*/

void increment_pointer(int* &ptr){
  //incremets and prints pointer value and increments the pointer
  *ptr++;
  std::cout<<"ptr is now "<<*ptr<<"!!!"<<std::endl;
  ptr+=1;
  ptr=ptr-ptr;
}
int abc;

int main(){
  int * ptr;
  ptr = new int(5);
  increment_pointer(ptr);
  std::cout<<*ptr<<std::endl;
  return 0;
}
