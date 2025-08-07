import java.time.LocalDate;
import java.time.LocalTime;
import java.time.Year;
import java.time.Month;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

class AppointmentScheduler {

    private final static DateTimeFormatter parser = DateTimeFormatter.ofPattern("MM/dd/yyyy HH:mm:ss");
    
    public LocalDateTime schedule(String appointmentDateDescription) {
        return LocalDateTime.parse(appointmentDateDescription, parser);
    }

    public boolean hasPassed(LocalDateTime appointmentDate) {
        return LocalDateTime.now().isAfter(appointmentDate);
    }

    public boolean isAfternoonAppointment(LocalDateTime appointmentDate) {
        var appointment = appointmentDate.toLocalTime();
        return appointment.compareTo(LocalTime.NOON) > -1 && appointment.compareTo(LocalTime.of(18, 0)) < 0;
    }

    public String getDescription(LocalDateTime appointmentDate) {
        return appointmentDate.format(DateTimeFormatter.ofPattern("'You have an appointment on' EEEE, LLLL d, y, 'at' h:mm a."));
    }

    public LocalDate getAnniversaryDate() {
        return LocalDate.of(Year.now().getValue(), Month.SEPTEMBER, 15);
    }
}
