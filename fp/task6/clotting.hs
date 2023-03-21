main = do
    let l = [1..10]
    print (foldl (+) 0 l)
    print (foldl (+) 0 [x | x <- l, even x])
    print (foldl (-) 0 [x | x <- l, odd x])
