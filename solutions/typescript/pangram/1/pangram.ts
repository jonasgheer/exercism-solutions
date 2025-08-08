export default class Pangram {
  private static alphabet = "abcdefghijklmnopqrstuvwxyz".split("");

  constructor(readonly sentence: string) {}

  isPangram() {
    return Pangram.alphabet.every((c) =>
      this.sentence.toLowerCase().includes(c)
    );
  }
}
