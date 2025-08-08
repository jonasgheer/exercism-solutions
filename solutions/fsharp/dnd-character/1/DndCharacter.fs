module DndCharacter

type Character =
    { Strength: int
      Dexterity: int
      Constitution: int
      Intelligence: int
      Wisdom: int
      Charisma: int
      Hitpoints: int }

let modifier x = int (floor ((float (x) - 10.) / 2.))

let ability () =
    let rnd = System.Random()

    [ for _ in [ 1 .. 4 ] do
          rnd.Next(1, 7) ]
    |> List.sort
    |> List.skip 1
    |> List.sum

let createCharacter () =
    let constitution = ability ()

    { Strength = ability ()
      Wisdom = ability ()
      Dexterity = ability ()
      Constitution = constitution
      Intelligence = ability ()
      Charisma = ability ()
      Hitpoints = 10 + modifier constitution }
