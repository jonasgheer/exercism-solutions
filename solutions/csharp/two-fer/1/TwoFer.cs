using System;

public static class TwoFer
{
    public static string Speak() {
        return "One for you, one for me.";
    }

    public static string Speak(String name) {
        return $"One for {name}, one for me.";
    }
}
