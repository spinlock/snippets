package template;

public class Ufs {

    private int[] ufs;

    public Ufs(int n) {
        ufs = new int[n];
        for (int i = 0; i < n; i++) {
            ufs[i] = -1;
        }
    }

    public void union(int i, int j) {
        int pi = find(i);
        int pj = find(j);
        if (pi != pj) {
            ufs[pi] += ufs[pj];
            ufs[pj] = pi;
        }
    }

    public int size(int i) {
        int pi = find(i);
        return -ufs[pi];
    }

    public int find(int i) {
        if (ufs[i] < 0) {
            return i;
        } else {
            ufs[i] = find(ufs[i]);
            return ufs[i];
        }
    }

}
