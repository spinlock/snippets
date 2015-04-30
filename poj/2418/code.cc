#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct node_s;
typedef struct node_s node_t;
typedef struct node_s {
    node_t *next[256];
    bool tail;
    int hits;
    char *text;
} node_t;

node_t *
node_alloc() {
    node_t *x = (node_t *)malloc(sizeof(node_t));
    memset(x->next, 0, sizeof(x->next));
    x->tail = false;
    x->hits = 0;
    x->text = NULL;
    return x;
}

void
node_free(node_t *x) {
    if (x == NULL) {
        return;
    }
    for (int i = 0; i < 256; i ++) {
        node_free(x->next[i]);
    }
    free(x->text);
    free(x);
}

node_t *
node_get(node_t *x, char b) {
    int i = (unsigned char)b;
    if (x->next[i] == NULL) {
        x->next[i] = node_alloc();
    }
    return x->next[i];
}

typedef struct {
    node_t *root;
    int size;
    int hits;
} tree_t;

void
tree_init(tree_t *t) {
    t->root = node_alloc();
    t->size = 0;
    t->hits = 0;
}

void
tree_free(tree_t *t) {
    node_free(t->root);
}

int
tree_nodes_walk(node_t **ns, int p, node_t *x) {
    if (x->tail) {
        ns[p ++] = x;
    }
    for (int i = 0; i < 256; i ++) {
        if (x->next[i] != NULL) {
            p = tree_nodes_walk(ns, p, x->next[i]);
        }
    }
    return p;
}

void
tree_insert(tree_t *t, char *s) {
    node_t *x = t->root;
    int i;
    for (i = 0; s[i] != '\0'; i ++) {
        if (s[i] == '\r' || s[i] == '\n') {
            s[i] = '\0';
            break;
        }
        x = node_get(x, s[i]);
    }
    if (!x->tail) {
        char *p = (char *)malloc(sizeof(char) * (i + 1));
        for (int j = 0; j <= i; j ++) {
            p[j] = s[j];
        }
        x->text = p;
        x->tail = true;
        t->size ++;
    }
    x->hits ++;
    t->hits ++;
}

int
main(void) {
    char buff[4096];
    tree_t __t, *t = &__t;
    tree_init(t);
    while (true) {
        char *s = fgets(buff, sizeof(buff), stdin);
        if (s == NULL) {
            break;
        }
        tree_insert(t, s);
    }
    node_t **ns = (node_t **)malloc(sizeof(node_t *) * t->size);
    tree_nodes_walk(ns, 0, t->root);
    for (int i = 0; i < t->size; i ++) {
        printf("%s %.4f\n", ns[i]->text, (100 * (double)ns[i]->hits) / t->hits);
    }
    free(ns);
    tree_free(t);
    return 0;
}
