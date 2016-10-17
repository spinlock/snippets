package template;

public class Sort {

    private static void xswap(int[] data, int i, int j) {
        int t = data[i];
        data[i] = data[j];
        data[j] = t;
    }

    private static void msort(int[] data, int[] buff, int beg, int end) {
        if (beg >= end) {
            return;
        }
        int mid = beg + (end - beg) / 2;
        msort(data, buff, beg, mid);
        msort(data, buff, mid + 1, end);
        int i = beg, j = mid + 1;
        int k = beg;
        while (i <= mid && j <= end) {
            if (data[i] <= data[j]) {
                buff[k++] = data[i++];
            } else {
                buff[k++] = data[j++];
            }
        }
        while (i <= mid) {
            buff[k++] = data[i++];
        }
        while (j <= end) {
            buff[k++] = data[j++];
        }
        System.arraycopy(buff, beg, data, beg, end - beg + 1);
    }

    public static void msort(int[] data) {
        if (data.length > 1) {
            int[] buff = new int[data.length];
            msort(data, buff, 0, data.length - 1);
        }
    }

    private static void hdown(int[] data, int i, int n) {
        while (i < n) {
            int l = i * 2 + 1;
            int r = i * 2 + 2;
            int m = i;
            if (l < n && data[l] > data[m]) {
                m = l;
            }
            if (r < n && data[r] > data[m]) {
                m = r;
            }
            if (m == i) {
                return;
            }
            xswap(data, i, m);
            i = m;
        }
    }

    public static void hsort(int[] data) {
        if (data.length > 1) {
            int n = data.length;
            for (int i = n - 1; i >= 0; i--) {
                hdown(data, i, n);
            }
            for (int i = n - 1; i != 0; i--) {
                xswap(data, 0, i);
                hdown(data, 0, i);
            }
        }
    }

    private static void qsort(int[] data, int beg, int end) {
        if (beg >= end) {
            return;
        }
        int mid = beg + (end - beg) / 2;
        xswap(data, beg, mid);
        int i = beg;
        for (int j = i + 1; j <= end; j++) {
            if (data[j] < data[beg]) {
                xswap(data, j, ++i);
            }
        }
        xswap(data, beg, i);
        qsort(data, beg, i - 1);
        qsort(data, i + 1, end);
    }

    public static void qsort(int[] data) {
        if (data.length > 1) {
            qsort(data, 0, data.length - 1);
        }
    }

}
