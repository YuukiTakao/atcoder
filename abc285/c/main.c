#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define Mod 1000000007
typedef	long long ll;
ll mod_pow(int a, int exp) {
	if (exp == 0)
		return 1;
	ll p = a, ans = 1;
	for (int i = 0; i < 30; i++) {
		int div = 1 << i;
		if ((exp / div) % 2 == 1)
			ans = (ans * p);
		p = (p * p);
	}
	return ans;
}
int	main(void) {
	char S[30];
	scanf("%s", S);
	// printf("%s\n", S);
	int sl = strlen(S);
	// printf("%d\n", sl);

	ll offset = 64;
	ll base = 26;
	ll ans = 0;
	for (int i = sl -1; i >= 0; i--) {
		ans += mod_pow(base, sl-1-i) * (S[i] - offset);
		// printf("i=%d sl-1-i=%d ((int)(S[i]-'A')+1=%lld mod_pow(base, sl-1-i)=%lld ans=%lld\n", i, sl-1-i, (S[i] - offset), mod_pow(base, sl-1-i), ans);
	}
	printf("%lld\n", ans);
	return (0);
}