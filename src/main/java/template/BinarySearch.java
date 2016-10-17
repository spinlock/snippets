package template;

public class BinarySearch {

    public static int search(int[] data, int value) {
        return search(data, value, 0, data.length - 1);
    }

    public static int search(int[] data, int value, int beg, int end) {
        while (beg <= end) {
            int mid = beg + (end - beg) / 2;
            if (data[mid] == value) {
                return mid;
            }
            if (value < data[mid]) {
                end = mid - 1;
            } else {
                beg = mid + 1;
            }
        }
        return -(beg + 1);
    }

}
