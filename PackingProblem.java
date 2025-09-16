public class PackingProblem {
  // 5 7
  // 5 4
  // 2 3
  // 3 2
  // 4 3
  // 2 1
  public static void main(String[] args) {
    int m = 10; // max weight for suitcase;

    int[] values = { 5, 5, 2, 3, 4, 2 };
    int[] weights = { 7, 4, 3, 2, 3, 1 };

    System.out.println(packing(values, weights, m));
  }

  public static int packing(int[] values, int[] weights, int maxWeight) {
    int[] res = new int[maxWeight + 1];

    for (int i = 0; i < values.length; i++) {
      int w = weights[i];
      int v = values[i];

      for (int j = maxWeight; j >= w; j--) {
        res[j] = Math.max(res[j], v + res[j - w]);
      }
    }

    return res[maxWeight];
  }
}
