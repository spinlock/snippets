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
process(int n) {
    int **g = (int **)malloc(sizeof(int *) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        g[i] = newarray(n);
    }

    for (int i = 1; i < n; i ++) {
        char f, t;
        int k, d;
        scanf("%c %d", &f, &k);
        f = f - 'A' + 1;
        for (int j = 0; j < k; j ++) {
            scanf(" %c %d", &t, &d);
            t = t - 'A' + 1;
            g[f][t] = g[t][f] = d;
        }
        scanf("\n");
    }

    int *d = newarray(n);
    int *p = newarray(n);
    int *mark = newarray(n);

    for (int v = 2; v <= n; v ++) {
        if (mark[v] == 0 && g[1][v] != 0) {
            d[v] = g[1][v];
            p[v] = 1;
        }
    }
    while (true) {
        int x = 0;
        for (int v = 2; v <= n; v ++) {
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
        for (int v = 2; v <= n; v ++) {
            if (mark[v] == 0 && g[x][v] != 0) {
                if (p[v] == 0 || d[v] > g[x][v]) {
                    d[v] = g[x][v];
                    p[v] = x;
                }
            }
        }
    }

    int sum = 0;
    for (int v = 2; v <= n; v ++) {
        sum += d[v];
    }

    free(d);
    free(p);
    free(mark);
    for (int i = 1; i <= n; i ++) {
        free(g[i]);
    }
    free(g);
    return sum;
}

int
main(void) {
    int n;
    while (true) {
        scanf("%d\n", &n);
        if (n == 0) {
            return 0;
        }
        printf("%d\n", process(n));
    }
    return 0;
}
