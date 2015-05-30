#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *
newarray(int n) {
    int size = sizeof(int) * n;
    int *p = (int *)malloc(size);
    return (int *)memset(p, 0, size);
}

typedef struct {
    int u, v;
    int d;
} edge_t;

void
swap(int *x, int *y) {
    int t = *x; *x = *y; *y = t;
}

void
ufs_init(int *ufs, int n) {
    for (int i = 1; i <= n; i ++) {
        ufs[i] = -1;
    }
}

int
ufs_find(int *ufs, int x) {
    int px = ufs[x];
    if (px < 0) {
        return x;
    } else {
        ufs[x] = ufs_find(ufs, px);
        return ufs[x];
    }
}

bool
ufs_join(int *ufs, int x, int y) {
    int px = ufs_find(ufs, x);
    int py = ufs_find(ufs, y);
    if (px != py) {
        ufs[px] += ufs[py];
        ufs[py] = px;
        return true;
    } else {
        return false;
    }
}

void
eswap(edge_t *edges, int i, int j) {
    if (i != j) {
        swap(&edges[i].u, &edges[j].u);
        swap(&edges[i].v, &edges[j].v);
        swap(&edges[i].d, &edges[j].d);
    }
}

void
esort(edge_t *edges, int beg, int end) {
    if (beg >= end) {
        return;
    }
    int pivot = beg;
    for (int i = beg + 1; i <= end; i ++) {
        if (edges[i].d < edges[beg].d) {
            pivot ++;
            eswap(edges, i, pivot);
        }
    }
    eswap(edges, beg, pivot);
    esort(edges, beg, pivot - 1);
    esort(edges, pivot + 1, end);
}

int
main(void) {
    int n, m;
    scanf("%d", &n);

    int *x = newarray(n);
    int *y = newarray(n);
    for (int i = 1; i <= n; i ++) {
        scanf("%d %d", &x[i], &y[i]);
    }

    edge_t *edges = (edge_t *)malloc(sizeof(edge_t) * n * n);
    int k = 0;
    for (int u = 1; u <= n; u ++) {
        for (int v = u + 1; v <= n; v ++) {
            int dx = x[u] - x[v];
            int dy = y[u] - y[v];
            edges[k].u = u;
            edges[k].v = v;
            edges[k].d = dx * dx + dy * dy;
            k ++;
        }
    }
    esort(edges, 0, k - 1);

    int *ufs = newarray(n);
    ufs_init(ufs, n);

    scanf("%d", &m);
    for (int i = 1; i <= m; i ++) {
        int u, v;
        scanf("%d %d", &u, &v);
        ufs_join(ufs, u, v);
    }

    for (int i = 0; i < k; i ++) {
        if (ufs_join(ufs, edges[i].u, edges[i].v)) {
            printf("%d %d\n", edges[i].u, edges[i].v);
        }
    }

    free(x);
    free(y);
    free(ufs);
    free(edges);
    return 0;
}
