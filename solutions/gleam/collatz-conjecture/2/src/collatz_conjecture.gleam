import gleam/int.{is_even}

pub type Error {
  NonPositiveNumber
}

pub fn steps(number: Int) -> Result(Int, Error) {
  case number > 0 {
    True -> steps_with_count(number, 0)
    False -> Error(NonPositiveNumber)
  }
}

fn steps_with_count(number: Int, count: Int) -> Result(Int, Error) {
  case number, is_even(number) {
    1, _ -> Ok(count)
    _, True -> steps_with_count(number/2, count + 1)
    _, False -> steps_with_count(3*number + 1, count + 1)
  }
}
