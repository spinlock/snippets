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
    int n, a, b;
    scanf("%d %d %d", &n, &a, &b);

    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        g[i] = newarray(n);
    }
    for (int i = 1; i <= n; i ++) {
        for (int j = 1; j <= n; j ++) {
            g[i][j] = -1;
        }
    }

    for (int i = 1; i <= n; i ++) {
        int m, v;
        scanf("%d", &m);
        for (int j = 1; j <= m; j ++) {
            scanf("%d", &v);
            if (j == 1) {
                g[i][v] = 0;
            } else {
                g[i][v] = 1;
            }
        }
    }

    int *mark = newarray(n);
    int *d = newarray(n);
    int *p = newarray(n);

    mark[a] = 1;
    p[a] = a;
    for (int v = 1; v <= n; v ++) {
        if (mark[v] == 0 && g[a][v] >= 0) {
            d[v] = g[a][v];
            p[v] = a;
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
            if (mark[v] == 0 && g[x][v] >= 0) {
                if (p[v] == 0 || d[v] > d[x] + g[x][v]) {
                    d[v] = d[x] + g[x][v];
                    p[v] = x;
                }
            }
        }
    }

    if (p[b] != 0) {
        printf("%d\n", d[b]);
    } else {
        printf("%d\n", -1);
    }

    free(mark);
    free(d);
    free(p);

    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);
    return 0;
}
