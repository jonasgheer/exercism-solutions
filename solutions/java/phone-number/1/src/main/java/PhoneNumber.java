class PhoneNumber {

    private final String cleanPhoneNumber;

    PhoneNumber(String numberString) {
        numberString = numberString.replaceAll("[-.()\s+]", ""); 
        if (numberString.length() > 11) {
            throw new IllegalArgumentException("must not be greater than 11 digits");
        }
        if (numberString.length() == 11) {
            if (!numberString.startsWith("1")) {
                throw new IllegalArgumentException("11 digits must start with 1");
            }
            numberString = numberString.substring(1);
        }
        if (numberString.length() != 10) {
            throw new IllegalArgumentException("must not be fewer than 10 digits");
        }
        
        if (numberString.chars().filter(Character::isLetter).count() != 0) {
            throw new IllegalArgumentException("letters not permitted");
        }
        if (numberString.chars().filter(Character::isDigit).count() != 10) {
            throw new IllegalArgumentException("punctuations not permitted");
        }
        
        var areaCode = numberString.substring(0, 3);
        var exchangeCode = numberString.substring(3, 6);
        var subscriberNumber = numberString.substring(6, 10);

        if (areaCode.startsWith("0")) {
            throw new IllegalArgumentException("area code cannot start with zero");
        }
        if (areaCode.startsWith("1")) {
            throw new IllegalArgumentException("area code cannot start with one");
        }

        if (exchangeCode.startsWith("0")) {
            throw new IllegalArgumentException("exchange code cannot start with zero");
        }
        if (exchangeCode.startsWith("1")) {
            throw new IllegalArgumentException("exchange code cannot start with one");
        }
        
        cleanPhoneNumber = numberString;
    }

    String getNumber() {
        return cleanPhoneNumber;
    }

}