import java.util.Map;
import java.util.HashMap;

class ParallelLetterFrequency {

    private String[] texts;

    ParallelLetterFrequency(String[] texts) {
        this.texts = texts;
    }

    Map<Character, Integer> countLetters() {
        var map = new HashMap<Character, Integer>();
        for (var line : this.texts) {
            line = line.toLowerCase().replaceAll("[-!?;:,.0-9\"\'\n()\s]", "").trim();
            for (var c : line.toCharArray()) {
                if (map.containsKey(c)) {
                    map.put(c, (map.get(c) + 1));
                } else {
                    map.put(c, 1);
                }
            }
        }
        return map;
    }

}
