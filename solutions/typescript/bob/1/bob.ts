export default class Bob {
  private isYelling = (text: string) =>
    text === text.toUpperCase() && /[A-Z]{1,}/.test(text);
  private isAsking = (text: string) => /[?]$/.test(text);
  private isSayingNothing = (text: string) => text === "";

  hey(text: string): string {
    text = text.trim();
    if (this.isYelling(text) && this.isAsking(text)) {
      return "Calm down, I know what I'm doing!";
    } else if (this.isSayingNothing(text)) {
      return "Fine. Be that way!";
    } else if (this.isYelling(text)) {
      return "Whoa, chill out!";
    } else if (this.isAsking(text)) {
      return "Sure.";
    } else {
      return "Whatever.";
    }
  }
}
