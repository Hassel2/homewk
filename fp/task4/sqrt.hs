import Text.Printf

mysqrt :: Float -> Float
mysqrt x = f x 1
    where
        f x a 
            | abs(a - a1) <= 0.000001 = a1
            | otherwise = f x a1
            where
                a1= 0.5 * (a + x/a)

main = do
    printf "(mysqrt 9)\n"
    printf "  => %.13f\n" (mysqrt 9)
    printf "(mysqrt 2)\n"
    printf "  => %.13f\n" (mysqrt 2)
    printf "(mysqrt 3)\n"
    printf "  => %.13f\n" (mysqrt 4)
    printf "(mysqrt 169)\n"
    printf "  => %.13f\n" (mysqrt 169)
