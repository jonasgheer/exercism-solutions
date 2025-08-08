public class CarsAssemble {

    private final int PRODUCTION_SPEED = 221;

    public double productionRatePerHour(int speed) {
        return PRODUCTION_SPEED * speed * successRate(speed);
    }

    public int workingItemsPerMinute(int speed) {
        return (int) (productionRatePerHour(speed) / 60.0);
    }

    private double successRate(int speed) {
        return switch(speed) {
            case 0 -> 0.0;
            case 1,2,3,4 -> 1.0;
            case 5,6,7,8 -> 0.9;
            case 9 -> 0.8;
            case 10 -> 0.77;
            default -> throw new IllegalArgumentException();
        };
    }
}
