export default class Bob {
  hey(text: string): string {
    text = text.trim();
    const isYelling = text === text.toUpperCase() && /[A-Z]+/.test(text);
    const isAsking = text.endsWith("?");
    const isSayingNothing = text === "";

    if (isYelling && isAsking) return "Calm down, I know what I'm doing!";
    if (isSayingNothing) return "Fine. Be that way!";
    if (isYelling) return "Whoa, chill out!";
    if (isAsking) return "Sure.";
    return "Whatever.";
  }
}
