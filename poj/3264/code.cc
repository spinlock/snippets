#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

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
intree_init_walk(node_t *t, int x, int *values, int beg, int end) {
    t[x].beg = beg;
    t[x].end = end;
    if (beg == end) {
        t[x].min = values[beg];
        t[x].max = values[beg];
    } else {
        int mid = (beg + end) / 2;
        int l = 2 * x + 1;
        int r = 2 * x + 2;
        intree_init_walk(t, l, values, beg, mid);
        intree_init_walk(t, r, values, mid + 1, end);
        t[x].min = min_int(t[l].min, t[r].min);
        t[x].max = max_int(t[l].max, t[r].max);
    }
}

node_t *
intree_init(int *values, int n) {
    node_t *t = (node_t *)malloc(sizeof(node_t) * (4 * n + 1));
    if (n != 0) {
        intree_init_walk(t, 0, values, 0, n - 1);
    } else {
        t[0].end = t[0].beg - 1;
    }
    return t;
}

void
intree_minmax_walk(node_t *t, int x, int beg, int end, int *pmin, int *pmax) {
    if (beg == t[x].beg && end == t[x].end) {
        *pmin = min_int(*pmin, t[x].min);
        *pmax = max_int(*pmax, t[x].max);
    } else {
        int mid = (t[x].beg + t[x].end) / 2;
        int l = 2 * x + 1;
        int r = 2 * x + 2;
        if (mid >= end) {
            intree_minmax_walk(t, l, beg, end, pmin, pmax);
        } else if (beg >= mid + 1) {
            intree_minmax_walk(t, r, beg, end, pmin, pmax);
        } else {
            intree_minmax_walk(t, l, beg, mid, pmin, pmax);
            intree_minmax_walk(t, r, mid + 1, end, pmin, pmax);
        }
    }
}

void
intree_minmax(node_t *t, int beg, int end, int *pmin, int *pmax) {
    beg = max_int(beg, t[0].beg);
    end = min_int(end, t[0].end);
    if (beg > end) {
        *pmin = *pmax = 0;
    } else {
        *pmin = INT_MAX, *pmax = INT_MIN;
        intree_minmax_walk(t, 0, beg, end, pmin, pmax);
    }
}

int
main(void) {
    int n, q;
    scanf("%d %d", &n, &q);
    int *height = (int *)malloc(sizeof(int) * (n + 1));
    height[0] = 0;
    for (int i = 1; i <= n; i ++) {
        scanf("%d", &height[i]);
    }
    node_t *tree = intree_init(height, n + 1);
    for (int i = 0; i < q; i ++) {
        int beg, end;
        int min, max;
        scanf("%d %d", &beg, &end);
        intree_minmax(tree, beg, end, &min, &max);
        printf("%d\n", max - min);
    }
    free(tree);
    free(height);
    return 0;
}
