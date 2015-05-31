#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *
newarray(int n) {
    int size = sizeof(int) * (n + 1);
    int *p = (int *)malloc(size);
    return (int *)memset(p, 0, size);
}

int dx[4] = {1, 0, -1, 0};
int dy[4] = {0, 1, 0, -1};


int
dfs(int **g, int m, int n, int i, int j) {
    if (i < 1 || i > m) {
        return 0;
    }
    if (j < 1 || j > n) {
        return 0;
    }
    if (g[i][j] == 0) {
        return 0;
    }
    int sum = 1;
    g[i][j] = 0;
    for (int k = 0; k < 4; k ++) {
        sum += dfs(g, m, n, i + dx[k], j + dy[k]);
    }
    return sum;
}

int
main(void) {
    int m, n;
    scanf("%d %d\n", &n, &m);

    int **g = (int **)malloc(sizeof(int *) * (m + 1));
    for (int i = 1; i <= m; i ++) {
        g[i] = newarray(n);
    }
    for (int i = 1; i <= m; i ++) {
        for (int j = 1; j <= n; j ++) {
            char c;
            scanf("%c", &c);
            if (c == '*') {
                g[i][j] = 1;
            }
        }
        scanf("\n");
    }

    int max = 0;
    for (int i = 1; i <= m; i ++) {
        for (int j = 1; j <= n; j ++) {
            int v = dfs(g, m, n, i, j);
            if (v > max) {
                max = v;
            }
        }
    }
    printf("%d\n", max);
    return 0;
}
