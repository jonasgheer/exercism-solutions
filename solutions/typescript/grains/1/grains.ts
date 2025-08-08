export default class Grains {
  static square(nr: number) {
    if (nr < 1 || nr > 64) throw Error("invalid square");
    return 2 ** (nr - 1);
  }

  static total() {
    let total = 0;
    for (let i = 0; i < 64; i++) {
      total += 2 ** i;
    }
    return total;
  }
}
