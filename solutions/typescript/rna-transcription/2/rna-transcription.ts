type Dna = "A" | "C" | "G" | "T";
type Rna = "A" | "C" | "G" | "U";

const dnaMap: Record<Dna, Rna> = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

function isDnaArray(dna: string[]): dna is Dna[] {
  if (dna.every((c) => c in dnaMap)) {
    return true;
  }
  return false;
}

class Transcriptor {
  toRna(dna: string): string {
    const dnaArray = dna.split("");
    if (!isDnaArray(dnaArray)) {
      throw new Error("Invalid input DNA.");
    }
    return dnaArray.map((c) => dnaMap[c]).join("");
  }
}

export default Transcriptor;
