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
    g->mod[x] = (g->mod[x] + g->mod[px]) % 2;
    return g->array[x];
}

bool ufs_union(ufs_t *g, int x, int y) {
    int px = ufs_find(g, x);
    int py = ufs_find(g, y);
    if (px == py) {
        return g->mod[x] != g->mod[y];
    } else {
        g->array[px] += g->array[py];
        g->array[py] = px;
        g->mod[py] = (g->mod[x] + 3 - g->mod[y]) % 2;
        return true;
    }
}

int ufs_gang(ufs_t *g, int x, int y) {
    int px = ufs_find(g, x);
    int py = ufs_find(g, y);
    if (px != py) {
        return -1;
    } else if (g->mod[x] == g->mod[y]) {
        return 1;
    } else {
        return 0;
    }
}

void process() {
    int n, m;
    scanf("%d %d\n", &n, &m);
    ufs_t __g, *g = &__g;
    ufs_init(g, n + 1);
    for (int i = 0; i < m; i ++) {
        char o;
        int x, y;
        scanf("%c %d %d\n", &o, &x, &y);
        if (o == 'D') {
            ufs_union(g, x, y);
        } else if (o == 'A') {
            int m = ufs_gang(g, x, y);
            if (m == -1) {
                printf("Not sure yet.\n");
            } else if (m == 0) {
                printf("In different gangs.\n");
            } else {
                printf("In the same gang.\n");
            }
        }
    }
    ufs_free(g);
}

int main(void) {
    int t;
    scanf("%d", &t);
    for (int i = 0; i < t; i ++) {
        process();
    }
}
