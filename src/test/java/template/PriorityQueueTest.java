package template;

import org.junit.Test;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

public class PriorityQueueTest {

    @Test
    public void testIncreaseOrder() {
        PriorityQueue<Integer> queue = new PriorityQueue<>(5, (i1, i2) -> i1 - i2);
        for (int i = 0; i < 5; i++) {
            assertTrue(queue.push(i));
        }
        for (int i = 0; i < 5; i++) {
            assertFalse(queue.push(0));
        }
        for (int i = 0; i < 5; i++) {
            assertFalse(queue.push(i - 100));
        }
        assertTrue(queue.size() == 5);
        for (int i = 0; i < 5; i++) {
            assertTrue(queue.pop() == i);
        }
        assertTrue(queue.size() == 0);
    }

    @Test
    public void testDecreaseOrder() {
        PriorityQueue<Integer> queue = new PriorityQueue<>(5, (i1, i2) -> i2 - i1);
        for (int i = 0; i < 5; i++) {
            assertTrue(queue.push(i));
        }
        for (int i = 0; i < 5; i++) {
            assertTrue(queue.push(-1));
        }
        assertTrue(queue.size() == 5);
        for (int i = 0; i < 5; i++) {
            assertTrue(queue.pop() == -1);
        }
        assertTrue(queue.size() == 0);
    }

}