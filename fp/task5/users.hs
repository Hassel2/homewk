import Text.Printf

userAges :: String -> [(String, Int)] -> (String, Int)
userAges key dict
    | key `elem` [fst x | x <- dict] = head (filter (\(x, y) -> x == key) dict)
    | otherwise = ("", 0)

{- userAges "Dick" [("Eugene", 13),("Geoerge", 44),("Jack", 35),("Denis", 10)]
 - userAges "Jack" [("Eugene", 13),("Geoerge", 44),("Jack", 35),("Denis", 10)] -}

