#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void
xswap(int *array, int i, int j) {
    if (i != j) {
        int t = array[i];
        array[i] = array[j];
        array[j] = t;
    }
}

void
qsort(int *array, int beg, int end) {
    if (beg >= end) {
        return;
    }
    int pivot = beg;
    for (int j = beg + 1; j <= end; j ++) {
        if (array[j] <= array[beg]) {
            pivot ++;
            xswap(array, j, pivot);
        }
    }
    xswap(array, beg, pivot);
    qsort(array, beg, pivot - 1);
    qsort(array, pivot + 1, end);
}

void
heap_down(int *array, int p, int size) {
    while (p < size) {
        int l = 2 * p + 1;
        int r = 2 * p + 2;
        int m = p;
        if (l < size && array[l] > array[m]) {
            m = l;
        }
        if (r < size && array[r] > array[m]) {
            m = r;
        }
        if (p == m) {
            return;
        }
        xswap(array, p, m);
        p = m;
    }
}

void
heap_build(int *array, int size) {
    for (int i = size / 2; i >= 0; i --) {
        heap_down(array, i, size);
    }
}

int
main(void) {
    int t;
    scanf("%d", &t);
    for (int l = 0; l < t; l ++) {
        int m, n;
        scanf("%d %d", &m, &n);

        int *ans = (int *)malloc(sizeof(int) * n);
        for (int j = 0; j < n; j ++) {
            scanf("%d", &ans[j]);
        }
        qsort(ans, 0, n - 1);

        int *tmp = (int *)malloc(sizeof(int) * n);
        for (int j = 0; j < n; j ++) {
            tmp[j] = ans[j];
        }
        heap_build(tmp, n);

        int *pls = (int *)malloc(sizeof(int) * n);
        for (int i = 1; i < m; i ++) {
            for (int j = 0; j < n; j ++) {
                scanf("%d", &pls[j]);
            }
            qsort(pls, 0, n - 1);

            for (int j = 0; j < n; j ++) {
                tmp[j] += pls[0];
            }

            for (int k = 1; k < n; k ++) {
                for (int j = 0; j < n; j ++) {
                    int v = ans[j] + pls[k];
                    if (v >= tmp[0]) {
                        break;
                    }
                    tmp[0] = v;
                    heap_down(tmp, 0, n);
                }
            }

            for (int j = 0; j < n; j ++) {
                ans[j] = tmp[j];
            }
            qsort(ans, 0, n - 1);
        }

        for (int j = 0; j < n; j ++) {
            printf("%d ", ans[j]);
        }
        printf("\n");

        free(ans);
        free(tmp);
        free(pls);
    }
    return 0;
}
