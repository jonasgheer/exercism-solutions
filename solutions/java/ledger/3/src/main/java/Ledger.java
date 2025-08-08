import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Stream;

import static java.util.Locale.forLanguageTag;

public class Ledger {
    public LedgerEntry createLedgerEntry(String date, String desc, int change) {
        return new LedgerEntry(LocalDate.parse(date), desc, change);
    }

    public String format(String currency, String locale, LedgerEntry[] entries) {
        CurrencyLocale currencyLocale = CurrencyLocale.from(currency, locale);
        if (entries.length == 0) {
            return currencyLocale.locale.header;
        }
        return format(currencyLocale, entries);
    }


    private String format(CurrencyLocale currencyLocale, LedgerEntry[] entries) {
        var negativeEntries = new ArrayList<LedgerEntry>();
        var positiveEntries = new ArrayList<LedgerEntry>();
        for (LedgerEntry entry : entries) {
            if (entry.change() >= 0) {
                positiveEntries.add(entry);
            } else {
                negativeEntries.add(entry);
            }
        }
        negativeEntries.sort(Comparator.comparing(LedgerEntry::localDate));
        positiveEntries.sort(Comparator.comparing(LedgerEntry::localDate));

        var allEntries = Stream.concat(negativeEntries.stream(), positiveEntries.stream()).toList();
        var result = new StringBuilder(currencyLocale.locale.header);
        for (LedgerEntry e : allEntries) {
            String date = e.localDate().format(currencyLocale.locale.dateTimeFormatter);
            String desc = e.description().length() > 25 ? e.description().substring(0, 22) + "..." : e.description();
            String amount = currencyLocale.formatAmount(e.change());

            result.append("\n");
            result.append(String.format("%s | %-25s | %13s",
                    date,
                    desc,
                    amount));
        }
        return result.toString();
    }

    public record LedgerEntry(LocalDate localDate, String description, double change) {
    }

    private enum CurrencyLocale {
        EN_US_USD(Locale.EN_US, Currency.USD),
        EN_US_EUR(Locale.EN_US, Currency.EUR),
        NL_NL_USD(Locale.NL_NL, Currency.USD),
        NL_NL_EUR(Locale.NL_NL, Currency.EUR);

        private final Locale locale;
        private final Currency currency;

        CurrencyLocale(Locale locale, Currency currency) {
            this.locale = locale;
            this.currency = currency;
        }

        private String formatAmount(double change) {
            String amount = this.locale.formatAmount(change);
            if (this.locale.equals(CurrencyLocale.Locale.EN_US)) {
                if (change < 0)
                    amount = "(" + this.currency.currencySymbol + amount.substring(1);
                else
                    amount = this.currency.currencySymbol + amount + " ";
            } else {
                amount = this.currency.currencySymbol + " " + amount + " ";
            }
            return amount;
        }

        static CurrencyLocale from(String currency, String locale) {
            for (var lc : CurrencyLocale.values()) {
                if (lc.currency.name().equals(currency) && lc.locale.localeStr.equals(locale)) {
                    return lc;
                }
            }
            throw new IllegalArgumentException(locale + "+" + currency + "is not a valid combination of currency and locale");
        }

        private enum Locale {
            EN_US("en-US", DateTimeFormatter.ofPattern("MM/dd/yyyy"), "%(,.2f", "Date       | Description               | Change       "),
            NL_NL("nl-NL", DateTimeFormatter.ofPattern("dd/MM/yyyy"), "%,.2f", "Datum      | Omschrijving              | Verandering  ");

            private final String localeStr;
            private final java.util.Locale locale;
            private final DateTimeFormatter dateTimeFormatter;
            private final String formattingFlags;
            private final String header;

            Locale(String localeStr, DateTimeFormatter dateTimeFormatter, String formattingFlags, String header) {
                this.localeStr = localeStr;
                this.locale = forLanguageTag(localeStr);
                this.dateTimeFormatter = dateTimeFormatter;
                this.formattingFlags = formattingFlags;
                this.header = header;
            }

            String formatAmount(double amount) {
                return String.format(locale, this.formattingFlags, amount / 100);
            }
        }

        private enum Currency {
            USD("$"), EUR("€");

            private final String currencySymbol;

            Currency(String currencySymbol) {
                this.currencySymbol = currencySymbol;
            }
        }
    }
}
