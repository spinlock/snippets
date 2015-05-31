#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int u, v;
    int cost;
} edge_t;

void
eswap(edge_t **edges, int i, int j) {
    if (i != j) {
        edge_t *t = edges[i];
        edges[i] = edges[j];
        edges[j] = t;
    }
}

bool
ecomp(edge_t **edges, int i, int j) {
    return edges[i]->cost > edges[j]->cost;
}

void
esort(edge_t **edges, int beg, int end) {
    if (beg >= end) {
        return;
    }
    int pivot = beg;
    for (int i = beg + 1; i <= end; i ++) {
        if (ecomp(edges, i, beg)) {
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

    edge_t *epool = (edge_t *)malloc(sizeof(edge_t) * m);
    for (int i = 0; i < m; i ++) {
        edge_t *e = epool + i;
        scanf("%d %d %d", &e->u, &e->v, &e->cost);
    }

    edge_t **edges = (edge_t **)malloc(sizeof(edge_t *) * m);
    for (int i = 0; i < m; i ++) {
        edges[i] = &epool[i];
    }

    esort(edges, 0, m - 1);

    int *ufs = ufs_init(n);
    int costs = 0;
    for (int i = 0; i < m; i ++) {
        edge_t *e = edges[i];
        if (ufs_join(ufs, e->u, e->v)) {
            costs += e->cost;
        }
    }

    for (int i = 2; i <= n; i ++) {
        if (ufs_find(ufs, i) != ufs_find(ufs, 1)) {
            costs = -1;
        }
    }
    printf("%d\n", costs);

    ufs_free(ufs);
    free(epool);
    free(edges);
    return 0;
}
