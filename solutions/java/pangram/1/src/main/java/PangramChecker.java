import java.util.Map;
import java.util.HashMap;

public class PangramChecker {

    private final Map<Character, Boolean> alphabet = new HashMap<>();

    PangramChecker() {
        for (char c : "abcdefghijklmnopqrstuvwxyz".toCharArray()) {
            alphabet.put(c, false);
        }
    }

    public boolean isPangram(String input) {
        input = input.toLowerCase();
        for (char c : input.toCharArray()) {
            alphabet.put(c, true);
        }
        for (var value : alphabet.values()) {
            if (value == false) return false;
        }
        return true;
    }

}
