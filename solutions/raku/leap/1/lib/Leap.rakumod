unit module Leap;

sub is-leap-year ($year) is export {
    if $year %% 100 {
        if $year %% 400 {
            return True;
        } else {
            return False;
        }
    }
    return $year %% 4;
}
