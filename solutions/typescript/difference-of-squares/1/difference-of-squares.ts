export default class Square {
  readonly squareOfSum: number;
  readonly sumOfSquares: number;
  readonly difference: number;

  constructor(n: number) {
    const oneToN: number[] = [];
    for (let i = 1; i <= n; i++) {
      oneToN.push(i);
    }
    this.squareOfSum = oneToN.reduce((acc, curr) => acc + curr) ** 2;
    this.sumOfSquares = oneToN.reduce((acc, curr) => acc + curr ** 2);
    this.difference = this.squareOfSum - this.sumOfSquares;
  }
}