export function isIsogram(word: string): boolean {
    word = word.toLowerCase().replaceAll(/[- ]/g, "")
    return word.length === new Set(word).size
}
