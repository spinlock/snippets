#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct node_s {
    node_s *next[2];
    bool tail;
} node_t;

node_t *
trie_node_init() {
    node_t *x = (node_t *)malloc(sizeof(node_t));
    x->tail = false;
    memset(x->next, 0, sizeof(x->next));
    return x;
}

void
trie_node_free(node_t *x) {
    if (x == NULL) {
        return;
    }
    for (int i = 0; i < 2; i ++) {
        trie_node_free(x->next[i]);
    }
    free(x);
}

typedef struct {
    node_t *root;
} tree_t;

void
trie_tree_init(tree_t *t) {
    t->root = trie_node_init();
}

void
trie_tree_free(tree_t *t) {
    trie_node_free(t->root);
}

bool
trie_tree_insert_noprefix(tree_t *t, char *s) {
    node_t *x = t->root;
    for (char *p = s; *p == '0' || *p == '1'; p ++) {
        if (x->tail) {
            return false;
        }
        int i = (*p) - '0';
        if (x->next[i] == NULL) {
            x->next[i] = trie_node_init();
        }
        x = x->next[i];
    }
    if (x->tail) {
        return false;
    }
    for (int i = 0; i < 2; i ++) {
        if (x->next[i] != NULL) {
            return false;
        }
    }
    x->tail = true;
    return true;
}

int
process(tree_t *t) {
    char buf[1024];
    bool decodable = true;
    while (true) {
        char *s = fgets(buf, sizeof(buf), stdin);
        if (s == NULL) {
            return -1;
        }
        if (s[0] == '9') {
            break;
        }
        if (decodable) {
            decodable = trie_tree_insert_noprefix(t, s);
        }
    }
    return decodable ? 0 : 1;
}

int
main(void) {
    char buffer[1024];
    for (int i = 1; ; i ++) {
        tree_t __t, *t = &__t;
        trie_tree_init(t);
        int ret = process(t);
        trie_tree_free(t);
        if (ret < 0) {
            return 0;
        } else if (ret == 0) {
            printf("Set %d is immediately decodable\n", i);
        } else {
            printf("Set %d is not immediately decodable\n", i);
        }
    }
}
