import gleam/int

pub type Error {
  NonPositiveNumber
}

pub fn steps(number: Int) -> Result(Int, Error) {
  steps_with_count(number, 0)
}

fn steps_with_count(number: Int, count: Int) -> Result(Int, Error) {
  case number {
    1 -> Ok(count)
    _ ->  case number > 0 {
            True -> {
              case int.is_even(number) {
                True -> steps_with_count(number/2, count + 1)
                False -> steps_with_count(3*number + 1, count + 1)
              }
            }
            False -> Error(NonPositiveNumber)
          }
  }
}
