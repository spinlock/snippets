#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int *
newarray(int n) {
    int size = sizeof(int) * (n + 1);
    int *p = (int *)malloc(size);
    return (int *)memset(p, 0, size);
}

int
max(int v1, int v2) {
    if (v1 > v2) {
        return v1;
    } else {
        return v2;
    }
}

int
main(void) {
    int n;
    scanf("%d", &n);
    int *a = newarray(n);
    for (int i = 1; i <= n; i ++) {
        scanf("%d", &a[i]);
    }
    int *s = newarray(n);
    for (int i = 1; i <= n; i ++) {
        s[i] = 1;
        for (int k = 1; k < i; k ++) {
            if (a[i] > a[k]) {
                s[i] = max(s[i], 1 + s[k]);
            }
        }
    }
    int mlen = 0;
    for (int i = 1; i <= n; i ++) {
        mlen = max(mlen, s[i]);
    }
    printf("%d\n", mlen);
    free(s);
    free(a);
    return 0;
}
