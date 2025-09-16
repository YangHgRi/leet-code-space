public class QuickSort {
  public static void main(String[] args) {
    int[] array = { 34, 7, 23, 32, 5, 62 };

    if (array == null || array.length <= 1) {
      return;
    }

    int begin = 0;
    int end = array.length - 1;

    quickSort(array, begin, end);

    for (int e : array) {
      System.err.print(e + " ");
    }
  }

  public static void quickSort(int[] arr, int left, int right) {
    if (left >= right) {
      return;
    }

    int pivot = arr[(left + right) / 2];

    int i = left;
    int j = right;

    while (i <= j) {
      while (arr[i] < pivot) {
        i++;
      }
      while (arr[j] > pivot) {
        j--;
      }

      if (i <= j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
        i++;
        j--;
      }
    }

    if (left < j) {
      quickSort(arr, left, j);
    }
    if (i < right) {
      quickSort(arr, i, right);
    }
  }
}
