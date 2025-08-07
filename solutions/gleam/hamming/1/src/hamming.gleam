import gleam/string.{split}
import gleam/list.{strict_zip, filter, length}

pub fn distance(strand1: String, strand2: String) -> Result(Int, Nil) {
  case strict_zip(split(strand1, ""), split(strand2, "")) {
    Ok(pairs) -> {
        pairs
        |> filter(fn(pair) {pair.0 != pair.1})
        |> length
        |> Ok
    }
    Error(_) -> Error(Nil)
  }
}
