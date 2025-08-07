export default class Matrix {
  readonly columns: number[][];
  readonly rows: number[][];

  constructor(numbers: string) {
    const rows = numbers
      .split("\n")
      .map((row) => row.split(" ").map((n) => Number(n)));

    let columns: number[][] = [];
    rows[0].forEach(() => {
      columns.push([]);
    });
    for (let i = 0; i < rows.length; i++) {
      for (let j = 0; j < rows[0].length; j++) {
        columns[j].push(rows[i][j]);
      }
    }
    this.rows = rows;
    this.columns = columns;
  }
}
