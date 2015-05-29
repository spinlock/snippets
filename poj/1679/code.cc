#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *
newarray(int n) {
    int size = sizeof(int) * (n + 1);
    int *p = (int *)malloc(size);
    return (int *)memset(p, 0, size);
}

void
process(int n, int m) {
    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        g[i] = newarray(n);
    }
    for (int i = 1; i <= m; i ++) {
        int u, v, d;
        scanf("%d %d %d", &u, &v, &d);
        g[u][v] = g[v][u] = d;
    }

    int *mark = newarray(n);
    int *d = newarray(n);
    int *p = newarray(n);

    mark[1] = 1;
    for (int v = 1; v <= n; v ++) {
        if (!mark[v] && g[1][v] != 0) {
            d[v] = g[1][v];
            p[v] = 1;
        }
    }

    bool flag = true;
    while (flag) {
        int x = 0;
        for (int v = 1; v <= n; v ++) {
            if (!mark[v] && p[v] != 0) {
                if (x == 0 || d[x] > d[v]) {
                    x = v;
                }
            }
        }
        if (x == 0) {
            break;
        }
        for (int v = 1; v <= n; v ++) {
            if (mark[v] && g[x][v] != 0 && g[x][v] <= d[x]) {
                if (p[x] != v) {
                    flag = false;
                }
            }
        }
        mark[x] = 1;
        for (int v = 1; v <= n; v ++) {
            if (!mark[v] && g[x][v] != 0) {
                if (p[v] == 0 || d[v] > g[x][v]) {
                    d[v] = g[x][v];
                    p[v] = x;
                }
            }
        }
    }

    int sum = 0;
    for (int v = 1; v <= n; v ++) {
        sum += d[v];
    }
    if (flag) {
        printf("%d\n", sum);
    } else {
        printf("Not Unique!\n");
    }

    free(d);
    free(p);
    free(mark);

    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);
}

int
main(void) {
    int t, n, m;
    scanf("%d\n", &t);
    for (int i = 0; i < t; i ++) {
        scanf("%d %d", &n, &m);
        process(n, m);
    }
    return 0;
}
