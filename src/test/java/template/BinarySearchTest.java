package template;

import org.junit.Test;

import static org.junit.Assert.assertTrue;

public class BinarySearchTest {

    @Test
    public void testBinarySearch() {
        assertTrue(BinarySearch.search(new int[]{}, 0) == -1);
        int[] array = new int[10];
        for (int i = 0; i < array.length; i++) {
            array[i] = 2 * i;
        }
        for (int i = 0; i < array.length; i++) {
            assertTrue(BinarySearch.search(array, i * 2) == i);
            assertTrue(BinarySearch.search(array, i * 2 - 1) == -(i + 1));
            assertTrue(BinarySearch.search(array, i * 2 + 1) == -(i + 2));
        }
    }

}