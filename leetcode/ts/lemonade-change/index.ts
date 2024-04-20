function lemonadeChange(bills: number[]): boolean {
  const change: Object = {
    5: 0,
    10: 0
  };

  for (let i = 0; i < bills.length; i++) {
    const bill = bills[i];

    if (bill === 5) {
      change[5]++;
    }

    if (bill === 10) {
      if (!change[5]) {
        return false;
      }

      change[10]++;
      change[5]--;
    }

    if (bill === 20) {
      if (change[10] && change[5]) {
        change[10]--;
        change[5]--;

        continue;
      }

      if (change[5] >= 3) {
        change[5] -= 3;

        continue;
      }

      return false;
    }
  }

  return true;
}
