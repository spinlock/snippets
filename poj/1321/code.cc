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
dfs(int **g, int *col, int n, int k, int i) {
    if (k == 0) {
        return 1;
    } else if (i > n) {
        return 0;
    }
    int cnt = 0;
    for (int j = 1; j <= n; j ++) {
        if (col[j] == 0 && g[i][j] != 0) {
            col[j] = 1;
            cnt += dfs(g, col, n, k - 1, i + 1);
            col[j] = 0;
        }
    }
    cnt += dfs(g, col, n, k, i + 1);
    return cnt;
}

int
process(int n, int k) {
    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        g[i] = newarray(n);
    }

    for (int i = 1; i <= n; i ++) {
        for (int j = 1; j <= n; j ++) {
            char b;
            scanf("%c", &b);
            if (b == '#') {
                g[i][j] = 1;
            }
        }
        scanf("\n");
    }

    int *col = newarray(n);

    int cnt = dfs(g, col, n, k, 1);

    free(col);
    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);
    return cnt;
}

int
main(void) {
    int n, k;
    while (true) {
        scanf("%d %d\n", &n, &k);
        if (n < 0 || k < 0) {
            return 0;
        }
        printf("%d\n", process(n, k));
    }
    return 0;
}
