import Text.Printf

fun :: Float -> Float -> Float -> Float
fun x y z
  | (x > y && y > z) || (y > x && x > z) = x**2 + y**2
  | (y > z && z > x) || (z > y && y > x) = z**2 + y**2
  | otherwise = x**2 + z**2

main = do
    printf "  => %f\n" (fun 3 2 1)
    printf "(fun 2 3 1)\n"
    printf "  => %f\n" (fun 2 3 1)
    printf "(fun 2 1 3)\n"
    printf "  => %f\n" (fun 2 1 3)
    printf "(fun 1 2 3)\n"
    printf "  => %f\n" (fun 1 2 3)
