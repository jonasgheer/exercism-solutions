const numeralMap: Record<string, number> = {
    "M": 1000,
    "CM": 900,
    "D": 500,
    "CD": 400,
    "C": 100,
    "XC": 90,
    "L": 50,
    "XL": 40,
    "X": 10,
    "IX": 9,
    "V": 5,
    "IV": 4,
    "I": 1
}

export const toRoman = (n: number): string => {
    let romanNumeral = ""
    for (let [roman, decimal] of Object.entries(numeralMap)) {
        if (n === 0) return romanNumeral;
        const quotient = Math.floor(n / decimal)
        if (quotient === 0) continue;
        romanNumeral += roman.repeat(quotient)
        n -= decimal * quotient
    }
    return romanNumeral
}

