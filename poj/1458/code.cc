#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool
isspace(char *p) {
    const char *s = " \t\r\n";
    for (int i = strlen(s) - 1; i >= 0; i --) {
        if (*p == s[i]) {
            return true;
        }
    }
    return false;
}

char *
trimprefix(char *p) {
    while (*p != '\0' && isspace(p)) {
        p ++;
    }
    return p;
}

int
maxlen(char *p) {
    int len = 0;
    while (*p != '\0' && !isspace(p)) {
        p ++, len ++;
    }
    return len;
}

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
process(char *input) {
    char *a = trimprefix(input);
    int m = maxlen(a);
    if (m == 0) {
        return 0;
    }
    char *b = trimprefix(a + m);
    int n = maxlen(b);
    if (n == 0) {
        return 0;
    }

    int *so = newarray(n);
    int *sn = newarray(n);

    for (int i = 1; i <= m; i ++) {
        for (int j = 1; j <= n; j ++) {
            if (a[i - 1] == b[j - 1]) {
                sn[j] = 1 + so[j - 1];
            } else {
                sn[j] = max(sn[j - 1], so[j]);
            }
        }
        int *t = so; so = sn; sn = t;
    }

    int max = so[n];

    free(so);
    free(sn);
    return max;
}

int
main(void) {
    static char buff[1024*1024];
    while (true) {
        char *s = fgets(buff, sizeof(buff), stdin);
        if (s == NULL) {
            return 0;
        }
        printf("%d\n", process(s));
    }
    return 0;
}
