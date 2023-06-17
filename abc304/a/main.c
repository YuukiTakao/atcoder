#include <stdio.h>
#include <stdlib.h>

typedef struct {
	char name[11];
	int age;
} Person;

int	main(void) {
	int N;
	scanf("%d", &N);
	Person **person = calloc(N, sizeof(Person *));  // Fixed here

	int min = 1000000000;
	int minI = -1;
	for (int i = 0; i < N; i++) {
		person[i] = calloc(1, sizeof(Person));

		scanf("%s", person[i]->name);
		scanf("%d", &(person[i]->age));
		// printf("%s ", person[i]->name);
		// printf("%d\n", person[i]->age);
		if (person[i]->age < min) {
			min = person[i]->age;
			minI = i;
		}
	}
	// printf("min=%d\n", min);
	// printf("minI=%d\n", minI);
	for (int i = 0; i < N; i++)
	{
		printf("%s\n", person[(i+minI) % N]->name);
	}
	
	return (0);
}