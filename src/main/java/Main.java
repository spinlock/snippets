import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Main {

    class Ufs {

        int[] ufs;

        public Ufs(int n) {
            ufs = new int[n];
            for (int i = 0; i < n; i++) {
                ufs[i] = -1;
            }
        }

        public int find(int i) {
            if (ufs[i] < 0) {
                return i;
            }
            ufs[i] = find(ufs[i]);
            return ufs[i];
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

    }

    public void solve(InputStream in) throws Exception {
        BufferedReader stdin = new BufferedReader(new InputStreamReader(in));
        while (true) {
            StringTokenizer tokenizer = new StringTokenizer(stdin.readLine());
            int n = Integer.parseInt(tokenizer.nextToken());
            int m = Integer.parseInt(tokenizer.nextToken());

            if (n == 0 && m == 0) {
                return;
            }

            Ufs ufs = new Ufs(n);
            for (int i = 0; i < m; i++) {
                StringTokenizer t = new StringTokenizer(stdin.readLine());
                int k = Integer.parseInt(t.nextToken());
                int l = Integer.parseInt(t.nextToken());
                for (int j = 1; j < k; j++) {
                    int p = Integer.parseInt(t.nextToken());
                    ufs.union(l, p);
                }
            }
            System.out.println(ufs.size(0));
        }
    }

    public static void main(String[] args) {
        try {
            new Main().solve(System.in);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

/*
class Node {

    Node[] nodes;
    boolean tail;

    Node find(byte b) {
        if (nodes == null) {
            return null;
        }
        int i = b & 0xff;
        return nodes[i];
    }

    Node insert(byte b) {
        if (nodes == null) {
            nodes = new Node[256];
        }
        int i = b & 0xff;
        if (nodes[i] == null) {
            nodes[i] = new Node();
        }
        return nodes[i];
    }

}


class Tree {

    Node root = new Node();

    void insert(byte[] keyword) {
        if (keyword.length != 0) {
            Node last = root;
            for (byte b : keyword) {
                last = last.insert(b);
            }
            last.tail = true;
        }
    }

    boolean contains(String s) {
        Node node = root;
        for (byte b : s.getBytes()) {
            if (node == null) {
                return false;
            }
            node = node.find(b);
        }
        return node != null && node.tail;
    }

}

public class Main {

    static class Pair {
        int id, count;

        Pair(int id, int count) {
            this.id = id;
            this.count = count;
        }
    }

    private static int getValue(ArrayList<Integer> array, int i) {
        if (i >= 0 && i < array.size()) {
            return array.get(i);
        }
        return 0;
    }

    private static int test(int n) {
        ArrayList<Integer> values = new ArrayList();
        values.add(1);
        for (int i = 1; i <= n; i++) {
            for (int j = 0; j < values.size(); j++) {
                System.out.print(values.get(j));
                System.out.print(" ");
            }
            System.out.println();
            ArrayList<Integer> updated = new ArrayList();
            for (int j = 0; j <= i; j++) {
                updated.add(getValue(values, j - 1) + getValue(values, j));
            }
            values = updated;
        }
        return 0;
    }

    public static void main(String[] args) {
        test(10);
        try {
            Tree tree = new Tree();
            BufferedReader input = new BufferedReader(new FileReader("/Users/spinlock/Projects/snippets/src/input"));
            String keywords = input.readLine();
            for (String keyword : keywords.split(" ")) {
                if (keyword == null || keyword.length() == 0) {
                    continue;
                }
                tree.insert(keyword.toLowerCase().getBytes());
            }
            Map<Integer, Pair> counts = new HashMap<>();
            int m = Integer.parseInt(input.readLine());
            for (int i = 0; i < m; i++) {
                int id = Integer.parseInt(input.readLine());
                int n = 0;
                Node node = tree.root;
                do {
                    byte b = (byte) input.read();
                    if (b >= 'A' && b <= 'Z') {
                        b += 'a' - 'A';
                    }
                    if (b >= 'a' && b <= 'z') {
                        if (node != null) {
                            node = node.find(b);
                        }
                    } else {
                        if (node != null && node.tail) {
                            n++;
                        }
                        node = tree.root;
                    }
                    if (b == '\n') {
                        break;
                    }
                } while (true);
                if (counts.containsKey(id)) {
                    counts.get(id).count += n;
                } else {
                    counts.put(id, new Pair(id, n));
                }
            }
            List<Pair> pairs = new ArrayList<>(counts.values());
            Collections.sort(pairs, new Comparator<Pair>() {
                public int compare(Pair o1, Pair o2) {
                    if (o1.count > o2.count) {
                        return -1;
                    }
                    if (o1.count < o2.count) {
                        return 1;
                    }
                    return o1.id - o2.id;
                }
            });
            for (int i = 0; i < pairs.size(); i++) {
                if (i != 0) {
                    System.out.print(" ");
                }
                Pair p = pairs.get(i);
                System.out.print(p.id);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}

*/


