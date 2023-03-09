import Text.Printf
tip :: Float -> Float
tip x = x * 0.15

main = do
    printf "%.5f" (tip 7.54)
