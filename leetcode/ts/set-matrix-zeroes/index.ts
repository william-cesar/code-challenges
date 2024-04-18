/**
 Do not return anything, modify matrix in-place instead.
 */
function setZeroes(matrix: number[][]): void {
  let isFirstRowZero: boolean = false;
  let isFirstColZero: boolean = false;

  const m: number = matrix.length;
  const n: number = matrix[0].length;

  //  check if first row has zero
  for (let i = 0; i < n; i++) {
    if (matrix[0][i] === 0) {
      isFirstRowZero = true;
      break;
    }
  }

  //  check if first col has zero
  for (let i = 0; i < m; i++) {
    if (matrix[i][0] === 0) {
      isFirstColZero = true;
      break;
    }
  }

  // check which col and row has zero in whole matrix
  for (let i = 1; i < m; i++) {
    for (let j = 1; j < n; j++) {
      if (matrix[i][j] === 0) {
        matrix[i][0] = 0;
        matrix[0][j] = 0;
      }
    }
  }

  //  set matrix element to zero if starting element ion row and col has a zero
  for (let i = 1; i < m; i++) {
    for (let j = 1; j < n; j++) {
      if (matrix[i][0] === 0 || matrix[0][j] === 0) {
        matrix[i][j] = 0;
      }
    }
  }

  //   if first row has zero then make all ele at first row zero
  if (isFirstRowZero) {
    for (let i = 0; i < n; i++) {
      matrix[0][i] = 0;
    }
  }

  //   if first col has zero then make all ele at first col zero
  if (isFirstColZero) {
    for (let i = 0; i < m; i++) {
      matrix[i][0] = 0;
    }
  }
}
