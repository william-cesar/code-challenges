/* 
  take the first row from the matrix
  then pop the last elements of the other rows
  reverse the rows to get them ordered for the next array.shift 
  reverse the matrix to get the last row in the next iteraction
*/

function spiralOrder(matrix: number[][]): number[] {
  const spiral: Array<number> = [];

  const tRows: number = matrix.length;
  const tColumns: number = matrix[0].length;
  const tItems: number = tRows * tColumns;

  while (spiral.length < tItems) {
    spiral.push(...matrix.shift()!);

    matrix.forEach((row) => spiral.push(row.pop()!));

    matrix.reverse().map((row) => row.reverse());
  }

  return spiral;
}
