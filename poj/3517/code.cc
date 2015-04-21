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
} sbtree_t;

node_t *
sbtree_init_walk(sbtree_t *t, int *keys, int beg, int end) {
    if (beg > end) {
        return &t->nil;
    } else if (beg == end) {
        node_t *x = &t->nodes[beg];
        x->key = keys[beg];
        x->size = 1;
        x->left = &t->nil;
        x->right = &t->nil;
        return x;
    } else {
        int mid = (beg + end) / 2;
        node_t *x = &t->nodes[mid];
        x->key = keys[mid];
        x->size = end - beg + 1;
        x->left = sbtree_init_walk(t, keys, beg, mid - 1);
        x->right = sbtree_init_walk(t, keys, mid + 1, end);
        return x;
    }
}

void
sbtree_init(sbtree_t *t, int *keys, int n) {
    t->nil.size = 0;
    t->nil.left = &t->nil;
    t->nil.right = &t->nil;
    t->nodes = (node_t *)malloc(sizeof(node_t) * n);
    t->root = sbtree_init_walk(t, keys, 0, n - 1);
}

void
sbtree_free(sbtree_t *t) {
    free(t->nodes);
}

node_t *
sbtree_find_min(node_t *x) {
    while (x->left->size != 0) {
        x = x->left;
    }
    return x;
}

node_t *
sbtree_find_max(node_t *x) {
    while (x->right->size != 0) {
        x = x->right;
    }
    return x;
}

void sbtree_lbalance(node_t **p);
void sbtree_rbalance(node_t **p);
void sbtree_maintain(node_t **p);

void
sbtree_lrotate(node_t **p) {
    node_t *x = *p;
    node_t *y = x->right;
    x->right = y->left;
    y->left = x;
    y->size = x->size;
    x->size = x->left->size + x->right->size + 1;
    *p = y;
}

void
sbtree_rrotate(node_t **p) {
    node_t *x = *p;
    node_t *y = x->left;
    x->left = y->right;
    y->right = x;
    y->size = x->size;
    x->size = x->left->size + x->right->size + 1;
    *p = y;
}

void
sbtree_lbalance(node_t **p) {
    node_t *x = *p;
    if (x->right->size < x->left->left->size) {
        sbtree_rrotate(&x);
    } else if (x->right->size < x->left->right->size) {
        sbtree_lrotate(&x->left);
        sbtree_rrotate(&x);
    } else {
        return;
    }
    sbtree_rbalance(&x->right);
    sbtree_lbalance(&x->left);
    sbtree_maintain(&x);
    *p = x;
}

void
sbtree_rbalance(node_t **p) {
    node_t *x = *p;
    if (x->left->size < x->right->right->size) {
        sbtree_lrotate(&x);
    } else if (x->left->size < x->right->left->size) {
        sbtree_rrotate(&x->right);
        sbtree_lrotate(&x);
    } else {
        return;
    }
    sbtree_lbalance(&x->left);
    sbtree_rbalance(&x->right);
    sbtree_maintain(&x);
    *p = x;
}

void
sbtree_maintain(node_t **p) {
    sbtree_lbalance(p);
    sbtree_rbalance(p);
}

bool
sbtree_delete_walk(sbtree_t *t, node_t **p, int rank, int *pkey) {
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
            node_t *m = sbtree_find_max(x->left);
            x->key = m->key;
            sbtree_delete_walk(t, &x->left, -1, NULL);
        } else if (x->right->size != 0) {
            node_t *m = sbtree_find_min(x->right);
            x->key = m->key;
            sbtree_delete_walk(t, &x->right, 0, NULL);
        }
    } else if (x->left->size < rank) {
        rank -= x->left->size + 1;
        updated = sbtree_delete_walk(t, &x->right, rank, pkey);
    } else {
        updated = sbtree_delete_walk(t, &x->left, rank, pkey);
    }
    if (updated) {
        x->size --;
        if (x->size == 0) {
            *p = &t->nil;
        } else {
            sbtree_maintain(&x);
            *p = x;
        }
    }
    return updated;
}

bool
sbtree_delete(sbtree_t *t, int rank, int *pkey) {
    return sbtree_delete_walk(t, &t->root, rank, pkey);
}

int
sbtree_size(sbtree_t *t) {
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
        int *keys = (int *)malloc(sizeof(int) * n);
        for (int i = 0; i < n; i ++) {
            keys[i] = i + 1;
        }
        sbtree_t __t, *t = &__t;
        sbtree_init(t, keys, n);
        int rank = m - 1;
        int last;
        while (sbtree_size(t) != 0) {
            rank = rank % sbtree_size(t);
            sbtree_delete(t, rank, &last);
            rank = rank + k - 1;
        }
        printf("%d\n", last);
        sbtree_free(t);
        free(keys);
    }
    return 0;
}
