package template;

import java.util.HashMap;
import java.util.Map;

public class TrieTree {

    class TreeNode {
        Map<Integer, TreeNode> children;
        boolean leaf;

        public boolean hasChildren() {
            return children != null && !children.isEmpty();
        }

        public TreeNode get(int c, boolean create) {
            if (children != null) {
                TreeNode node = children.get(c);
                if (node != null) {
                    return node;
                }
            }
            if (create) {
                if (children == null) {
                    children = new HashMap<>();
                }
                TreeNode node = new TreeNode();
                children.put(c, node);
                return node;
            }
            return null;
        }
    }

    private TreeNode root = new TreeNode();

    public boolean insert(String s) {
        TreeNode node = root;
        for (int i = 0; i < s.length(); i++) {
            node = node.get(s.charAt(i), true);
        }
        if (node.leaf) {
            return false;
        }
        node.leaf = true;
        return true;
    }

    public boolean insertNoPrefix(String s) {
        TreeNode node = root;
        for (int i = 0; i < s.length(); i++) {
            if (node.leaf) {
                return false;
            }
            node = node.get(s.charAt(i), true);
        }
        if (node.leaf || node.hasChildren()) {
            return false;
        }
        node.leaf = true;
        return true;
    }

    public boolean contains(String s) {
        TreeNode node = root;
        for (int i = 0; i < s.length(); i++) {
            if (node == null) {
                return false;
            }
            node = node.get(s.charAt(i), false);
        }
        return node != null && node.leaf;
    }

}
