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
max(int v1, int v2) {
    if (v1 > v2) {
        return v1;
    } else {
        return v2;
    }
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
    g[i][j] = 0;
    int sum = 1;
    for (int k = 0; k < 4; k ++) {
        sum += dfs(g, m, n, i + dx[k], j + dy[k]);
    }
    return sum;
}

int
main(void) {
    int n, m, k;
    scanf("%d %d %d", &m, &n, &k);
    int **g = (int **)malloc(sizeof(int *) * (m + 1));
    for (int i = 1; i <= m; i ++) {
        g[i] = newarray(n);
    }

    for (int i = 1; i <= k; i ++) {
        int x, y;
        scanf("%d %d", &x, &y);
        g[x][y] = 1;
    }

    int area = 0;
    for (int i = 1; i <= m; i ++) {
        for (int j = 1; j <= n; j ++) {
            area = max(area, dfs(g, m, n, i, j));
        }
    }
    printf("%d\n", area);

    for (int i = 1; i <= m; i ++) {
        free(g[i]);
    }
    free(g);
    return 0;
}
