module Accumulate

let accumulate (func: 'a -> 'b) (input: 'a list): 'b list =
    let rec map elements agg =
        match elements with
        | [] -> agg
        | head :: rest -> map rest (func head :: agg)

    List.rev (map input [])
