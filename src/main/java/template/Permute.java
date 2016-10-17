package template;

public class Permute {

    private int[] values;

    public Permute(int[] values) {
        this.values = values;
    }

    public Permute(String s) {
        values = new int[s.length()];
        for (int i = 0; i < s.length(); i++) {
            values[i] = s.charAt(i);
        }
    }

    public String toString() {
        StringBuilder sb = new StringBuilder();
        for (Integer v : values) {
            if (sb.length() != 0) {
                sb.append(":");
            }
            sb.append(v);
        }
        return sb.toString();
    }

    private void swap(int i, int j) {
        int t = values[i];
        values[i] = values[j];
        values[j] = t;
    }

    private void reverse(int i, int j) {
        while (i < j) {
            swap(i++, j--);
        }
    }

    public boolean next() {
        int n = values.length;
        if (n <= 1) {
            return false;
        }
        int i = n - 1;
        while (i != 0 && values[i - 1] >= values[i]) {
            i--;
        }
        if (i != 0) {
            int j = n - 1;
            while (values[j] <= values[i - 1]) {
                j--;
            }
            swap(i - 1, j);
        }
        reverse(i, n - 1);
        return i != 0;
    }

    public boolean prev() {
        int n = values.length;
        if (n <= 1) {
            return false;
        }
        int i = n - 1;
        while (i != 0 && values[i - 1] <= values[i]) {
            i--;
        }
        if (i != 0) {
            int j = n - 1;
            while (values[j] >= values[i - 1]) {
                j--;
            }
            swap(i - 1, j);
        }
        reverse(i, n - 1);
        return i != 0;
    }

}
