const colorMap = {
  black: "0",
  brown: "1",
  red: "2",
  orange: "3",
  yellow: "4",
  green: "5",
  blue: "6",
  violet: "7",
  grey: "8",
  white: "9",
} as const;

type Color = keyof typeof colorMap;

export class ResistorColor {
  private colors: Array<Color>;

  constructor(colors: Color[]) {
    if (colors.length < 2) {
      throw new Error("At least two colors need to be present");
    }
    this.colors = colors;
  }

  value = (): number =>
    Number(colorMap[this.colors[0]] + colorMap[this.colors[1]]);
}
