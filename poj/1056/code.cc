#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct node_s;
typedef struct node_s node_t;
typedef struct node_s {
    node_t *next[2];
    bool tail;
} node_t;

node_t *
node_alloc() {
    node_t *x = (node_t *)malloc(sizeof(node_t));
    for (int i = 0; i < 2; i ++) {
        x->next[i] = NULL;
    }
    x->tail = false;
    return x;
}

void
node_free(node_t *x) {
    if (x == NULL) {
        return;
    }
    for (int i = 0; i < 2; i ++) {
        node_free(x->next[i]);
    }
    free(x);
}

bool
node_has_next(node_t *x) {
    for (int i = 0; i < 2; i ++) {
        if (x->next[i] != NULL) {
            return true;
        }
    }
    return false;
}

node_t *
node_get_next(node_t *x, int i) {
    if (x->next[i] == NULL) {
        x->next[i] = node_alloc();
    }
    return x->next[i];
}

typedef struct {
    node_t *root;
} tree_t;

void
tree_init(tree_t *t) {
    t->root = node_alloc();
}

void
tree_free(tree_t *t) {
    node_free(t->root);
}

bool
tree_insert_noprefix(tree_t *t, char *s) {
    node_t *x = t->root;
    for (int i = 0; s[i] == '0' || s[i] == '1'; i ++) {
        if (x->tail) {
            return false;
        }
        x = node_get_next(x, s[i] - '0');
    }
    if (x->tail || node_has_next(x)) {
        return false;
    } else {
        x->tail = true;
        return true;
    }
}

int
main(void) {
    char buff[4096];
    for (int i = 1; ; i ++) {
        bool decodable = true;
        tree_t __t, *t = &__t;
        tree_init(t);
        while (true) {
            char *s = fgets(buff, sizeof(buff), stdin);
            if (s == NULL) {
                tree_free(t);
                return 0;
            }
            if (s[0] == '9') {
                break;
            }
            if (decodable) {
                decodable = tree_insert_noprefix(t, s);
            }
        }
        tree_free(t);
        if (decodable) {
            printf("Set %d is immediately decodable\n", i);
        } else {
            printf("Set %d is not immediately decodable\n", i);
        }
    }
}
