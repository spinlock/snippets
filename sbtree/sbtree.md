## Size Balanced Tree
***

### maintain
***

##### rotate
```
LEFT_ROTATE(x)
    y <- x.right
    x.right <- y.left
    y.left <- x
    y.size <- x.size
    x.size <- x.left.size + x.right.size + 1

RIGHT_ROTATE(x)
    y <- x.left
    x.left <- y.right
    y.right <- x
    y.size <- x.size
    x.size <- x.left.size + x.right.size + 1
```

##### balance
```
LEFT_BALANCE(x) // left-child is heavier than right-child
    IF x.right.size < x.left.left.size THEN
        RIGHT_ROTATE(x)
    ELSE IF x.right.size < x.left.right.size THEN
        LEFT_ROTATE(x.left)
        RIGHT_ROTATE(x)
    ELSE
        RETURN
    END
    RIGHT_BALANCE(x.right)
    LEFT_BALANCE(x.left) // tail recursion

RIGHT_BALANCE(x)
    IF x.left.size < x.right.right.size THEN
        LEFT_ROTATE(x)
    ELSE IF x.left.size < x.right.left.size THEN
        RIGHT_TOTATE(x.right)
        LEFT_ROTATE(x)
    ELSE
        RETURN
    END
    LEFT_BALANCE(x.left)
    RIGHT_BALANCE(x.right)
```

##### maintain
```
MAINTAIN(x)
    LEFT_BALANCE(x)
    RIGHT_BALANXE(x)
```

### search, contains, rank, select, predecessor, successor
***

##### search & contains
```
SEARCH(key)
    x <- root
    FOR x.size != 0 DO
        IF x.key == key THEN
            RETURN x
        ELSE IF x.key < key THEN
            x <- x.right
        ELSE
            x <- x.right
        END
    DONE
    RETURN nil

CONTAINS(key)
    RETURN search(key) != nil
```

##### rank, select
```
RANK(key)
    x <- root
    r <- 0
    FOR x.size != 0 DO
        IF x.key == key THEN
            RETURN r + x.left.size
        ELSE IF x.key < key THEN
            r <- r + x.left.size + 1
            x <- x.right
        ELSE
            x <- x.left
        END
    DONE
    RETURN -(r + 1)

SELECT(r)
    x <- root
    IF r < 0 || r >= x.size THEN
        RETURN nil
    END
    FOR TRUE DO
        IF x.left.size == r THEN
            RETURN x
        ELSE IF x.left.size < r THEN
            r <- r - x.left.size - 1
            x <- x.right
        ELSE
            x <- x.left
        END
    DONE
```

##### predecessor, successor
```
PREDECESSOR(key)
    x, pred <- root, nil
    FOR x.size != 0 DO
        IF x.key >= key THEN
            x <- x.left
        ELSE
            x, pred <- x.right, x
        END
    DONE
    RETURN pred

SUCCESSOR(key)
    x, succ <- root, nil
    FOR x.size != 0 DO
        IF x.key <= key THEN
            x <- x.right
        ELSE
            x, succ <- x.left, x
        END
    DONE
    RETURN succ
```

#### insert, remove
***

```
INSERT(key, value)
    RETURN INSERT_RECURSIVE(key, value, root)

INSERT_RECURSIVE(key, value, x)
    IF x.size == 0 THEN
        x <- NewNode(key, value)
        RETURN TRUE, nil
    ELSE IF x.key == key THEN
        x.value, oldValue <- value, x.value
        RETURN FALSE, oldValue
    ELSE IF x.key < key THEN
        addNode, oldValue <- INSERT_RECURSIVE(key, value, x.right)
    ELSE
        addNode, oldValue <- INSERT_RECURSIVE(key, value, x.left)
    END
    IF addNode THEN
        x.size <- x.size + 1
        MAINTAIN(x)
    END
    RETURN addNode, oldValue
```

```
DELETE(key)
    RETURN DELETE_RECURSIVE(key, root)

DELETE_RECURSIVE(key, x)
    IF x.size == 0 THEN
        RETURN FALSE, nil
    ELSE IF x.key == key THEN
        delNode, oldValue <- TRUE, x.value
        IF x.left.size > x.right.size THEN
            p <- FINDMAX(x.left)
            x.key, x.value <- p.key, p.value
            DELETE_RECURSIVE(p.key, x.left)
        ELSE IF x.right.size != 0 THEN
            p <- FINDMIN(x.right)
            x.key, x.value <- p.key, p.value
            DELETE_RECURSIVE(p.key, x.right)
        ELSE
            x.key, x.value <- nil, nil
        END
    ELSE IF x.key < key THEN
        delNode, oldValue <- DELETE_RECURSIVE(key, x.right)
    ELSE
        delNode, oldValue <- DELETE_RECURSIVE(key, x.left)
    END
    IF delNode THEN
        x.size <- x.size - 1
        MAINTAIN(x)
    END
    RETURN delNode, oldValue

FINDMIN(x)
    FOR x.left.size != 0 DO
        x <- x.left
    DONE
    RETURN x

FINDMAX(x)
    FOR x.right.size != 0 DO
        x <- x.right
    DONE
    RETURN x
```
