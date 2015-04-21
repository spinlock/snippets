#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int *array, *mod;
} ufs_t;

void ufs_init(ufs_t *g, int n) {
    g->array = (int *)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i ++) {
        g->array[i] = -1;
    }
    g->mod = (int *)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i ++) {
        g->mod[i] = 0;
    }
}

void ufs_free(ufs_t *g) {
    free(g->array);
    free(g->mod);
}

int ufs_find(ufs_t *g, int x) {
    int px = g->array[x];
    if (px < 0) {
        return x;
    }
    g->array[x] = ufs_find(g, px);
    g->mod[x] = (g->mod[x] + g->mod[px]) % 3;
    return g->array[x];
}

bool ufs_union(ufs_t *g, int x, int y, int eat) {
    int px = ufs_find(g, x);
    int py = ufs_find(g, y);
    if (px == py) {
        return (g->mod[x] + eat) % 3 == g->mod[y];
    } else {
        g->array[px] += g->array[py];
        g->array[py] = px;
        g->mod[py] = (g->mod[x] + 3 + eat - g->mod[y]) % 3;
        return true;
    }
}

bool check(int x, int beg, int end) {
    return x >= beg && x <= end;
}

int main(void) {
    int n, m;
    scanf("%d %d", &n, &m);
    ufs_t __g, *g = &__g;
    ufs_init(g, n + 1);
    int failed = 0;
    for (int i = 0; i < m; i ++) {
        int d, x, y;
        scanf("%d %d %d", &d, &x, &y);
        if (!check(x, 1, n) || !check(y, 1, n)) {
            failed ++;
        } else if (d == 1) {
            if (!ufs_union(g, x, y, 0)) {
                failed ++;
            }
        } else if (d == 2) {
            if (!ufs_union(g, x, y, 1)) {
                failed ++;
            }
        }
    }
    ufs_free(g);
    printf("%d\n", failed);
    return 0;
}
