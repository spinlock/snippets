#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

typedef struct {
    int size;
    long long *buff;
} heap_t;

void
heap_init(heap_t *h, int size) {
    h->size = 0;
    h->buff = (long long *)malloc(sizeof(long long) * size);
}

void
heap_free(heap_t *h) {
    free(h->buff);
}

void
heap_swap(heap_t *h, int i, int j) {
    long long t = h->buff[i];
    h->buff[i] = h->buff[j];
    h->buff[j] = t;
}

void
heap_up(heap_t *h, int i) {
    while (i != 0) {
        int p = (i - 1) / 2;
        if (h->buff[p] < h->buff[i]) {
            return;
        }
        heap_swap(h, p, i);
        i = p;
    }
}

void
heap_down(heap_t *h, int p) {
    while (p < h->size) {
        int l = p * 2 + 1;
        int r = p * 2 + 2;
        int m = p;
        if (l < h->size && h->buff[l] < h->buff[m]) {
            m = l;
        }
        if (r < h->size && h->buff[r] < h->buff[m]) {
            m = r;
        }
        if (p == m) {
            return;
        }
        heap_swap(h, p, m);
        p = m;
    }
}

long long
heap_pop(heap_t *h) {
    if (h->size == 0) {
        return -1;
    }
    long long v = h->buff[0];
    h->size --;
    if (h->size != 0) {
        heap_swap(h, 0, h->size);
        heap_down(h, 0);
    }
    return v;
}

void
heap_push(heap_t *h, long long v) {
    h->buff[h->size] = v;
    h->size ++;
    heap_up(h, h->size - 1);
}

int
main(void) {
    int n;
    scanf("%d", &n);
    heap_t __h, *h = &__h;
    heap_init(h, n);
    for (int i = 0; i < n; i ++) {
        scanf("%lld", &h->buff[i]);
    }
    h->size = n;

    for (int i = h->size / 2; i >= 0; i --) {
        heap_down(h, i);
    }
    long long sum = 0;
    while (h->size != 1) {
        long long v1 = heap_pop(h);
        long long v2 = heap_pop(h);
        long long cut = v1 + v2;
        heap_push(h, cut);
        sum += cut;
    }
    printf("%lld\n", sum);
    return 0;
}
