import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Locale;

public class Ledger {
    public LedgerEntry createLedgerEntry(String date, String desc, int change) {
        return new LedgerEntry(LocalDate.parse(date), desc, change);
    }

    private static Map<String, String> headers = Map.of(
        "en-US", "Date       | Description               | Change       ",
        "nl-NL", "Datum      | Omschrijving              | Verandering  "
    );

    public String format(String currency, String locale, LedgerEntry[] entries) {
        if (!currency.equals("USD") && !currency.equals("EUR")) {
            throw new IllegalArgumentException("Invalid currency");
        }
        if (!locale.equals("en-US") && !locale.equals("nl-NL")) {
            throw new IllegalArgumentException("Invalid locale");
        }
        if (entries.length <= 0) {
            return headers.get(locale);
        }
    
        var negative = new ArrayList<LedgerEntry>();
        var positive = new ArrayList<LedgerEntry>();
        for (LedgerEntry entry : entries) {
            if (entry.change() >= 0) {
                positive.add(entry);
            } else {
                negative.add(entry);
            }
        }
        negative.sort((o1, o2) -> o1.localDate().compareTo(o2.localDate()));
        positive.sort((o1, o2) -> o1.localDate().compareTo(o2.localDate()));

        List<LedgerEntry> all = new ArrayList<>();
        all.addAll(negative);
        all.addAll(positive);

        var datetimePattern = locale.equals("en-US") ? "MM/dd/yyyy" : "dd/MM/yyyy";
        var str = headers.get(locale);
        for (LedgerEntry e : all) {
            String date = e.localDate().format(DateTimeFormatter.ofPattern(datetimePattern));
            String desc = e.description().length() > 25 ? e.description().substring(0, 22) + "..." : e.description();
            String amount = formatAmount(e.change(), currency, locale);

            str += "\n";
            str += String.format("%s | %-25s | %13s",
                                  date,
                                  desc,
                                  amount);
        }
        return str;
    }

    private String formatAmount(double change, String currency, String locale) {
        var currencySymbol = currency.equals("USD") ? "$" : "€";
        
        String amount = String.format(Locale.forLanguageTag(locale), locale.equals("en-US") ? "%(,.2f" : "%,.2f", change / 100);
        if (locale.equals("en-US")) {
            if (change < 0)
                amount = "(" + currencySymbol + amount.substring(1);
            else
                amount = currencySymbol + amount + " ";
        } else {
            amount = currencySymbol + " " + amount + " ";
        }

        return amount;
    }

    public static record LedgerEntry(LocalDate localDate, String description, double change) {}
}
