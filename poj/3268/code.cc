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
dijstra(int **g, int n, int s, int *d) {
    int *mark = newarray(n);
    int *p = newarray(n);

    mark[s] = 1;
    for (int v = 1; v <= n; v ++) {
        if (mark[v] == 0 && g[s][v] != 0) {
            d[v] = g[s][v];
            p[v] = s;
        }
    }

    while (true) {
        int x = 0;
        for (int v = 1; v <= n; v ++) {
            if (mark[v] == 0 && p[v] != 0) {
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
            if (mark[v] == 0 && g[x][v] != 0) {
                if (p[v] == 0 || d[v] > g[x][v] + d[x]) {
                    d[v] = g[x][v] + d[x];
                    p[v] = x;
                }
            }
        }
    }

    free(p);
    free(mark);
}

void
swap(int *x, int *y) {
    int t = *x; *x = *y; *y = t;
}

int
main(void) {
    int n, m, s;
    scanf("%d %d %d", &n, &m, &s);
    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        g[i] = newarray(n);
    }

    for (int i = 1; i <= m; i ++) {
        int a, b, d;
        scanf("%d %d %d", &a, &b, &d);
        g[a][b] = d;
    }

    int *d1 = newarray(n);
    int *d2 = newarray(n);

    dijstra(g, n, s, d1);

    for (int i = 1; i <= n; i ++) {
        for (int j = i + 1; j <= n; j ++) {
            swap(&g[i][j], &g[j][i]);
        }
    }

    dijstra(g, n, s, d2);

    int max = 0;
    for (int i = 1; i <= n; i ++) {
        int cost = d1[i] + d2[i];
        if (cost > max) {
            max = cost;
        }
    }
    printf("%d\n", max);

    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);

    free(d1);
    free(d2);
    return 0;
}
