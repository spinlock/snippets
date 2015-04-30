#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void
xswap(int *array, int i, int j) {
    if (i != j) {
        int t = array[i]; array[i] = array[j]; array[j] = t;
    }
}

void
qsort(int *array, int beg, int end) {
    if (beg >= end) {
        return;
    }
    int pivot = beg;
    for (int i = beg + 1; i <= end; i ++) {
        if (array[i] <= array[beg]) {
            pivot ++;
            xswap(array, i, pivot);
        }
    }
    xswap(array, beg, pivot);
    qsort(array, beg, pivot - 1);
    qsort(array, pivot + 1, end);
}

void
hdown(int *array, int p, int size) {
    while (p < size) {
        int l = p * 2 + 1;
        int r = p * 2 + 2;
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

int
main(void) {
    int t, m, n;
    scanf("%d", &t);
    for (; t != 0; t --) {
        scanf("%d %d", &m, &n);
        int *sum = (int *)malloc(sizeof(int) * n);
        int *pls = (int *)malloc(sizeof(int) * n);
        int *tmp = (int *)malloc(sizeof(int) * n);

        for (int i = 0; i < n; i ++) {
            scanf("%d", &sum[i]);
        }
        for (int i = n / 2; i >= 0; i --) {
            hdown(sum, i, n);
        }

        for (int l = 1; l < m; l ++) {
            for (int i = 0; i < n; i ++) {
                tmp[i] = sum[i];
            }
            qsort(tmp, 0, n - 1);

            for (int i = 0; i < n; i ++) {
                scanf("%d", &pls[i]);
            }
            qsort(pls, 0, n - 1);

            for (int i = 0; i < n; i ++) {
                sum[i] += pls[0];
            }

            for (int i = 1; i < n; i ++) {
                for (int j = 0; j < n; j ++) {
                    int x = pls[i] + tmp[j];
                    if (x >= sum[0]) {
                        break;
                    }
                    sum[0] = x;
                    hdown(sum, 0, n);
                }
            }
        }

        qsort(sum, 0, n - 1);
        for (int i = 0; i < n; i ++) {
            printf("%d ", sum[i]);
        }
        printf("\n");

        free(sum);
        free(pls);
        free(tmp);
    }
    return 0;
}
