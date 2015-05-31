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
main(void) {
    int n, m;
    scanf("%d", &n);
    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        g[i] = newarray(n);
    }
    for (int i = 1; i <= n; i ++) {
        for (int j = 1; j <= n; j ++) {
            scanf("%d", &g[i][j]);
        }
    }
    scanf("%d", &m);

    for (int i = 1; i <= m; i ++) {
        int u, v;
        scanf("%d %d", &u, &v);
        g[u][v] = g[v][u] = 0;
    }

    int *mark = newarray(n);
    int *d = newarray(n);
    int *p = newarray(n);
    mark[1] = 1;
    for (int v = 1; v <= n; v ++) {
        d[v] = g[1][v];
        p[v] = 1;
    }

    while (true) {
        int x = 0;
        for (int v = 1; v <= n; v ++) {
            if (mark[v] == 0) {
                if (x == 0 || d[x] > d[v]) {
                    x = v;
                }
            }
        }
        if (x == 0) {
            break;
        }
        mark[x] = 1;
        for (int v = 1; v <= n; v ++) {
            if (mark[v] == 0 && d[v] > g[x][v]) {
                d[v] = g[x][v];
                p[v] = x;
            }
        }
    }
    int sum = 0;
    for (int v = 2; v <= n; v ++) {
        sum += d[v];
    }
    printf("%d\n", sum);

    free(mark);
    free(d);
    free(p);
    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);
    return 0;
}
