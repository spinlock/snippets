#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *
newarray(int n) {
    int size = sizeof(int) * (n + 1);
    int *p = (int *)malloc(size);
    return (int *)memset(p, 0, size);
}

int
log2(int w) {
    int p = 0;
    for (int x = 1; x < w; x *= 2) {
        p ++;
    }
    return p;
}

int
process(int m, int w) {
    int *values1 = newarray(m);
    int *weight1 = newarray(m);
    for (int i = 1; i <= m; i ++) {
        scanf("%d %d", &values1[i], &weight1[i]);
    }

    int p = log2(w) + 1;
    int *values2 = newarray(m * p);
    int *weight2 = newarray(m * p);
    int n = 0;
    for (int i = 1; i <= m; i ++) {
        int vx = values1[i];
        int wx = weight1[i];
        while (wx <= w) {
            n ++;
            values2[n] = vx;
            weight2[n] = wx;
            vx *= 2;
            wx *= 2;
        }
    }

    int *so = newarray(w);
    int *sn = newarray(w);
    for (int j = 1; j <= w; j ++) {
        so[j] = -1;
    }
    for (int i = 1; i <= n; i ++) {
        int vx = values2[i];
        int wx = weight2[i];
        for (int j = 1; j <= w; j ++) {
            sn[j] = so[j];
            if (j >= wx && so[j - wx] >= 0) {
                int t = so[j - wx] + vx;
                if (sn[j] < 0 || sn[j] > t) {
                    sn[j] = t;
                }
            }
        }
        int *sp = so; so = sn; sn = sp;
    }
    int min = so[w];

    free(so);
    free(sn);
    free(values1);
    free(values2);
    free(weight1);
    free(weight2);
    return min;
}

int
main(void) {
    int t;
    scanf("%d", &t);
    for (int i = 0; i < t; i ++) {
        int e, f, n;
        scanf("%d %d", &e, &f);
        scanf("%d", &n);
        int min = process(n, f - e);
        if (min >= 0) {
            printf("The minimum amount of money in the piggy-bank is %d.\n", min);
        } else {
            printf("This is impossible.\n");
        }
    }
    return 0;
}
