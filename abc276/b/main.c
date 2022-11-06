#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

// asc
int int_compare(const void * n1, const void * n2)
{
	if (*(int *)n1 > *(int *)n2)
		return 1;
	else if (*(int *)n1 < *(int *)n2)
		return -1;
	else
		return 0;
}

typedef struct {
	int vertex;
	int	*arr;
	int	len;
	int	cap;
}	ad_list_t;

ad_list_t *arr_resize(ad_list_t *self, int cap)
{
    int byte = sizeof(int);
    int size = cap * byte;
    int *tmp = realloc(self->arr, size);
    if (!tmp)
		return NULL;
    self->arr = tmp;
    self->cap = cap;
    return self;
}

ad_list_t *push_back(ad_list_t *self, int elem)
{
	if (self->len >= self->cap)
	{
		if (! arr_resize(self, self->cap * 2))
			return (NULL);
	}
	self->arr[self->len] = elem;
	self->len++;
	return (self);
}

ad_list_t *array_new(void)
{
    ad_list_t *self = calloc(1, sizeof(*self));
    if (!self)
		return NULL;

    self->cap = 100;
    self->arr = calloc(self->cap, sizeof(int));
    if (!self->arr) {
        free(self);
        return NULL;
    }
    return self;
}

int	main(void)
{
	int			N;
	int			M;
	ad_list_t	**ad_list;

	int	res = scanf("%d%d", &N, &M);
	if (res == -1) 
		return (-1);

	ad_list = (ad_list_t **)calloc(N + 2, sizeof(ad_list_t));
	if(ad_list == NULL)
		return -1;

	for (int i = 0; i < M; i++)
	{
		int a, b;
		int	res = scanf("%d%d", &a, &b);
		if (res == -1) 
			return (-1);
		if (ad_list[a] == NULL)
			ad_list[a] = array_new();

		if (ad_list[b] == NULL)
			ad_list[b] = array_new();
		
		ad_list[a]->vertex = a;
		push_back(ad_list[a], b);

		ad_list[b]->vertex = b;
		push_back(ad_list[b], a);
	}

	for (int i = 1; i <= N; i++)
	{
		if (ad_list[i] != NULL)
		{
			qsort(ad_list[i]->arr, ad_list[i]->len, sizeof(int), int_compare);
			printf("%d", ad_list[i]->len);
			for (int j = 0; j < ad_list[i]->len; j++)
				printf(" %d", ad_list[i]->arr[j]);
		}
		else 
			printf("%d", 0);
		printf("\n");
	}
	return (0);
}
