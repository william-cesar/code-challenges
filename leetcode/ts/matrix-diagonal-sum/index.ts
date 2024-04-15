function diagonalSum(mat: number[][]): number {
  let rtl = 0;
  let ltr = mat.length - 1;

  let sum = 0;

  mat.forEach((row: number[]) => {
    if (rtl !== ltr) {
      sum += row[rtl] + row[ltr];
    } else {
      sum += row[rtl];
    }

    rtl++;
    ltr--;
  });

  return sum;
}
