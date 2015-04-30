#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int *size;
    int *dist;
} ufs_t;

void
ufs_init(ufs_t *ufs, int n) {
    ufs->size = (int *)malloc(sizeof(int) * n);
    ufs->dist = (int *)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i ++) {
        ufs->size[i] = -1;
        ufs->dist[i] = 0;
    }
}

void
ufs_free(ufs_t *ufs) {
    free(ufs->size);
    free(ufs->dist);
}

int
ufs_find(ufs_t *ufs, int x) {
    int px = ufs->size[x];
    if (px < 0) {
        return x;
    } else {
        int npx = ufs_find(ufs, px);
        if (npx != px) {
            ufs->size[x] = npx;
            ufs->dist[x] = (ufs->dist[x] + ufs->dist[px]) % 3;
        }
        return npx;
    }
}

bool
ufs_union(ufs_t *ufs, int x, int y, int eat) {
    int px = ufs_find(ufs, x);
    int py = ufs_find(ufs, y);
    if (px == py) {
        return ufs->dist[y] == (ufs->dist[x] + eat) % 3;
    } else {
        ufs->size[px] += ufs->size[py];
        ufs->size[py] = px;
        ufs->dist[py] = (ufs->dist[x] + 3 - ufs->dist[y] + eat) % 3;
        return true;
    }
}

int
main(void) {
    int n, k;
    scanf("%d %d", &n, &k);
    ufs_t __ufs, *ufs = &__ufs;
    ufs_init(ufs, n + 1);
    int lie = 0;
    for (int i = 0; i < k; i ++) {
        int op, x, y;
        scanf("%d %d %d", &op, &x, &y);
        int eat = (op == 1) ? 0 : 1;
        if (x > n || y > n) {
            lie ++;
        } else if (!ufs_union(ufs, x, y, eat)) {
            lie ++;
        }
    }
    ufs_free(ufs);
    printf("%d\n", lie);
    return 0;
}
