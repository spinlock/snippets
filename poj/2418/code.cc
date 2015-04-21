#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct node_s {
    char *text;
    char *keys;
    node_s **next;

    int hits;
    int size;
    bool tail;
} node_t;

node_t *
trie_node_init() {
    node_t *x = (node_t *)malloc(sizeof(node_t));
    x->size = 0;
    x->hits = 0;
    x->tail = false;
    return x;
}

int
trie_node_index(node_t *x, char key) {
    int beg = 0, end = x->size - 1;
    while (beg <= end) {
        int mid = (beg + end) / 2;
        if (x->keys[mid] == key) {
            return mid;
        } else if (x->keys[mid] < key) {
            beg = mid + 1;
        } else {
            end = mid - 1;
        }
    }
    return -(beg + 1);
}

node_t *
trie_node_insert(node_t *x, char key, int i) {
    char *keys = (char *)malloc(sizeof(char) * (x->size + 1));
    node_t **next = (node_t **)malloc(sizeof(node_t *) * (x->size + 1));
    for (int j = 0; j < i; j ++) {
        keys[j] = x->keys[j];
        next[j] = x->next[j];
    }
    keys[i] = key;
    next[i] = trie_node_init();
    for (int j = i; j < x->size; j ++) {
        keys[j + 1] = x->keys[j];
        next[j + 1] = x->next[j];
    }
    if (x->size != 0) {
        free(x->keys);
        free(x->next);
    }
    x->keys = keys;
    x->next = next;
    x->size++;
    return x->next[i];
}

void
trie_node_free(node_t *x) {
    if (x->size != 0) {
        free(x->keys);
        for (int i = 0; i < x->size; i ++) {
            trie_node_free(x->next[i]);
        }
        free(x->next);
    }
    if (x->tail) {
        free(x->text);
    }
    free(x);
}

int
trie_node_tail_walk(node_t *x, node_t **xs, int n) {
    if (x->tail) {
        xs[n] = x;
        n ++;
    }
    for (int i = 0; i < x->size; i ++) {
        n = trie_node_tail_walk(x->next[i], xs, n);
    }
    return n;
}

typedef struct {
    node_t *root;
    int size;
} tree_t;

void
trie_tree_init(tree_t *t) {
    t->root = trie_node_init();
    t->size = 0;
}

void
trie_tree_free(tree_t *t) {
    trie_node_free(t->root);
}

void
trie_tree_insert(tree_t *t, char *s) {
    node_t *x = t->root;
    for (char *p = s; *p != '\0'; p ++) {
        int i = trie_node_index(x, *p);
        if (i >= 0) {
            x = x->next[i];
        } else {
            x = trie_node_insert(x, *p, -(i + 1));
        }
    }
    if (!x->tail) {
        t->size ++;
        x->tail = true;
        x->text = strdup(s);
    }
    x->hits ++;
}

void
qswap(node_t **xs, int i, int j) {
    if (i != j) {
        node_t *t = xs[i];
        xs[i] = xs[j];
        xs[j] = t;
    }
}

void
qsort(node_t **xs, int beg, int end) {
    if (beg >= end) {
        return;
    }
    int pivot = beg;
    for (int j = beg + 1; j <= end; j ++) {
        if (strcmp(xs[j]->text, xs[beg]->text) <= 0) {
            pivot ++;
            qswap(xs, j, pivot);
        }
    }
    qswap(xs, beg, pivot);
    qsort(xs, beg, pivot - 1);
    qsort(xs, pivot + 1, end);
}

int
main(void) {
    char buf[1024];
    tree_t __t, *t = &__t;
    trie_tree_init(t);

    while (true) {
        char *s = fgets(buf, sizeof(buf) - 1, stdin);
        if (s == NULL) {
            break;
        }
        for (int i = 0; s[i] != '\0'; i ++) {
            if (s[i] == '\r' || s[i] == '\n') {
                s[i] = '\0';
                break;
            }
        }
        trie_tree_insert(t, s);
    }

    node_t **xs = (node_t **)malloc(sizeof(node_t *) * t->size);
    int n = trie_node_tail_walk(t->root, xs, 0);
    qsort(xs, 0, n - 1);
    int total = 0;
    for (int i = 0; i < n; i ++) {
        total += xs[i]->hits;
    }
    for (int i = 0; i < n; i ++) {
        printf("%s %.4f\n", xs[i]->text, float(xs[i]->hits * 100) / total);
    }
    free(xs);

    trie_tree_free(t);
    return 0;
}
