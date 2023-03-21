import First

cmp3 :: (Ord a) => [a] -> [a] -> String
cmp3 a b = ["less", "more", "first 3 element are equal"] !! cmp [first a, second a, third a] [first b, second b, third b] 0
    where
        cmp :: (Ord a) => [a] -> [a] -> Int -> Int
        cmp a b i
            | a !! i > b !! i = 1
            | a !! i < b !! i = 0
            | a !! i == b !! i && i == 2 = 2
            | otherwise = cmp a b (i + 1)
