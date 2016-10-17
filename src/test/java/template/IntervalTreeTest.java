package template;

import org.junit.Test;

import static org.junit.Assert.assertTrue;

public class IntervalTreeTest {

    private int[][] make(int n) {
        int[][] m = new int[n][];
        for (int i = 0; i < n; i++) {
            m[i] = new int[n];
        }
        return m;
    }

    @Test
    public void testIntervalTree() {
        final int n = 100;
        int[] data = new int[n];
        for (int i = 0; i < n; i++) {
            data[i] = i;
        }
        int[][] min = make(n);
        int[][] max = make(n);
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                min[i][j] = Integer.MAX_VALUE;
                max[i][j] = Integer.MIN_VALUE;
                for (int k = i; k <= j; k++) {
                    min[i][j] = Math.min(min[i][j], data[k]);
                    max[i][j] = Math.max(max[i][j], data[k]);
                }
            }
        }
        IntervalTree tree = new IntervalTree(data);
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                assertTrue(min[i][j] == tree.min(i, j));
                assertTrue(max[i][j] == tree.max(i, j));
            }
        }
    }

}