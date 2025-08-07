import java.util.Map;
import static java.util.Map.entry;

class BottleSong {

    private final Map<Integer, String> numToText = Map.ofEntries(
        entry(0, "no"),
        entry(1, "one"),
        entry(2, "two"),
        entry(3, "three"),
        entry(4, "four"),
        entry(5, "five"),
        entry(6, "six"),
        entry(7, "seven"),
        entry(8, "eight"),
        entry(9, "nine"),
        entry(10, "ten")
    );

    String recite(int startBottles, int takeDown) {
        var song = "";
        for (int i = 0; i < takeDown; i++) {
            song += String.format("""
                          %s green %s hanging on the wall,
                          %s green %s hanging on the wall,
                          And if one green bottle should accidentally fall,
                          There'll be %s green %s hanging on the wall.
                          """,
                        capitalize(numToText.get(startBottles)), 
                        bottles(startBottles),
                        capitalize(numToText.get(startBottles)), 
                        bottles(startBottles),
                        numToText.get(startBottles - 1),
                        bottles(startBottles - 1));
            if ((takeDown - 1) != i) song += "\n";
            startBottles--;
        }
        return song;
    }

    private String bottles(int startBottles) {
        return startBottles == 1 ? "bottle" : "bottles";
    }

    private String capitalize(String str) {
        var first = Character.toUpperCase(str.charAt(0));
        var arr = str.toCharArray();
        arr[0] = first;
        return new String(arr);
    }

}