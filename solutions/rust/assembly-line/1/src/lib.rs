pub fn production_rate_per_hour(speed: u8) -> f64 {
    if speed == 0 {
        return 0.0;
    }
    221.0 * speed as f64 * success_rate(speed)
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60.0) as u32
}

fn success_rate(speed: u8) -> f64 {
    match speed {
        1..=4 => 1.0,
        5..=8 => 0.9,
        9 | 10 => 0.77,
        _ => panic!("invalid speed"),
    }
}
