import Text.Printf

collatz :: Int -> [Int]
collatz x = col x (a ++ [x])
    where 
        a = []
        col :: Int -> [Int] -> [Int]
        col 1 a = a
        col x a
            | even x = col (x `div` 2) (a ++ [x `div` 2])
            | odd x = col (x*3+1) (a ++ [x*3+1])

main = do
    printf "(collatz 3)\n  => "
    print (collatz 3)
    printf "(collatz 9)\n  => "
    print (collatz 9)
    printf "(collatz 17)\n  => "
    print (collatz 17)
    printf "(collatz 20)\n  => "
    print (collatz 20)
    printf "(collatz 50)\n  => "
    print (collatz 50)
