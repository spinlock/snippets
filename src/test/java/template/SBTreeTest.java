package template;

import org.junit.Test;

import java.util.ArrayList;

import static org.junit.Assert.assertEquals;

public class SBTreeTest {

    @Test
    public void testAdd() {
        SBTree<Integer> tree = new SBTree<>();
        for (int i = 0; i < 1024; i++) {
            tree.add(i);
            assertEquals(i + 1, tree.size());
        }
        for (int i = 0; i < 1024; i++) {
            tree.add(i);
            assertEquals(1024, tree.size());
        }

        ArrayList<Integer> values = new ArrayList<>(tree.toList());
        assertEquals(1024, values.size());
        for (int i = 0; i < values.size(); i++) {
            assertEquals(i, (int) values.get(i));
        }
    }

    @Test
    public void testRemove() {
        SBTree<Integer> tree = new SBTree<>();
        for (int i = 0; i < 1024; i++) {
            tree.add(i);
            assertEquals(i + 1, tree.size());
        }
        for (int i = 0; i < 1024; i++) {
            tree.remove(i);
            assertEquals(1023 - i, tree.size());
        }
    }

    @Test
    public void testRank() {
        SBTree<Integer> tree = new SBTree<>();
        for (int i = 0; i < 1024; i++) {
            tree.add(i);
        }
        for (int i = 0; i < 1024; i++) {
            assertEquals(i, tree.rank(i));
        }
        assertEquals(-1, tree.rank(-1));
        assertEquals(-1025, tree.rank(1024));
    }

    @Test
    public void testSelect() {
        SBTree<Integer> tree = new SBTree<>();
        for (int i = 0; i < 1024; i++) {
            tree.add(i);
        }
        for (int i = 0; i < 1024; i++) {
            assertEquals(i, (int) tree.select(i));
        }
    }

}