import Text.Printf
import Data.Maybe

bisection :: (Float -> Float) -> Float -> Float -> Float -> Maybe Float
bisection f a b eps 
    | (f a > 0 && f b > 0) || (f a < 0 && f b < 0) = Nothing
    | abs (f c) <= eps = Just c
    | (f a < 0 && f c > 0) || (f a > 0 && f c < 0) = bisection f a c eps
    | otherwise = bisection f c b eps
    where 
        c = (b + a) / 2


main = do
    printf "(bisection (lambda (x) x) -3 3. 0.001)\n  => "
    print (fromJust (bisection id (-3) 3 0.001))
    printf "(bisection (lambda (x) (sin x)) 2 4. 0.001)\n  => "
    print (fromJust (bisection sin 2 4 0.001))
    printf "(bisection (lambda (x) (cos x)) 0 2. 0.001)\n  => "
    print (fromJust (bisection cos 0 2 0.001))
    printf "(bisection (lambda (x) (* x x)) -1 1. 0.001)\n  => "
    print (bisection (**2) (-1) 1 0.001)
    printf "(bisection (lambda (x) (* x x x)) -1 1. 0.001)\n  => "
    print (fromJust (bisection (**3) (-1) 1 0.001))
