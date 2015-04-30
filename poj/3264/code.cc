#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int beg, end;
    int min, max;
} node_t;

typedef struct {
    node_t *nodes;
} tree_t;

int
min_int(int v1, int v2) {
    if (v1 < v2) {
        return v1;
    } else {
        return v2;
    }
}

int
max_int(int v1, int v2) {
    if (v1 > v2) {
        return v1;
    } else {
        return v2;
    }
}

void
tree_init_walk(tree_t *t, int p, int *values, int beg, int end) {
    node_t *x = &t->nodes[p];
    x->beg = beg;
    x->end = end;
    if (beg == end) {
        x->min = values[beg];
        x->max = values[end];
    } else {
        int mid = beg + (end - beg) / 2;
        int l = p * 2 + 1;
        int r = p * 2 + 2;
        tree_init_walk(t, l, values, beg, mid);
        tree_init_walk(t, r, values, mid + 1, end);
        x->min = min_int(t->nodes[l].min, t->nodes[r].min);
        x->max = max_int(t->nodes[l].max, t->nodes[r].max);
    }
}

void
tree_init(tree_t *t, int *values, int n) {
    int max = 4 * n + 1;
    t->nodes = (node_t *)malloc(sizeof(node_t) * max);
    memset(t->nodes, 0, sizeof(node_t) * max);
    if (n != 0) {
        tree_init_walk(t, 0, values, 0, n - 1);
    }
}

void
tree_free(tree_t *t) {
    free(t->nodes);
}

void
tree_minmax_walk(tree_t *t, int p, int beg, int end, int *pmin, int *pmax) {
    node_t *x = &t->nodes[p];
    int min, max;
    if (x->beg == beg && x->end == end) {
        min = x->min;
        max = x->max;
    } else {
        int mid = x->beg + (x->end - x->beg) / 2;
        int l = p * 2 + 1;
        int r = p * 2 + 2;
        if (end <= mid) {
            tree_minmax_walk(t, l, beg, end, &min, &max);
        } else if (beg > mid) {
            tree_minmax_walk(t, r, beg, end, &min, &max);
        } else {
            int minl, maxl;
            tree_minmax_walk(t, l, beg, mid, &minl, &maxl);
            int minr, maxr;
            tree_minmax_walk(t, r, mid + 1, end, &minr, &maxr);
            min = min_int(minl, minr);
            max = max_int(maxl, maxr);
        }
    }
    *pmin = min;
    *pmax = max;
}

int
tree_delta(tree_t *t, int beg, int end) {
    beg = max_int(beg, t->nodes[0].beg);
    end = min_int(end, t->nodes[0].end);
    if (beg > end) {
        return 0;
    }
    int min, max;
    tree_minmax_walk(t, 0, beg, end, &min, &max);
    return max - min;
}

int
main(void) {
    int n, m;
    scanf("%d %d", &n, &m);

    int *values = (int *)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i ++) {
        scanf("%d", &values[i]);
    }

    tree_t __t, *t = &__t;
    tree_init(t, values, n);
    for (int i = 0; i < m; i ++) {
        int beg, end;
        scanf("%d %d", &beg, &end);
        printf("%d\n", tree_delta(t, beg - 1, end - 1));
    }
    tree_free(t);
    return 0;
}
