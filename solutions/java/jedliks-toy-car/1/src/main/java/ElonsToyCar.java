import java.lang.System;

public class ElonsToyCar {

    private int battery = 100;
    
    public static ElonsToyCar buy() {
        return new ElonsToyCar();
    }

    public String distanceDisplay() {
        return String.format("Driven %s meters", (100 - battery) * 20);
    }

    public String batteryDisplay() {
        if (battery == 0) {
            return "Battery empty";
        } 
        return String.format("Battery at %s%%", battery);
    }

    public void drive() {
        if (battery == 0) return;
        battery--;
    }
}
