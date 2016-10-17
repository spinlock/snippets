package template;

import java.util.LinkedList;

public class TreeVisit {

    class TreeNode {

        TreeNode left, right;

        public void visit() {
        }
    }

    private static void preOrderEnqueue(LinkedList<TreeNode> queue, TreeNode node) {
        if (node != null) {
            queue.addFirst(node);
        }
    }

    public static void preOrderVisit(TreeNode root) {
        LinkedList<TreeNode> queue = new LinkedList();
        preOrderEnqueue(queue, root);
        while (!queue.isEmpty()) {
            TreeNode node = queue.removeFirst();
            node.visit();
            preOrderEnqueue(queue, node.right);
            preOrderEnqueue(queue, node.left);
        }
    }

    private static void inOrderEnqueue(LinkedList<TreeNode> queue, TreeNode node) {
        while (node != null) {
            queue.addFirst(node);
            node = node.left;
        }
    }

    public static void inOrderVisit(TreeNode root) {
        LinkedList<TreeNode> queue = new LinkedList<>();
        inOrderEnqueue(queue, root);
        while (!queue.isEmpty()) {
            TreeNode node = queue.removeFirst();
            node.visit();
            inOrderEnqueue(queue, node.right);
        }
    }

    private static void postOrderEnqueue(LinkedList<TreeNode> queue, TreeNode node) {
        while (node != null) {
            queue.addFirst(node.left);
            node = node.left;
        }
    }

    public static void postOrderVisit(TreeNode root) {
        LinkedList<TreeNode> queue = new LinkedList<>();
        postOrderEnqueue(queue, root);
        TreeNode lastVisit = null;
        while (!queue.isEmpty()) {
            TreeNode node = queue.removeFirst();
            if (lastVisit == node.right) {
                lastVisit = node;
                node.visit();
            } else {
                lastVisit = null;
                queue.addFirst(node);
                postOrderEnqueue(queue, node.right);
            }
        }
    }

    private static void levelOrderEnqueue(LinkedList<TreeNode> queue, TreeNode node) {
        if (node != null) {
            queue.addLast(node);
        }
    }

    public static void levelOrderVisit(TreeNode root) {
        LinkedList<TreeNode> queue = new LinkedList<>();
        levelOrderEnqueue(queue, root);
        while (!queue.isEmpty()) {
            LinkedList<TreeNode> children = new LinkedList<>();
            while (!queue.isEmpty()) {
                TreeNode node = queue.removeFirst();
                node.visit();
                levelOrderEnqueue(children, node.left);
                levelOrderEnqueue(children, node.right);
            }
            queue = children;
        }
    }

}