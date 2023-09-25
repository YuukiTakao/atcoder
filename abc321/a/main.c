#include <stdio.h>
#include <stdlib.h>

int	main(void) {
	int n;
	scanf("%d", &n);
	int digits[5];

	int keta = 0;
	for (int i=0; n > 0; i++) {
		digits[i] = n % 10;
		keta++;
		n /= 10;
	}
	for (int i = 0; i+1<keta; i++) {
		if (digits[i] >= digits[i+1]) {
			printf("No\n");
			return 0;
		}
	}
	printf("Yes\n");
	return (0);
}