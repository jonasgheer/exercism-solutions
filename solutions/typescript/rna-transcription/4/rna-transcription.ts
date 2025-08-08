type Dna = "A" | "C" | "G" | "T";
type Rna = "A" | "C" | "G" | "U";

const dnaMap: Record<Dna, Rna> = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

function isDna(dna: string): dna is Dna {
  if (!(dna in dnaMap)) {
    throw Error("Invalid input DNA.");
  }
  return true;
}

class Transcriptor {
  toRna(dna: string): string {
    return dna
      .split("")
      .map((c) => isDna(c) && dnaMap[c])
      .join("");
  }
}

export default Transcriptor;
