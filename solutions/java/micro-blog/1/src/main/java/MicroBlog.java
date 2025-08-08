class MicroBlog {
    public String truncate(String input) {
        try {
            var offset = input.offsetByCodePoints(0, 5);
            return input.substring(0, offset);
        } catch (IndexOutOfBoundsException e) {
            return input;
        }
    }
}
