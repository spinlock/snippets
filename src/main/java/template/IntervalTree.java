package template;

public class IntervalTree {

    class TreeNode {
        int beg, end;
        int min, max;

        TreeNode left, right;

        TreeNode(int[] data, int beg, int end) {
            this.beg = beg;
            this.end = end;
            if (beg == end) {
                this.min = data[beg];
                this.max = data[beg];
            } else {
                int mid = beg + (end - beg) / 2;
                left = new TreeNode(data, beg, mid);
                right = new TreeNode(data, mid + 1, end);
                this.min = Math.min(left.min, right.min);
                this.max = Math.max(left.max, right.max);
            }
        }

        public int min(int beg, int end) {
            beg = Math.max(beg, this.beg);
            end = Math.min(end, this.end);
            if (beg > end) {
                return Integer.MAX_VALUE;
            }
            if (beg == this.beg && end == this.end) {
                return this.min;
            }
            int value = Integer.MAX_VALUE;
            if (beg <= left.end) {
                value = Math.min(value, left.min(beg, end));
            }
            if (end >= right.beg) {
                value = Math.min(value, right.min(beg, end));
            }
            return value;
        }

        public int max(int beg, int end) {
            beg = Math.max(beg, this.beg);
            end = Math.min(end, this.end);
            if (beg > end) {
                return Integer.MIN_VALUE;
            }
            if (beg == this.beg && end == this.end) {
                return this.max;
            }
            int value = Integer.MIN_VALUE;
            if (beg <= left.end) {
                value = Math.max(value, left.max(beg, end));
            }
            if (end >= right.beg) {
                value = Math.max(value, right.max(beg, end));
            }
            return value;
        }
    }

    private TreeNode root;

    public IntervalTree(int[] data) {
        if (data.length != 0) {
            root = new TreeNode(data, 0, data.length - 1);
        }
    }

    public int min(int beg, int end) {
        if (root != null) {
            return root.min(beg, end);
        }
        return Integer.MAX_VALUE;
    }

    public int max(int beg, int end) {
        if (root != null) {
            return root.max(beg, end);
        }
        return Integer.MIN_VALUE;
    }

}
