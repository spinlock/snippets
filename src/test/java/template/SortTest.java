package template;

import org.junit.Test;

import java.util.Arrays;
import java.util.Random;

import static org.junit.Assert.assertTrue;

public class SortTest {

    private int[] duplicate(int[] array) {
        int[] dup = new int[array.length];
        System.arraycopy(array, 0, dup, 0, array.length);
        return dup;
    }

    private int[] sorted(int[] array) {
        int[] dup = duplicate(array);
        Arrays.sort(dup);
        return dup;
    }

    private int[] randArray(int n) {
        Random r = new Random();
        int[] array = new int[n];
        for (int i = 0; i < n; i++) {
            array[i] = r.nextInt(10000);
        }
        return array;
    }

    private void verify(int[] array, int[] expect) {
        for (int i = 0; i < array.length; i++) {
            assertTrue(array[i] == expect[i]);
        }
    }

    @Test
    public void testMergeSort() {
        for (int i = 0; i < 10; i++) {
            int[] array = randArray(1000);
            int[] expect = sorted(array);
            Sort.msort(array);
            verify(array, expect);
        }
    }

    @Test
    public void testQuickSort() {
        for (int i = 0; i < 10; i++) {
            int[] array = randArray(1000);
            int[] expect = sorted(array);
            Sort.qsort(array);
            verify(array, expect);
        }
    }

    @Test
    public void testHeapSort() {
        for (int i = 0; i < 10; i++) {
            int[] array = randArray(1000);
            int[] expect = sorted(array);
            Sort.hsort(array);
            verify(array, expect);
        }
    }

}