import java.io.BufferedReader;
import java.io.FileReader;
import java.util.*;

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

    public static void main(String[] args) {
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

