function solution(inputString: string): boolean {
  return inputString === [...inputString].reverse().join('');
}
