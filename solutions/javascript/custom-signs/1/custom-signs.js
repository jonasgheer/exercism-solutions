export function buildSign(occasion, name) {
  return `Happy ${occasion} ${name}!`;
}

export function buildBirthdaySign(age) {
  const adjective = age >= 50 ? "mature" : "young";
  return `Happy Birthday! What a ${adjective} fellow you are.`;
}

export function graduationFor(name, year) {
  return `Congratulations ${name}!\nClass of ${year}`;
}

export function costOf(sign, currency) {
  return `Your sign costs ${(sign.length * 2 + 20).toFixed(2)} ${currency}.`;
}
