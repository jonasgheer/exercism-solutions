export default class Robot {
  private static usedNames = new Set<string>();

  private _name: string;

  constructor() {
    let generatedName = generateName();
    while (Robot.usedNames.has(generatedName)) {
      generatedName = generateName();
    }
    Robot.usedNames.add(generatedName);
    this._name = generatedName;
  }

  get name(): string {
    return this._name;
  }

  resetName(): void {
    let generatedName = generateName();
    while (Robot.usedNames.has(generatedName)) {
      generatedName = generateName();
    }
    Robot.usedNames.add(generatedName);
    this._name = generatedName;
  }

  static releaseNames(): void {
    Robot.usedNames = new Set();
  }
}

function generateName(): string {
  return (
    randomChar() + randomChar() + randomInt(9) + randomInt(9) + randomInt(9)
  );
}

function randomChar(): string {
  const start = 65;
  return String.fromCharCode(start + randomInt(25));
}

function randomInt(max: number): number {
  return Math.floor(Math.random() * Math.floor(max));
}
