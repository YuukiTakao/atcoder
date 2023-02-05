#include <stdio.h>
#include <stdlib.h>

typedef struct {
	int *parent;
	int *size;
}	UF;

// Nは頂点数
UF *init_uf(int N) {
	UF *uf = calloc(1, sizeof(UF));
	uf->parent = calloc(N+1, sizeof(int));
	uf->size = calloc(N+1, sizeof(int));
	for (int i = 1; i <= N; i++) {
		uf->parent[i] = -1;
		uf->size[i] = 1;
	}
	return uf;
}
int	root(UF *uf, int x) {
	while(1) {
		if (uf->parent[x] == -1)
			break;
		x = uf->parent[x];
	}
	return x;
}
void	unite(UF *uf, int u, int v) {
	int RootU = root(uf, u);
	int RootV = root(uf, v);
	if (RootU == RootV)
		return ;
	if (uf->size[RootU] < uf->size[RootV]) {
		uf->parent[RootU] = RootV;
		uf->size[RootV] = uf->size[RootU] + uf->size[RootV];
	} else {
		uf->parent[RootV] = RootU;
		uf->size[RootU] = uf->size[RootV] + uf->size[RootU];
	}
}
int	uf_same(UF *uf, int u, int v) {
	return root(uf, u) == root(uf, v);
}

int	main(void)
{
	int n, m;

	scanf("%d %d", &n, &m);
	// printf("%d %d\n", n, m);

	UF *uf = init_uf(n);
	// printf("size=%d parent=%d\n", uf->size[1], uf->parent[1]);

	int ans = 0;
	for (int i = 1; i <= m; i++) {
		int a, b;
		scanf("%d %d", &a, &b);
		// printf("%d %d\n", a,b);
		
		if (uf_same(uf, a, b)){
			ans++;
		}
		unite(uf, a, b);
	}
	printf("%d\n", ans);
	return (0);
}