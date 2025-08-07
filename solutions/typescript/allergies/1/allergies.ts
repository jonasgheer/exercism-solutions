const ItemScore = {
  eggs: 1,
  peanuts: 2,
  shellfish: 4,
  strawberries: 8,
  tomatoes: 16,
  chocolate: 32,
  pollen: 64,
  cats: 128,
} as const;

type Item = keyof typeof ItemScore;

export default class Allergies {
  private allergies: Item[] = [];

  constructor(score: number) {
    if (score > 256) {
      score = score % 256;
    }

    for (const [item, s] of Object.entries(ItemScore).sort(
      ([, s1], [, s2]) => s2 - s1
    )) {
      if (score - s >= 0) {
        this.allergies.push(item as Item);
        score -= s;
      }
    }
  }

  list(): Item[] {
    return this.allergies.reverse();
  }

  allergicTo(item: Item): boolean {
    return this.allergies.includes(item);
  }
}
