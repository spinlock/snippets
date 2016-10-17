package template;

import java.util.ArrayList;
import java.util.Comparator;

public class PriorityQueue<T> {

    private ArrayList<T> heap;
    private int capacity;
    private Comparator<T> comparator;

    public PriorityQueue(int capacity, Comparator<T> comparator) {
        this.capacity = capacity;
        this.heap = new ArrayList<T>(capacity);
        this.comparator = comparator;
    }

    public int size() {
        return heap.size();
    }

    private void swap(int i, int j) {
        T temp = heap.get(i);
        heap.set(i, heap.get(j));
        heap.set(j, temp);
    }

    private boolean less(int i, int j) {
        return lessValue(heap.get(i), heap.get(j));
    }

    private boolean lessValue(T a, T b) {
        return comparator.compare(a, b) < 0;
    }

    private void shiftUp(int i) {
        while (i != 0) {
            int p = (i - 1) / 2;
            if (less(i, p)) {
                swap(i, p);
                i = p;
            } else {
                return;
            }
        }
    }

    private void shiftDown(int i) {
        int n = heap.size();
        while (i < n) {
            int l = (i * 2) + 1;
            int r = (i * 2) + 2;
            int m = i;
            if (l < n && less(l, m)) {
                m = l;
            }
            if (r < n && less(r, m)) {
                m = r;
            }
            if (m != i) {
                swap(i, m);
                i = m;
            } else {
                return;
            }
        }
    }

    public boolean push(T value) {
        if (heap.size() < capacity) {
            heap.add(value);
            shiftUp(heap.size() - 1);
            return true;
        }
        if (capacity != 0 && lessValue(heap.get(0), value)) {
            heap.set(0, value);
            shiftDown(0);
            return true;
        }
        return false;
    }

    public T pop() {
        if (heap.size() != 0) {
            T top = heap.get(0);
            swap(0, heap.size() - 1);
            heap.remove(heap.size() - 1);
            shiftDown(0);
            return top;
        }
        return null;
    }

}
