function maximumWealth(accounts: number[][]): number {
  let richest: number = 0;

  accounts.forEach((account: number[]) => {
    let sum: number = 0;

    account.forEach((value: number) => (sum += value));

    if (sum > richest) {
      richest = sum;
    }
  });

  return richest;
}
