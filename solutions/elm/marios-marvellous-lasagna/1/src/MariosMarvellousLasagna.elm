module MariosMarvellousLasagna exposing (remainingTimeInMinutes)


remainingTimeInMinutes numLayers minutesInOven =
    let
        expectedMinutesInOven = 40
        preparationTimeInMinutes layers =
            layers * 2
    in
    preparationTimeInMinutes numLayers + expectedMinutesInOven - minutesInOven
    
