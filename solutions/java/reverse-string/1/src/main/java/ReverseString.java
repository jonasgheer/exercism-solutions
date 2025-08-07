class ReverseString {

    String reverse(String inputString) {
        char[] chars = inputString.toCharArray();
        int front = 0;
        int back = chars.length - 1;
        while (front < chars.length / 2) {
            char temp = chars[front];
            chars[front] = chars[back];
            chars[back] = temp;
            front++;
            back--;
        }
        return String.valueOf(chars);
    }
  
}
