#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int beg, end;
    long long sum, add;
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
tree_init_walk(node_t *t, int *values, int x, int beg, int end) {
    t[x].beg = beg;
    t[x].end = end;
    if (beg == end) {
        t[x].add = 0;
        t[x].sum = values[beg];
    } else {
        int mid = beg + (end - beg) / 2;
        int l = x * 2 + 1;
        int r = x * 2 + 2;
        tree_init_walk(t, values, l, beg, mid);
        tree_init_walk(t, values, r, mid + 1, end);
        t[x].add = 0;
        t[x].sum = t[l].sum + t[r].sum;
    }
}

node_t *
tree_init(int *values, int n) {
    node_t *t = (node_t *)malloc(sizeof(node_t) * (4 * n + 1));
    if (n != 0) {
        tree_init_walk(t, values, 0, 0, n - 1);
    }
    return t;
}

void
tree_free(node_t *t) {
    free(t);
}

long long
tree_sum_walk(node_t *t, int x, int beg, int end) {
    if (t[x].beg == beg && t[x].end == end) {
        return t[x].sum;
    } else {
        int mid = t[x].beg + (t[x].end - t[x].beg) / 2;
        int l = x * 2 + 1;
        int r = x * 2 + 2;
        if (t[x].add != 0) {
            t[l].add += t[x].add;
            t[l].sum += t[x].add * (t[l].end - t[l].beg + 1);
            t[r].add += t[x].add;
            t[r].sum += t[x].add * (t[r].end - t[r].beg + 1);
            t[x].add = 0;
        }
        if (end <= mid) {
            return tree_sum_walk(t, l, beg, end);
        } else if (beg > mid) {
            return tree_sum_walk(t, r, beg, end);
        } else {
            long long suml = tree_sum_walk(t, l, beg, mid);
            long long sumr = tree_sum_walk(t, r, mid + 1, end);
            return suml + sumr;
        }
    }
}

long long
tree_sum(node_t *t, int beg, int end) {
    beg = max_int(beg, t[0].beg);
    end = min_int(end, t[0].end);
    if (beg > end) {
        return 0;
    }
    return tree_sum_walk(t, 0, beg, end);
}

void
tree_add_walk(node_t *t, int x, int beg, int end, int add) {
    t[x].sum += (end - beg + 1) * add;
    if (t[x].beg == beg && t[x].end == end) {
        t[x].add += add;
    } else {
        int mid = t[x].beg + (t[x].end - t[x].beg) / 2;
        int l = x * 2 + 1;
        int r = x * 2 + 2;
        if (end <= mid) {
            tree_add_walk(t, l, beg, end, add);
        } else if (beg > mid) {
            tree_add_walk(t, r, beg, end, add);
        } else {
            tree_add_walk(t, l, beg, mid, add);
            tree_add_walk(t, r, mid + 1, end, add);
        }
    }
}

void
tree_add(node_t *t, int beg, int end, int add) {
    beg = max_int(beg, t[0].beg);
    end = min_int(end, t[0].end);
    if (beg > end) {
        return;
    }
    tree_add_walk(t, 0, beg, end, add);
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
        char op;
        int beg, end, add;
        scanf("\n%c", &op);
        if (op == 'Q') {
            scanf("%d %d", &beg, &end);
            printf("%lld\n", tree_sum(t, beg - 1, end - 1));
        } else if (op == 'C') {
            scanf("%d %d %d", &beg, &end, &add);
            tree_add(t, beg - 1, end - 1, add);
        }
    }
    tree_free(t);
    return 0;
}
