{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use min" #-}
{-# HLINT ignore "Use max" #-}
import Text.Printf

ma :: Float -> Float -> Float
ma x y = if x > y
            then x
            else y

mi :: Float -> Float -> Float
mi x y = if x < y
            then x
            else y

main = do
    printf "(ma 3 4)\n"
    printf "  => %f\n" (ma 3 4)
    printf "(mi 3 4)\n"
    printf "  => %f\n" (mi 3 4)
    printf "(fun 3 2 1)\n"
