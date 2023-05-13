#include <stdio.h>

// SYNOP pettern 1 ~ 5
#define SYNOP1 1
#define SYNOP2 2
#define SYNOP3 3
#define SYNOP4 4
#define SYNOP5 5

// パターン名を返す
int detect_synop_pattern(int m)
{
	if(m <= 99){
		return SYNOP1;
	} else if(100 <= m && m <= 5000){
		return SYNOP2;
	} else if(6000 <= m && m <= 30000){
		return SYNOP3;
	} else if(35000 <= m && m <= 70000){
		return SYNOP4;
	} else if(70001 <= m){
		return SYNOP5;
	}
	// エラー処理
	return 1;
}

int calc_vv(int m, int synop)
{
	switch(synop){
		case SYNOP1:
			return 0;
		case SYNOP2:
			return m / 100;
		case SYNOP3:
			return m / 1000 + 50;
		case SYNOP4:
			return ((m / 1000) - 30) / 5 + 80;
		case SYNOP5:
			return 89;
	}
	// エラー処理
	return 1;
}


int	main(void)
{
	int m;
	scanf("%d",&m);

	int pattern = detect_synop_pattern(m);
	// printf("%d\n", synop);
	int vv = calc_vv(m, pattern);
	printf("%02d\n", vv);
	
	return (0);
}
