#include <iostream>
std::cout << "#include <iostream>" std::endl;

std::cout << "" std::endl;

std::cout << "" std::endl;
bool find_matches(int array[], int n){
std::cout << "bool find_matches(int array[], int n){" std::endl;
  //O(n^2)
std::cout << "  //O(n^2)" std::endl;
  for (int x=0;x<n-1;x++){
std::cout << "  for (int x=0;x<n-1;x++){" std::endl;
    for (int y=x+1;y<n;y++){
std::cout << "    for (int y=x+1;y<n;y++){" std::endl;
      if (array[x]==-array[y]){
std::cout << "      if (array[x]==-array[y]){" std::endl;
        std::cout<<array[x]<<array[y]<<std::endl;
std::cout << "        std::cout<<array[x]<<array[y]<<std::endl;" std::endl;
        return 1;
std::cout << "        return 1;" std::endl;
      }
std::cout << "      }" std::endl;
    }
std::cout << "    }" std::endl;
  }
std::cout << "  }" std::endl;
  return 0;
std::cout << "  return 0;" std::endl;
}
std::cout << "}" std::endl;

std::cout << "" std::endl;
int main(){
std::cout << "int main(){" std::endl;
  int v[]={5,3,65,7,4,-2,-2,4,5,-7,3,-2};
std::cout << "  int v[]={5,3,65,7,4,-2,-2,4,5,-7,3,-2};" std::endl;
  std::cout<<find_matches(v,12)<<std::endl;
std::cout << "  std::cout<<find_matches(v,12)<<std::endl;" std::endl;
  return 0;
std::cout << "  return 0;" std::endl;
}
std::cout << "}" std::endl;

std::cout << "" std::endl;
