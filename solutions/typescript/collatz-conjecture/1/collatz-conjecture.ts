class CollatzConjecture {
  static steps(n: number): number {
    if (n < 1) {
      throw Error("Only positive numbers are allowed");
    }
    let steps = 0;
    while (n !== 1) {
      if (n % 2 === 0) {
        n /= 2;
      } else {
        n = 3 * n + 1;
      }
      steps++;
    }
    return steps;
  }
}

export default CollatzConjecture;
