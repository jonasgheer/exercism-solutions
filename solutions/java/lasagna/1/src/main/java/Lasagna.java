public class Lasagna {
    public int expectedMinutesInOven() {
        return 40;
    }

    public int remainingMinutesInOven(int minutes) {
        return 40 - minutes;
    }

    public int preparationTimeInMinutes(int numLayers) {
        return numLayers * 2;
    }

    public int totalTimeInMinutes(int numLayers, int timeInOven) {
        return (2 * numLayers) + timeInOven;
    }
}
