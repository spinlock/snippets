package template;

import org.junit.Test;

import java.util.LinkedList;
import java.util.List;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

public class TrieTreeTest {

    @Test
    public void testTrieTree() {
        List<String> ks = new LinkedList<>();
        List<String> xs = new LinkedList<>();
        for (int i = 0; i < 32; i++) {
            ks.add(String.format("%d", i));
            xs.add(String.format("-%d", i));
        }

        TrieTree tree = new TrieTree();

        for (String k : ks) {
            tree.insert(k);
            assertTrue(tree.contains(k));
        }
        for (String k : ks) {
            assertTrue(tree.contains(k));
        }
        for (String x : xs) {
            assertFalse(tree.contains(x));
        }
        assertFalse(tree.contains(""));
        assertTrue(tree.insert(""));
        assertTrue(tree.contains(""));
    }

    @Test
    public void testTrieTreeNoPrefix() {
        TrieTree tree = new TrieTree();
        assertTrue(tree.insertNoPrefix("000"));
        assertTrue(tree.insertNoPrefix("1"));
        assertFalse(tree.insertNoPrefix("11"));
        assertFalse(tree.insertNoPrefix("00"));
        assertFalse(tree.insertNoPrefix("0000"));
    }

}