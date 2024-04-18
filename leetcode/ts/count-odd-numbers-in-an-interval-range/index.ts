function countOdds(low: number, high: number): number {
  const difference: number = high - low;

  if (low & 1 && high & 1) {
    return difference / 2 + 1;
  }

  return difference / 2;
}
