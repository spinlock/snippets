#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct node_s {
    node_s *next[10];
    bool tail;
} node_t;

node_t *
trie_node_init(node_t *x) {
    x->tail = false;
    memset(x->next, 0, sizeof(x->next));
    return x;
}

typedef struct {
    node_t *root;
    node_t *free_list;
    int next_free;
} tree_t;

node_t *
trie_tree_node_init(tree_t *t) {
    return trie_node_init(&t->free_list[t->next_free ++]);
}

void
trie_tree_init(tree_t *t, int n) {
    t->free_list = (node_t *)malloc(sizeof(node_t) * n);
    t->next_free = 0;
    t->root = trie_tree_node_init(t);
}

void
trie_tree_free(tree_t *t) {
    free(t->free_list);
}

bool
trie_tree_insert_noprefix(tree_t *t, char *s) {
    node_t *x = t->root;
    for (char *p = s; *p >= '0' && *p <= '9'; p ++) {
        if (x->tail) {
            return false;
        }
        int i = (*p) - '0';
        if (x->next[i] == NULL) {
            x->next[i] = trie_tree_node_init(t);
        }
        x = x->next[i];
    }
    if (x->tail) {
        return false;
    }
    for (int i = 0; i < 10; i ++) {
        if (x->next[i] != NULL) {
            return false;
        }
    }
    x->tail = true;
    return true;
}

void
process(int n, int depth) {
    tree_t __t, *t = &__t;
    trie_tree_init(t, n * depth);
    char buffer[1024];
    bool decodable = true;
    for (int i = 0; i < n; i ++) {
        char *s = fgets(buffer, sizeof(buffer) - 1, stdin);
        if (decodable) {
            decodable = trie_tree_insert_noprefix(t, s);
        }
    }
    if (decodable) {
        printf("YES\n");
    } else {
        printf("NO\n");
    }
    trie_tree_free(t);
}

int
main(void) {
    int t, n;
    scanf("%d\n", &t);
    for (int i = 0; i < t; i ++) {
        scanf("%d\n", &n);
        process(n, 40);
    }
    return 0;
}
