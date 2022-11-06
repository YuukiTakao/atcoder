#include <stdio.h>

int	main(void)
{
	char	S[101];
	int 	a_position;


	scanf("%s", S);
	a_position = -1;
	for (int i = 0; S[i] != '\0'; i++)
	{
		if (S[i] == 'a')
			a_position = i + 1;
	}
	printf("%d", a_position);
	return (0);
}