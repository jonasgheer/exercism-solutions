use std::num::ParseIntError;

/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let code = code.replace(" ", "");
    let code: Vec<&str> = code.split("").filter(|c| *c != "").collect();
    if code.len() <= 1 {
        return false;
    }
    let code: Result<Vec<u32>, ParseIntError> = code.iter().map(|c| c.parse::<u32>()).collect();
    let mut code = match code {
        Ok(numbers) => numbers,
        Err(_) => return false,
    };
    for i in (0..code.len() - 1).rev().step_by(2) {
        let double = if code[i] * 2 > 9 {
            code[i] * 2 - 9
        } else {
            code[i] * 2
        };
        code[i] = double;
    }

    return code.iter().sum::<u32>() % 10 == 0;
}
