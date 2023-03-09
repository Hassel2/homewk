import Text.Printf
sci :: Float -> Float -> Float
sci x pow = x * (10 ** pow)

main = do
    printf "%.5f" (sci 0.35 (-3))
