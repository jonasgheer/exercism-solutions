enum Color {
  black,
  brown,
  red,
  orange,
  yellow,
  green,
  blue,
  violet,
  grey,
  white,
}

export const colorCode = (color: keyof typeof Color): number => Color[color];

export const COLORS = Object.keys(Color).filter((color) =>
  isNaN(parseInt(color))
);
