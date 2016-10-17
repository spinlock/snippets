package template;

import org.junit.Test;

import java.util.HashSet;
import java.util.Set;

import static org.junit.Assert.assertTrue;

public class PermuteTest {

    private void testNext(String input, int n) {
        Permute p = new Permute(input);
        Set<String> set = new HashSet<>();
        while (true) {
            String s = p.toString();
            if (set.contains(s)) {
                break;
            }
            set.add(s);
            p.next();
        }
        assertTrue(set.size() == n);
    }

    @Test
    public void testNextPermutation() {
        testNext("a", 1);
        testNext("abcdefgh", 40320);
        testNext("aaa", 1);
        testNext("aaabbbcccddd", 369600);
    }

    private void testPrev(String input, int n) {
        Permute p = new Permute(input);
        Set<String> set = new HashSet<>();
        while (true) {
            String s = p.toString();
            if (set.contains(s)) {
                break;
            }
            set.add(s);
            p.prev();
        }
        assertTrue(set.size() == n);
    }

    @Test
    public void testPrevPermutation() {
        testPrev("a", 1);
        testPrev("abcdefgh", 40320);
        testPrev("aaa", 1);
        testPrev("aaabbbcccddd", 369600);
    }

}