export function translate2d(dx, dy) {
  return (x, y) => [x + dx, y + dy];
}

export function scale2d(sx, sy) {
  return (x, y) => [x * sx, y * sy];
}

export function composeTransform(f, g) {
  return (x, y) => g(...f(x, y));
}

export function memoizeTransform(f) {
  const memory = {x: null, y: null, result: null};
  return (x, y) => {
    if (memory.x === x && memory.y === y) 
      return memory.result;
    const result = f(x, y);
    memory.x = x;
    memory.y = y;
    memory.result = result;
    return result;
  }
}
