package template;

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

    private static int[] kmpdfa(String needle) {
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

    public static int kmpstr(String s, String needle) {
        int len1 = strlen(s);
        int len2 = strlen(needle);
        if (len1 < len2) {
            return -1;
        }
        int[] prefix = kmpdfa(needle);
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
                j = prefix[j];
            }
        }
    }

}
