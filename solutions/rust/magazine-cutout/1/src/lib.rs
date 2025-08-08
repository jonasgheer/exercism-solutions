use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut words = HashMap::new();
    for word in magazine {
        match words.get(word) {
            Some(_) => *words.entry(word).or_insert(0) += 1,
            None => {
                words.insert(word, 1);
            }
        }
    }
    for word in note {
        match words.get(word) {
            Some(count) => {
                if *count == 0 {
                    return false;
                }
                *words.entry(word).or_insert(0) -= 1;
            }
            None => return false,
        }
    }
    true
}
