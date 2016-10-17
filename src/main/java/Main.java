import java.io.BufferedReader;
import java.io.FileReader;
import java.util.*;

public class Main {

    static class TrieTree {

        class TreeNode {
            Map<Integer, TreeNode> children;
            boolean leaf;
            String name;
            int count;

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

            public boolean hasChildren() {
                return children != null && !children.isEmpty();
            }

        }

        TreeNode root = new TreeNode();
        List<TreeNode> leaves = new LinkedList<>();

        public void insert(String s) {
            TreeNode node = root;
            for (int i = 0; i < s.length(); i++) {
                node = node.get(s.charAt(i), true);
            }
            node.count++;
            if (node.leaf) {
                return;
            }
            node.leaf = true;
            node.name = s;
            leaves.add(node);
        }
    }

    public static void main(String[] args) {
        try {
            BufferedReader in = new BufferedReader(new FileReader("input"));
            TrieTree tree = new TrieTree();
            int total = 0;
            while (true) {
                String name = in.readLine();
                if (name == null) {
                    break;
                }
                tree.insert(name);
                total++;
            }
            List<TrieTree.TreeNode> leaves = tree.leaves;
            Collections.sort(leaves, (node1, node2) -> node1.name.compareTo(node2.name));
            for (TrieTree.TreeNode node : leaves) {
                System.out.println(node.name + " " + ((double) node.count / total));
            }
        } catch (Exception e) {
        }
    }

}

