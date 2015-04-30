#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int *size;
} ufs_t;

void
ufs_init(ufs_t *ufs, int n) {
    ufs->size = (int *)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i ++) {
        ufs->size[i] = -1;
    }
}

void
ufs_free(ufs_t *ufs) {
    free(ufs->size);
}

int
ufs_find(ufs_t *ufs, int x) {
    int px = ufs->size[x];
    if (px < 0) {
        return x;
    } else {
        int npx = ufs_find(ufs, px);
        ufs->size[x] = npx;
        return npx;
    }
}

int
ufs_size(ufs_t *ufs, int x) {
    int px = ufs_find(ufs, x);
    return -ufs->size[px];
}

void
ufs_union(ufs_t *ufs, int x, int y) {
    int px = ufs_find(ufs, x);
    int py = ufs_find(ufs, y);
    if (px != py) {
        ufs->size[px] += ufs->size[py];
        ufs->size[py] = px;
    }
}

int
main(void) {
    while (true) {
        int n, m;
        scanf("%d %d", &n, &m);
        if (n == 0) {
            return 0;
        }
        ufs_t __ufs, *ufs = &__ufs;
        ufs_init(ufs, n);
        for (int i = 0; i < m; i ++) {
            int k, x, y;
            scanf("%d %d", &k, &x);
            for (int j = 1; j < k; j ++) {
                scanf("%d", &y);
                ufs_union(ufs, x, y);
            }
        }
        printf("%d\n", ufs_size(ufs, 0));
        ufs_free(ufs);
    }
    return 0;
}
