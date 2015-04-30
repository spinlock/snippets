#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int beg, end;
    int min, max;
} node_t;

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
tree_init_walk(node_t *t, int x, int *values, int beg, int end) {
    t[x].beg = beg;
    t[x].end = end;
    if (beg == end) {
        t[x].min = values[beg];
        t[x].max = values[end];
    } else {
        int mid = beg + (end - beg) / 2;
        int l = x * 2 + 1;
        int r = x * 2 + 2;
        tree_init_walk(t, l, values, beg, mid);
        tree_init_walk(t, r, values, mid + 1, end);
        t[x].min = min_int(t[l].min, t[r].min);
        t[x].max = max_int(t[l].max, t[r].max);
    }
}

node_t *
tree_init(int *values, int n) {
    int max = 4 * n + 1;
    node_t *t = (node_t *)malloc(sizeof(node_t) * max);
    memset(t, 0, sizeof(node_t) * max);
    if (n != 0) {
        tree_init_walk(t, 0, values, 0, n - 1);
    }
    return t;
}

void
tree_free(node_t *t) {
    free(t);
}

void
tree_minmax_walk(node_t *t, int x, int beg, int end, int *pmin, int *pmax) {
    int min, max;
    if (t[x].beg == beg && t[x].end == end) {
        min = t[x].min;
        max = t[x].max;
    } else {
        int mid = t[x].beg + (t[x].end - t[x].beg) / 2;
        int l = x * 2 + 1;
        int r = x * 2 + 2;
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
tree_delta(node_t *t, int beg, int end) {
    beg = max_int(beg, t[0].beg);
    end = min_int(end, t[0].end);
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

    node_t *t = tree_init(values, n);
    for (int i = 0; i < m; i ++) {
        int beg, end;
        scanf("%d %d", &beg, &end);
        printf("%d\n", tree_delta(t, beg - 1, end - 1));
    }
    tree_free(t);
    return 0;
}
