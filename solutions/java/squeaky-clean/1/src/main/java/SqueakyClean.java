import java.util.Map;
import java.util.List;

class SqueakyClean {

    private static final Map<Character, Character> leetSpeak = Map.of(
        '3', 'e',
        '0', 'o',
        '4', 'a',
        '7', 't',
        '1', 'l'
    );
    
    static String clean(String identifier) {
        var strBuilder = new StringBuilder();
        var upCaseNext = false;
        for (var c : identifier.toCharArray()) {
            if (List.of('¡', '!', '#', '$', '.').contains(c)) continue;
            if (upCaseNext) {
                strBuilder.append(Character.toUpperCase(c));
                upCaseNext = false;
                continue;
            }
            if (leetSpeak.containsKey(c)) {
                strBuilder.append(leetSpeak.get(c));
                continue;
            }
            if (c == ' ') {
                strBuilder.append('_');
            } else if (c == '-') {
                upCaseNext = true;
            } else {
                strBuilder.append(c);
            }
        }
        return strBuilder.toString();
    }
}
