#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct node;

typedef struct node {
    int hits;
    bool leaf;
    struct node *left, *right;
} node_t;

node_t *
node_init() {
    int size = sizeof(node_t);
    node_t *n = (node_t *)malloc(size);
    return (node_t *)memset(n, 0, size);
}

void
node_free(node_t *n) {
    if (n == NULL) {
        return;
    }
    node_free(n->left);
    node_free(n->right);
    free(n);
}

typedef struct {
    node_t **buff;
    int size;
} heap_t;

void
heap_init(heap_t *h, int maxsize) {
    h->buff = (node_t **)malloc(sizeof(node_t *) * maxsize);
    h->size = 0;
}

void
heap_free(heap_t *h) {
    free(h->buff);
}

bool
heap_less(heap_t *h, int i, int j) {
    return h->buff[i]->hits < h->buff[j]->hits;
}

void
heap_swap(heap_t *h, int i, int j) {
    if (i != j) {
        node_t *n = h->buff[i];
        h->buff[i] = h->buff[j];
        h->buff[j] = n;
    }
}

void
heap_down(heap_t *h, int p) {
    while (p < h->size) {
        int l = p * 2 + 1;
        int r = p * 2 + 2;
        int m = p;
        if (l < h->size && heap_less(h, l, m)) {
            m = l;
        }
        if (r < h->size && heap_less(h, r, m)) {
            m = r;
        }
        if (m == p) {
            return;
        }
        heap_swap(h, m, p);
        p = m;
    }
}

void
heap_up(heap_t *h, int i) {
    while (i != 0) {
        int p = (i - 1) / 2;
        if (heap_less(h, p, i)) {
            return;
        }
        heap_swap(h, p, i);
        i = p;
    }
}

void
heap_rebuild(heap_t *h) {
    for (int i = h->size / 2; i >= 0; i --) {
        heap_down(h, i);
    }
}

node_t *
heap_pop(heap_t *h) {
    if (h->size == 0) {
        return NULL;
    }
    node_t *n = h->buff[0];
    h->size --;
    if (h->size != 0) {
        heap_swap(h, 0, h->size);
        heap_down(h, 0);
    }
    return n;
}

void
heap_push(heap_t *h, node_t *n) {
    h->buff[h->size] = n;
    h->size ++;
    heap_up(h, h->size - 1);
}

int
node_visit(int depth, node_t *n) {
    if (n == NULL) {
        return 0;
    }
    if (n->leaf) {
        return depth * n->hits;
    }
    int sum = 0;
    sum += node_visit(depth + 1, n->left);
    sum += node_visit(depth + 1, n->right);
    return sum;
}

void
process(char *s, int n) {
    const int max = 256;
    int hits[max];
    memset(hits, 0, sizeof(hits));
    for (int i = 0; i < n; i ++) {
        unsigned char b = s[i];
        hits[b] ++;
    }

    heap_t __heap, *h = &__heap;
    heap_init(h, max);

    for (int i = 0; i < max; i ++) {
        if (hits[i] != 0) {
            node_t *n = node_init();
            n->hits = hits[i];
            n->leaf = true;
            h->buff[h->size] = n;
            h->size ++;
        }
    }
    heap_rebuild(h);

    while (h->size != 1) {
        node_t *a = heap_pop(h);
        node_t *b = heap_pop(h);
        node_t *n = node_init();
        n->hits = a->hits + b->hits;
        n->left = a, n->right = b;
        heap_push(h, n);
    }
    node_t *root = heap_pop(h);

    int base = n * 8;
    int total;
    if (root->leaf) {
        total = n;
    } else {
        total = node_visit(0, root);
    }

    node_free(root);
    heap_free(h);

    printf("%d %d %.1f\n", base, total, float(base) / total);
}

int
main(void) {
    static char buff[1024*1024];
    while (true) {
        char *s = fgets(buff, sizeof(buff), stdin);
        if (s == NULL) {
            return 0;
        }
        int len = 0;
        for (char *p = s; *p != '\0' && *p != '\n'; p ++) {
            len ++;
        }
        if (strncmp(s, "END", len) == 0) {
            return 0;
        }
        process(s, len);
    }
    return 0;
}
