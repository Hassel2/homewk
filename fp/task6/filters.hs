main = do
    let l = [1..10]
    print (filter even l)
    print (filter odd l)
    print (filter (\x -> even x && x `div` 3 == 0) l)
    print (filter (\x -> even x || x^2 > 10) l)
