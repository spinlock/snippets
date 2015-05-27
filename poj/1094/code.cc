#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void
toposort(int n, int **g, int u, int *order, int *visit) {
    if (visit[u] != 0) {
        return;
    }
    visit[u] = -1;
    int max = 0;
    for (int v = 0; v < n; v ++) {
        if (g[u][v] == 0) {
            continue;
        }
        toposort(n, g, v, order, visit);
        if (visit[v] < 0) {
            return;
        }
        int d = order[v] + 1;
        if (max < d) {
            max = d;
        }
    }
    visit[u] = 1;
    order[u] = max;
}

int
process(int n, int **g, int *order, int *visit) {
    for (int i = 0; i < n; i ++) {
        order[i] = visit[i] = 0;
    }
    for (int u = 0; u < n; u ++) {
        toposort(n, g, u, order, visit);
        if (visit[u] < 0) {
            return -1;
        }
    }
    for (int u = 0; u < n; u ++) {
        for (int v = u + 1; v < n; v ++) {
            if (order[u] == order[v]) {
                return 0;
            }
        }
    }
    return 1;
}

int
main(void) {
    while (true) {
        int m, n;
        scanf("%d %d\n", &n, &m);
        if (n == 0 || m == 0) {
            return 0;
        }
        int **g = (int **)malloc(sizeof(int *) * n);
        for (int i = 0; i < n; i ++) {
            int size = sizeof(int) * n;
            g[i] = (int *)malloc(size);
            memset(g[i], 0, size);
        }
        int *order = (int *)malloc(sizeof(int) * n);
        int *visit = (int *)malloc(sizeof(int) * n);
        int state = 0;
        for (int i = 0; i < m; i ++) {
            char a, b;
            scanf("%c<%c\n", &a, &b);
            if (state != 0) {
                continue;
            }
            int x = a - 'A', y = b - 'A';
            if (g[y][x] != 0) {
                continue;
            } else {
                g[y][x] = 1;
            }
            int s = process(n, g, order, visit);
            if (s < 0) {
                printf("Inconsistency found after %d relations.\n", i + 1);
                state = -1;
            } else if (s != 0){
                state = i + 1;
            }
        }
        if (state == 0) {
            printf("Sorted sequence cannot be determined.\n");
        } else if (state > 0) {
            printf("Sorted sequence determined after %d relations: ", state);
            for (int i = 0; i < n; i ++) {
                for (int u = 0; u < n; u ++) {
                    if (order[u] == i) {
                        printf("%c", u + 'A');
                    }
                }
            }
            printf(".\n");
        }
        free(order);
        free(visit);
        for (int i = 0; i < n; i ++) {
            free(g[i]);
        }
        free(g);
    }
    return 0;
}
