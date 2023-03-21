module First where

first :: [a] -> a
first [] = error "Empty list"
first x = head x

second :: [a] -> a 
second [] = error "Empty list"
second x 
    | length x < 2 = error "List is too short"
    | otherwise = head (tail x)

third :: [a] -> a 
third [] = error "Empty list"
third x 
    | length x < 3 = error "List is too short"
    | otherwise = head (tail (tail x))
