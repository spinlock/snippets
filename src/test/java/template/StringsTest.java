package template;

import org.junit.Before;
import org.junit.Test;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

public class StringsTest {

    @Test
    public void atoiTests() {
        Map<String, Integer> tests = new HashMap<>();
        tests.put("", 0);
        tests.put("0", 0);
        tests.put("1", 1);
        tests.put("12345", 12345);
        tests.put("012345", 12345);
        tests.put("12345x", 12345);
        tests.put("18446744073709551616", Integer.MAX_VALUE);
        tests.put("-0", 0);
        tests.put("-1", -1);
        tests.put("12345", 12345);
        tests.put("-12345", -12345);
        tests.put("012345", 12345);
        tests.put("-012345", -12345);
        tests.put("-9223372036854775807", Integer.MIN_VALUE);
        tests.put("2147483649", Integer.MAX_VALUE);
        tests.put("-2147483649", Integer.MIN_VALUE);

        tests.forEach((key, value) -> {
            assertTrue(Strings.atoi(key) == value);
        });
    }

    class StrStrTest {
        String s, needle;
        int expect;

        public StrStrTest(String s, String needle, int expect) {
            this.s = s;
            this.needle = needle;
            this.expect = expect;
        }
    }

    private List<StrStrTest> strstrTests = new LinkedList<>();

    @Before
    public void prepareStrStr() {
        strstrTests.add(new StrStrTest("", "", 0));
        strstrTests.add(new StrStrTest("", "a", -1));
        strstrTests.add(new StrStrTest("", "foo", -1));
        strstrTests.add(new StrStrTest("fo", "foo", -1));
        strstrTests.add(new StrStrTest("foo", "foo", 0));
        strstrTests.add(new StrStrTest("oofofoofooo", "f", 2));
        strstrTests.add(new StrStrTest("oofofoofooo", "foo", 4));
        strstrTests.add(new StrStrTest("barfoobarfoo", "foo", 3));
        strstrTests.add(new StrStrTest("foo", "", 0));
        strstrTests.add(new StrStrTest("foo", "o", 1));
        strstrTests.add(new StrStrTest("abcABCabc", "A", 3));
        strstrTests.add(new StrStrTest("", "a", -1));
        strstrTests.add(new StrStrTest("x", "a", -1));
        strstrTests.add(new StrStrTest("x", "x", 0));
        strstrTests.add(new StrStrTest("abc", "a", 0));
        strstrTests.add(new StrStrTest("abc", "b", 1));
        strstrTests.add(new StrStrTest("abc", "c", 2));
        strstrTests.add(new StrStrTest("abc", "x", -1));
        strstrTests.add(new StrStrTest("abaabbcc", "aa", 2));
    }

    @Test
    public void testStrStr() {
        strstrTests.forEach((t) -> {
            assertTrue(Strings.strstr(t.s, t.needle) == t.expect);
        });
    }

    @Test
    public void testKmpStr() {
        strstrTests.forEach((t) -> {
            assertTrue(Strings.kmpstr(t.s, t.needle) == t.expect);
        });
    }

    @Test
    public void testBase2() {
        for (int i = 0; i < 30; i++) {
            int value = (1 << i) >> 1;
            assertEquals(Strings.convertBase(value, 16), String.format("%x", value));
        }
    }

}