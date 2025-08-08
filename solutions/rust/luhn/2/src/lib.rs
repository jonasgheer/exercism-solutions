/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let code = code.replace(" ", "");
    if code.len() <= 1 {
        return false;
    }
    let mut total = 0;
    for (i, c) in code.chars().rev().enumerate() {
        if let Some(n) = c.to_digit(10) {
            if i % 2 != 0 {
                let double = n * 2;
                total += if double > 9 { double - 9 } else { double }
            } else {
                total += n;
            }
        } else {
            return false;
        }
    }
    return total % 10 == 0;
}
