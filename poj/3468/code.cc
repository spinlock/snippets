#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

typedef struct {
    int beg, end;
    int64_t sum, add;
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
    t[x].add = 0;
    if (beg == end) {
        t[x].sum = (int64_t)values[beg];
    } else {
        int mid = (beg + end) / 2;
        int l = 2 * x + 1;
        int r = 2 * x + 2;
        intree_init_walk(t, l, values, beg, mid);
        intree_init_walk(t, r, values, mid + 1, end);
        t[x].sum = t[l].sum + t[r].sum;
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

int64_t
intree_sum_walk(node_t *t, int x, int beg, int end) {
    if (beg == t[x].beg && end == t[x].end) {
        return t[x].sum;
    } else {
        int mid = (t[x].beg + t[x].end) / 2;
        int l = 2 * x + 1;
        int r = 2 * x + 2;
        if (t[x].add != 0) {
            t[l].add += t[x].add;
            t[r].add += t[x].add;
            t[l].sum += (t[l].end - t[l].beg + 1) * t[x].add;
            t[r].sum += (t[r].end - t[r].beg + 1) * t[x].add;
            t[x].add = 0;
        }
        if (end <= mid) {
            return intree_sum_walk(t, l, beg, end);
        } else if (beg >= mid + 1) {
            return intree_sum_walk(t, r, beg, end);
        } else {
            int64_t suml = intree_sum_walk(t, l, beg, mid);
            int64_t sumr = intree_sum_walk(t, r, mid + 1, end);
            return suml + sumr;
        }
    }
    return 0;
}

int64_t
intree_sum(node_t *t, int beg, int end) {
    beg = max_int(beg, t[0].beg);
    end = min_int(end, t[0].end);
    if (beg > end) {
        return 0;
    }
    return intree_sum_walk(t, 0, beg, end);
}

void
intree_add_walk(node_t *t, int x, int beg, int end, int add) {
    if (beg == t[x].beg && end == t[x].end) {
        t[x].add += add;
    } else {
        int mid = (t[x].beg + t[x].end) / 2;
        int l = 2 * x + 1;
        int r = 2 * x + 2;
        if (end <= mid) {
            intree_add_walk(t, l, beg, end, add);
        } else if (beg >= mid + 1) {
            intree_add_walk(t, r, beg, end, add);
        } else {
            intree_add_walk(t, l, beg, mid, add);
            intree_add_walk(t, r, mid + 1, end, add);
        }
    }
    t[x].sum += (end - beg + 1) * add;
}

void
intree_add(node_t *t, int beg, int end, int add) {
    beg = max_int(beg, t[0].beg);
    end = min_int(end, t[0].end);
    if (beg > end || add == 0) {
        return;
    }
    intree_add_walk(t, 0, beg, end, add);
}

int
main(void) {
    int n, q;
    scanf("%d %d", &n, &q);
    int *values = (int *)malloc(sizeof(int) * (n + 1));
    values[0] = 0;
    for (int i = 1; i <= n; i ++) {
        scanf("%d", &values[i]);
    }
    node_t *t = intree_init(values, n + 1);
    for (int i = 0; i < q; i ++) {
        char op;
        int beg, end, add;
        scanf("\n%c", &op);
        if (op == 'Q') {
            scanf("%d %d", &beg, &end);
            printf("%lld\n", (long long)intree_sum(t, beg, end));
        } else if (op == 'C') {
            scanf("%d %d %d", &beg, &end, &add);
            intree_add(t, beg, end, add);
        }
    }
    free(t);
    free(values);
}
