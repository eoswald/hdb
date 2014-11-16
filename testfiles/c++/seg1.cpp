#include <iostream>
#include <string>
using namespace std;

struct people
{
	string forename;
	string lastname;
	int age;
};

int main()
{
	int num_numbers = 5; //ask the user how many numbers the sequence will contain
	string first_name[5] = {"Josh", "Derek", "Eric", "David", "Luke"};
	string last_name[5] = {"Makinen", "Meer", "Oswald", "Vorick", "Champine"};
	int age[5] = {1, 2, 3, 4, 5};
	cout << "Enter how people will be entered : \n";

	people * peoples;
	peoples = new people [num_numbers];

	for(int x = 0; x < num_numbers; x++) 
	{
		 cout<<"Enter forename "<< x <<":\n";
		 peoples[x].forename = first_name[x];
		 cout<<"Enter surname "<< x <<":\n";
		 peoples[x].lastname = last_name[x];
		 cout<<"Enter age "<< x <<":\n";
		 peoples[x].age = age[x];
	}

	for(int i = 0; i<= num_numbers; i++)
	{
			cout << peoples[i].forename;
			cout << peoples[i].lastname;
			cout << peoples[i].age;
	}

//delete[] peoples;
}
