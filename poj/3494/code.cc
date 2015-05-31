#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    bool one;
    int beg, end;
    int height;
} node_t;

int
min(int v1, int v2) {
    if (v1 < v2) {
        return v1;
    } else {
        return v2;
    }
}

int
max(int v1, int v2) {
    if (v1 > v2) {
        return v1;
    } else {
        return v2;
    }
}

void
process(int n, int m) {
    node_t **g = (node_t **)malloc(sizeof(node_t *) * (m + 1));
    for (int i = 0; i <= m; i ++) {
        int size = sizeof(node_t) * (n + 1);
        node_t *p = (node_t *)malloc(size);
        memset(p, 0, size);
        g[i] = p;
    }

    for (int i = 1; i <= m; i ++) {
        int v;
        for (int j = 1; j <= n; j ++) {
            scanf("%d", &v);
            g[i][j].one = v != 0;
        }
    }

    for (int j = 1; j <= n; j ++) {
        g[0][j].beg = 1;
        g[0][j].end = n;
    }

    for (int i = 1; i <= m; i ++) {
        int beg = 1, end = n;
        for (int j = 1; j <= n; j ++) {
            if (g[i][j].one) {
                g[i][j].height = g[i - 1][j].height + 1;
                g[i][j].beg = max(beg, g[i - 1][j].beg);
            } else {
                g[i][j].beg = 1;
                beg = j + 1;
            }
        }
        for (int j = n; j >= 1; j --) {
            if (g[i][j].one) {
                g[i][j].end = min(end, g[i - 1][j].end);
            } else {
                g[i][j].end = n;
                end = j - 1;
            }
        }
    }

    int area = 0;
    for (int i = 1; i <= m; i ++) {
        for (int j = 1; j <= n; j ++) {
            if (g[i][j].one) {
                int beg = g[i][j].beg;
                int end = g[i][j].end;
                if (beg <= end) {
                    area = max(area, g[i][j].height * (end - beg + 1));
                }
            }
        }
    }
    printf("%d\n", area);

    for (int i = 0; i <= m; i ++) {
        free(g[i]);
    }
    free(g);
}

int
main(void) {
    while (true) {
        int n, m;
        if (scanf("%d %d", &m, &n) == EOF) {
            return 0;
        }
        process(n, m);
    }
    return 0;
}
