#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const int m = 44;
const int n = 1995;

int g[n + 1][m + 1];
int d[m + 1];

void
reverse(int *path, int i, int j) {
    while (i < j) {
        int t = path[i];
        path[i] = path[j];
        path[j] = t;
        i ++;
        j --;
    }
}

int
eular(int u, int *path, int size) {
    for (int e = 1; e <= n; e ++) {
        int v = g[e][u];
        if (v != 0) {
            g[e][u] = g[e][v] = 0;
            size = eular(v, path, size);
            path[size] = e;
            size ++;
        }
    }
    return size;
}

int
process(void) {
    memset(g, 0, sizeof(g));
    memset(d, 0, sizeof(d));

    for (int i = 0; ; i ++) {
        int x, y, z;
        scanf("%d %d", &x, &y);
        if (x == 0 || y == 0) {
            if (i == 0) {
                return -1;
            }
            break;
        }
        scanf("%d", &z);
        d[x] ++;
        d[y] ++;
        g[z][x] = y;
        g[z][y] = x;
    }

    for (int u = 1; u <= m; u ++) {
        if (d[u] % 2 != 0) {
            printf("Round trip does not exist.\n");
            return 0;
        }
    }

    int path[n];
    int size = eular(1, path, 0);
    reverse(path, 0, size - 1);
    for (int i = 0; i < size; i ++) {
        printf("%d ", path[i]);
    }
    printf("\n");
    return 0;
}

int
main(void) {
    while (true) {
        if (process() != 0) {
            break;
        }
    }
    return 0;
}
