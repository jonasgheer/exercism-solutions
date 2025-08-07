module Raindrops

let convert (number) =
    let mutable sound = ""

    if number % 3 = 0 then
        sound <- sound + "Pling"

    if number % 5 = 0 then
        sound <- sound + "Plang"

    if number % 7 = 0 then
        sound <- sound + "Plong"

    if sound = "" then
        sound <- sprintf "%d" number

    sound
