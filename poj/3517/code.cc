#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct node_s {
    int key;
    int size;
    node_s *left, *right;
} node_t;

typedef struct {
    node_t *nodes;
    node_t *root;
    node_t nil;
} tree_t;

node_t *
tree_init_walk(tree_t *t, int beg, int end) {
    if (beg > end) {
        return &t->nil;
    } else {
        int mid = beg + (end - beg) / 2;
        node_t *x = &t->nodes[mid];
        x->key = mid;
        x->size = end - beg + 1;
        x->left = tree_init_walk(t, beg, mid - 1);
        x->right = tree_init_walk(t, mid + 1, end);
        return x;
    }
}

void
tree_init(tree_t *t, int n) {
    t->nil.size = 0;
    t->nil.left = &t->nil, t->nil.right = &t->nil;
    t->nodes = (node_t *)malloc(sizeof(node_t) * n);
    t->root = tree_init_walk(t, 0, n - 1);
}

void
tree_free(tree_t *t) {
    free(t->nodes);
}

node_t *
tree_find_min(node_t *x) {
    while (x->left->size != 0) {
        x = x->left;
    }
    return x;
}

node_t *
tree_find_max(node_t *x) {
    while (x->right->size != 0) {
        x = x->right;
    }
    return x;
}

void tree_lbalance(node_t **p);
void tree_rbalance(node_t **p);
void tree_maintain(node_t **p);

void
tree_lrotate(node_t **p) {
    node_t *x = *p;
    node_t *y = x->right;
    x->right = y->left;
    y->left = x;
    y->size = x->size;
    x->size = x->left->size + x->right->size + 1;
    *p = y;
}

void
tree_rrotate(node_t **p) {
    node_t *x = *p;
    node_t *y = x->left;
    x->left = y->right;
    y->right = x;
    y->size = x->size;
    x->size = x->left->size + x->right->size + 1;
    *p = y;
}

void
tree_lbalance(node_t **p) {
    node_t *x = *p;
    if (x->right->size < x->left->left->size) {
        tree_rrotate(&x);
    } else if (x->right->size < x->left->right->size) {
        tree_lrotate(&x->left);
        tree_rrotate(&x);
    } else {
        return;
    }
    tree_rbalance(&x->right);
    tree_lbalance(&x->left);
    tree_maintain(&x);
    *p = x;
}

void
tree_rbalance(node_t **p) {
    node_t *x = *p;
    if (x->left->size < x->right->right->size) {
        tree_lrotate(&x);
    } else if (x->left->size < x->right->left->size) {
        tree_rrotate(&x->right);
        tree_lrotate(&x);
    } else {
        return;
    }
    tree_lbalance(&x->left);
    tree_rbalance(&x->right);
    tree_maintain(&x);
    *p = x;
}

void
tree_maintain(node_t **p) {
    tree_lbalance(p);
    tree_rbalance(p);
}

bool
tree_delete_walk(tree_t *t, node_t **p, int rank, int *pkey) {
    node_t *x = *p;
    if (rank < 0) {
        rank += x->size;
    }
    bool updated = false;
    if (rank < 0 || rank >= x->size) {
        return false;
    } else if (x->left->size == rank) {
        updated = true;
        if (pkey != NULL) {
            *pkey = x->key;
        }
        if (x->left->size > x->right->size) {
            node_t *m = tree_find_max(x->left);
            x->key = m->key;
            tree_delete_walk(t, &x->left, -1, NULL);
        } else if (x->right->size != 0) {
            node_t *m = tree_find_min(x->right);
            x->key = m->key;
            tree_delete_walk(t, &x->right, 0, NULL);
        }
    } else if (x->left->size < rank) {
        rank -= x->left->size + 1;
        updated = tree_delete_walk(t, &x->right, rank, pkey);
    } else {
        updated = tree_delete_walk(t, &x->left, rank, pkey);
    }
    if (updated) {
        x->size --;
        if (x->size == 0) {
            *p = &t->nil;
        } else {
            tree_maintain(&x);
            *p = x;
        }
    }
    return updated;
}

bool
tree_delete(tree_t *t, int rank, int *pkey) {
    return tree_delete_walk(t, &t->root, rank, pkey);
}

int
tree_size(tree_t *t) {
    return t->root->size;
}

int
main(void) {
    int n, k, m;
    while (true) {
        scanf("%d %d %d", &n, &k, &m);
        if (n == 0) {
            return 0;
        }
        tree_t __t, *t = &__t;
        tree_init(t, n);
        int rank = m - 1;
        int last;
        while (tree_size(t) != 0) {
            rank = rank % tree_size(t);
            tree_delete(t, rank, &last);
            rank = rank + k - 1;
        }
        printf("%d\n", last + 1);
        tree_free(t);
    }
    return 0;
}
