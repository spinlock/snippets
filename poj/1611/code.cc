#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *ufs_init(int n) {
    int *array = (int *)malloc(sizeof(int) * n);
    for (int i = 0; i < n; i ++) {
        array[i] = -1;
    }
    return array;
}

void ufs_free(int *array) {
    free(array);
}

int ufs_find(int *array, int x) {
    int px = array[x];
    if (px < 0) {
        return x;
    }
    array[x] = ufs_find(array, px);
    return array[x];
}

int ufs_size(int *array, int x) {
    int px = ufs_find(array, x);
    return -array[px];
}

void ufs_union(int *array, int x, int y) {
    int px = ufs_find(array, x);
    int py = ufs_find(array, y);
    if (px != py) {
        array[px] += array[py];
        array[py] = px;
    }
}

int main(void) {
    while (true) {
        int n, m;
        scanf("%d %d", &n, &m);
        if (n == 0) {
            return 0;
        }
        int *array = ufs_init(n);
        for (int i = 0; i < m; i ++) {
            int k, x, y;
            scanf("%d %d", &k, &x);
            for (int j = 1; j < k; j ++) {
                scanf("%d", &y);
                ufs_union(array, x, y);
            }
        }
        printf("%d\n", ufs_size(array, 0));
        ufs_free(array);
    }
    return 0;
}
