module TracksOnTracksOnTracks exposing (..)

import List

newList : List String
newList = []


existingList : List String
existingList = [ "Elm", "Clojure", "Haskell" ]


addLanguage : String -> List String -> List String
addLanguage language languages = language :: languages
    

countLanguages : List String -> Int
countLanguages languages = List.length languages


reverseList : List String -> List String
reverseList languages = List.reverse languages


excitingList : List String -> Bool
excitingList languages = 
    case languages of 
        ["Elm"] -> True
        "Elm" :: _ -> True
        _ :: "Elm" :: rest -> if List.length rest > 1 then False else True
        _ -> False
