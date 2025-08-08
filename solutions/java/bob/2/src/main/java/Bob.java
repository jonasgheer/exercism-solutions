import java.util.regex.Pattern;
import java.util.function.Predicate;

class Bob {

    private final static Pattern isAlpha = Pattern.compile("[a-zA-Z]");
    private final static Predicate<String> isShout = msg -> isAlpha.matcher(msg).find() && msg == msg.toUpperCase();

    String hey(String input) {
        input = input.trim();
        if (input.isEmpty()) {
            return "Fine. Be that way!";            
        }
        
        var isShouting = isShout.test(input);
        var isQuestioning = input.endsWith("?");
        if (isQuestioning) {
            if (isShouting) {
                return "Calm down, I know what I'm doing!";
            }
            return "Sure.";
        }
        if (isShouting) {
            return "Whoa, chill out!";
        }
        return "Whatever.";
    }
}