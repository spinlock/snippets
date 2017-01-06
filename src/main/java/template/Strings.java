package template;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

public class Strings {

    private static int strlen(String s) {
        if (s == null) {
            return 0;
        }
        return s.length();
    }

    private static int atoiIgnoreSpaces(String s, int i) {
        while (i < s.length() && s.charAt(i) == ' ') {
            i++;
        }
        return i;
    }

    private static int atoiUpdateValue(int value, boolean positive, int delta) {
        if (positive) {
            final int max = Integer.MAX_VALUE;
            if (value > max / 10 || value * 10 > max - delta) {
                return max;
            }
            return value * 10 + delta;
        } else {
            final int min = Integer.MIN_VALUE;
            if (value < min / 10 || value * 10 < min + delta) {
                return min;
            }
            return value * 10 - delta;
        }
    }

    public static int atoi(String s) {
        int len = strlen(s);
        if (len == 0) {
            return 0;
        }
        int i = atoiIgnoreSpaces(s, 0);
        if (i == len) {
            return 0;
        }

        boolean positive = true;
        switch (s.charAt(i)) {
        case '-':
            positive = false;
        case '+':
            i++;
        }

        int value = 0;
        for (int j = atoiIgnoreSpaces(s, i); j < len; j++) {
            char c = s.charAt(j);
            if (c >= '0' && c <= '9') {
                value = atoiUpdateValue(value, positive, c - '0');
            } else {
                return value;
            }
        }
        return value;
    }

    public static int strstr(String s, String needle) {
        int len1 = strlen(s);
        int len2 = strlen(needle);
        if (len1 < len2) {
            return -1;
        }
        int i = 0, j = 0;
        while (true) {
            if (j == len2) {
                return i;
            }
            if (j + i == len1) {
                return -1;
            }
            if (s.charAt(i + j) == needle.charAt(j)) {
                j++;
            } else {
                i++;
                j = 0;
            }
        }
    }

    // suffix[i] = the length of the longest prefix of needle[0...i-1]
    private static int[] kmpdfa1(String needle) {
        int len = needle.length();
        int[] suffix = new int[len];
        for (int i = 2; i < len; i++) {
            char c = needle.charAt(i - 1);
            int l = suffix[i - 1];
            while (l != 0 && needle.charAt(l) != c) {
                l = suffix[l];
            }
            if (needle.charAt(l) == c) {
                suffix[i] = l + 1;
            }
        }
        return suffix;
    }

    public static int kmpstr1(String s, String needle) {
        int len1 = strlen(s);
        int len2 = strlen(needle);
        if (len1 < len2) {
            return -1;
        }
        int[] suffix = kmpdfa1(needle);
        int i = 0, j = 0;
        while (true) {
            if (j == len2) {
                return i - j;
            }
            if (i == len1) {
                return -1;
            }
            if (s.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
            } else if (j == 0) {
                i++;
            } else {
                j = suffix[j];
            }
        }
    }

    public static String convertBase(int value, int base) {
        final String code = "0123456789abcdef";
        char[] buff = new char[64];
        int len = 0;
        while (value != 0) {
            buff[len++] = code.charAt(value % base);
            value /= base;
        }
        if (len == 0) {
            buff[len++] = '0';
        }
        for (int i = 0, j = len - 1; i < j; i++, j--) {
            char c = buff[i];
            buff[i] = buff[j];
            buff[j] = c;
        }
        return new String(buff, 0, len);
    }

    // prefix[i] = the length of the longest prefix of needle[0...i]
    public static int[] kmpdfa2(String needle) {
        int n = needle.length();
        int[] prefix = new int[n];
        for (int i = 1; i < n; i++) {
            char c = needle.charAt(i);
            int l = prefix[i - 1];
            while (l != 0 && needle.charAt(l) != c) {
                l = prefix[l - 1];
            }
            if (needle.charAt(l) == c) {
                prefix[i] = l + 1;
            }
        }
        return prefix;
    }

    public static int kmpstr2(String s, String needle) {
        int len1 = strlen(s);
        int len2 = strlen(needle);
        if (len1 < len2) {
            return -1;
        }
        int[] prefix = kmpdfa2(needle);
        int i = 0, j = 0;
        while (true) {
            if (j == len2) {
                return i - j;
            }
            if (i == len1) {
                return -1;
            }
            if (s.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
            } else if (j == 0) {
                i++;
            } else {
                j = prefix[j - 1];
            }
        }
    }

    public static List<Integer> kmpstr2All(String s, String needle) {
        int len1 = strlen(s);
        int len2 = strlen(needle);
        if (len1 < len2 || len2 == 0) {
            return Arrays.asList();
        }
        List<Integer> list = new LinkedList<>();
        int[] prefix = kmpdfa2(needle);
        int i = 0, j = 0;
        while (true) {
            if (j == len2) {
                list.add(i - j);
                j = prefix[j - 1];
            }
            if (i == len1) {
                return list;
            }
            if (s.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
            } else if (j == 0) {
                i++;
            } else {
                j = prefix[j - 1];
            }
        }
    }

}
