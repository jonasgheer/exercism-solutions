public class EliudsEggs {
    public int eggCount(int number) {
        // so I guess Integer.bitCount is off limits
        var binaryString = Integer.toBinaryString(number);
        return (int) binaryString.chars().filter(c -> c == '1').count();
    }
}
