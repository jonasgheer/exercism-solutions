import java.util.regex.Pattern;

class Bob {

    String hey(String input) {
        if (input.matches("\s+") || input.isBlank()) {
            return "Fine. Be that way!";
        }
        input = input.replaceAll("[0-9,.'!%^*@#$(*^\s]", "");
        if (input.matches("[A-Z\s]+\\?")) {
            return "Calm down, I know what I'm doing!";
        } else if (input.matches("[A-Z\s]+")) {
            return "Whoa, chill out!";
        } else if (input.endsWith("?")) {
            return "Sure.";
        } else {
            return "Whatever.";
        }
    }
}