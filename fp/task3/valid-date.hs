import Text.Printf

isLeapYear :: Int -> Bool
isLeapYear y 
    | y `mod` 4 == 0 && y `mod` 100 /= 0 = True
    | y `mod` 400 == 0 = True
    | otherwise = False

checkDays :: Int -> Bool
checkDays m 
    | m `elem` [2, 4, 6, 9, 11] = False
    | otherwise = True

validDate :: Int -> Int -> Int -> Bool
validDate d m y 
    | d * m * y <= 0 = False
    | m > 12 = False
    | d > 31 || m == 2 && d > 29 = False
    | d == 31 = checkDays m
    | m == 2 && d == 29 = isLeapYear y
    | otherwise = True

main = do
    printf "(valid-date? 7 8 1949)\n  => "
    print (validDate 7 8 1949)
    printf "(valid-date? -7 8 1949)\n  => "
    print (validDate (-7) 8 1949)
    printf "(valid-date? 1 0 1949)\n  => "
    print (validDate 1 0 1949)
    printf "(valid-date? 29 2 1900)\n  => "
    print (validDate 29 2 1900)
    printf "(valid-date? 29 2 2000)\n  => "
    print (validDate 29 2 2000)
    printf "(valid-date? 31 9 2000)\n  => "
    print (validDate 31 9 2000)
    printf "(valid-date? 30 2 2000)\n  => "
    print (validDate 30 2 2000)
    printf "(valid-date? 29 5 2000)\n  => "
    print (validDate 29 5 2000)
    printf "(valid-date? 99 8 2000)\n  => "
    print (validDate 99 8 2000)
