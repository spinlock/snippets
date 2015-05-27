#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int
max(int v1, int v2) {
    if (v1 > v2) {
        return v1;
    } else {
        return v2;
    }
}

int
main(void) {
    int n;
    scanf("%d", &n);
    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        int size = sizeof(int) * (n + 1);
        g[i] = (int *)malloc(size);
        memset(g[i], 0, size);
    }

    for (int i = 1; i <= n; i ++) {
        for (int j = 1; j <= i; j ++) {
            scanf("%d", &g[i][j]);
        }
    }
    for (int i = n - 1; i >= 1; i --) {
        for (int j = 1; j <= i; j ++) {
            g[i][j] += max(g[i + 1][j], g[i + 1][j + 1]);
        }
    }
    printf("%d\n", g[1][1]);

    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);
    return 0;
}
