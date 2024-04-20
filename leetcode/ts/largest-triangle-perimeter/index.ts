/*
 To be a valid triangle it has to satisfy either the following conditions:
 a). sum of any two sides greater than third side
 b). difference of any two sides less than third side
*/

function largestPerimeter(nums: number[]): number {
  nums.sort((a, b) => a - b);

  let maxPerimeter: number = 0;
  const maxIteractions: number = nums.length - 3;

  for (let i = 0; i <= maxIteractions; i++) {
    const a: number = nums[i];
    const b: number = nums[i + 1];
    const c: number = nums[i + 2];

    const isTriangle: boolean = a + b > c;

    if (isTriangle) {
      maxPerimeter = a + b + c;
    }
  }

  return maxPerimeter;
}
