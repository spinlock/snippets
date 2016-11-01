package template;

import java.util.LinkedList;
import java.util.List;

public class SBTree<K extends Comparable> {

    final private Node nil = new Node();

    class Node {

        private K key;
        private int size;
        private Node left, right;

        public Node() {
            left = right = this;
        }

        public Node(K key) {
            left = right = nil;
            size = 1;
            this.key = key;
        }

        private Node leftRotate() {
            Node root = right;
            right = root.left;
            root.left = this;
            root.size = size;
            size = left.size + right.size + 1;
            return root;
        }

        private Node rightRotate() {
            Node root = left;
            left = root.right;
            root.right = this;
            root.size = size;
            size = left.size + right.size + 1;
            return root;
        }

        private Node leftBalance() {
            if (right.size < left.left.size) {
                return rightRotate().maintain();
            }
            if (right.size < left.right.size) {
                left = left.leftRotate();
                return rightRotate().maintain();
            }
            return this;
        }

        private Node rightBalance() {
            if (left.size < right.right.size) {
                return leftRotate().maintain();
            }
            if (left.size < left.right.size) {
                right = right.rightRotate();
                return leftRotate().maintain();
            }
            return this;
        }

        private Node maintain() {
            left = left.leftBalance();
            right = right.rightBalance();
            return balance();
        }

        private Node balance() {
            return leftBalance().rightBalance();
        }

        public int size() {
            return size;
        }

        public Node insert(K key) {
            if (size == 0) {
                return new Node(key);
            }
            int d = this.key.compareTo(key);
            if (d == 0) {
                return this;
            } else if (d < 0) {
                right = right.insert(key);
            } else {
                left = left.insert(key);
            }
            size = left.size + right.size + 1;
            return balance();
        }

        public Node remove(K key) {
            if (size == 0) {
                return nil;
            }
            int d = this.key.compareTo(key);
            if (d == 0) {
                if (left.size > right.size) {
                    this.key = findMax(left);
                    left = left.remove(this.key);
                } else if (right.size != 0) {
                    this.key = findMin(right);
                    right = right.remove(this.key);
                } else {
                    return nil;
                }
            } else if (d < 0) {
                right = right.remove(key);
            } else {
                left = left.remove(key);
            }
            size = left.size + right.size + 1;
            return balance();
        }

        private K findMax(Node node) {
            while (node.right.size != 0) {
                node = node.right;
            }
            return node.key;
        }

        private K findMin(Node node) {
            while (node.left.size != 0) {
                node = node.left;
            }
            return node.key;
        }

        public int rank(K key) {
            Node node = this;
            int index = 0;
            while (node.size != 0) {
                int d = node.key.compareTo(key);
                if (d == 0) {
                    return index + node.left.size;
                } else if (d < 0) {
                    index += node.left.size + 1;
                    node = node.right;
                } else {
                    node = node.left;
                }
            }
            return -(index + 1);
        }

        public K select(int index) {
            if (index < 0) {
                index += size;
            }
            Node node = this;
            while (node.size != 0) {
                if (index == node.left.size) {
                    return node.key;
                } else if (index < node.left.size) {
                    node = node.left;
                } else {
                    index -= node.left.size + 1;
                    node = node.right;
                }
            }
            return null;
        }

        public List<K> toList(List<K> keys) {
            if (size != 0) {
                left.toList(keys);
                keys.add(key);
                right.toList(keys);
            }
            return keys;
        }

    }

    private Node root = nil;

    public int size() {
        return root.size();
    }

    public void add(K key) {
        root = root.insert(key);
    }

    public void remove(K key) {
        root = root.remove(key);
    }

    public int rank(K key) {
        return root.rank(key);
    }

    public K select(int index) {
        return root.select(index);
    }

    public List<K> toList() {
        return root.toList(new LinkedList<>());
    }

}
