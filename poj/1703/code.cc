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
            ufs->dist[x] = (ufs->dist[x] + ufs->dist[px]) % 2;
        }
        return npx;
    }
}

int
ufs_ack(ufs_t *ufs, int x, int y) {
    int px = ufs_find(ufs, x);
    int py = ufs_find(ufs, y);
    if (px != py) {
        return 0;
    }
    if (ufs->dist[x] != ufs->dist[y]) {
        return 1;
    } else {
        return 2;
    }
}

void
ufs_diff(ufs_t *ufs, int x, int y) {
    int px = ufs_find(ufs, x);
    int py = ufs_find(ufs, y);
    if (px != py) {
        ufs->size[px] += ufs->size[py];
        ufs->size[py] = px;
        ufs->dist[py] = (ufs->dist[x] + 3 - ufs->dist[y]) % 2;
    }
}

int
main(void) {
    int t;
    scanf("%d", &t);
    for (; t != 0; t --) {
        int n, m;
        scanf("%d %d\n", &n, &m);
        ufs_t __ufs, *ufs = &__ufs;
        ufs_init(ufs, n + 1);
        for (; m != 0; m --) {
            char op;
            int x, y;
            scanf("%c %d %d\n", &op, &x, &y);
            if (op == 'A') {
                switch (ufs_ack(ufs, x, y)) {
                case 0:
                    printf("Not sure yet.\n");
                    break;
                case 1:
                    printf("In different gangs.\n");
                    break;
                case 2:
                    printf("In the same gang.\n");
                    break;
                }
            } else if (op == 'D') {
                ufs_diff(ufs, x, y);
            }
        }
        ufs_free(ufs);
    }
    return 0;
}
