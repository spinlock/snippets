#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int u, v, l;
} edge_t;

void
xswap(int *x, int *y) {
    int t = *x; *x = *y; *y = t;
}

void
eswap(edge_t *edges, int i, int j) {
    if (i != j) {
        xswap(&edges[i].u, &edges[j].u);
        xswap(&edges[i].v, &edges[j].v);
        xswap(&edges[i].l, &edges[j].l);
    }
}

void
esort(edge_t *edges, int beg, int end) {
    if (beg >= end) {
        return;
    }
    int pivot = beg;
    for (int i = beg + 1; i <= end; i ++) {
        if (edges[i].l < edges[beg].l) {
            pivot ++;
            eswap(edges, i, pivot);
        }
    }
    eswap(edges, beg, pivot);
    esort(edges, beg, pivot - 1);
    esort(edges, pivot + 1, end);
}

int *
ufs_init(int n) {
    int *ufs = (int *)malloc(sizeof(int) * (n + 1));
    for (int i = 1; i <= n; i ++) {
        ufs[i] = -1;
    }
    return ufs;
}

void
ufs_free(int *ufs) {
    free(ufs);
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
    if (px == py) {
        return false;
    } else {
        ufs[px] += ufs[py];
        ufs[py] = px;
        return true;
    }
}

int
main(void) {
    int n, m;
    scanf("%d %d", &n, &m);

    edge_t *edges = (edge_t *)malloc(sizeof(edge_t) * m);
    for (int i = 0; i < m; i ++) {
        scanf("%d %d %d", &edges[i].u, &edges[i].v, &edges[i].l);
    }
    esort(edges, 0, m - 1);

    int max = 0;
    int *ufs = ufs_init(n);
    for (int i = 0; i < m; i ++) {
        if (ufs_join(ufs, edges[i].u, edges[i].v)) {
            if (edges[i].l > max) {
                max = edges[i].l;
            }
        }
    }
    printf("%d\n", max);

    ufs_free(ufs);
    free(edges);
    return 0;
}
