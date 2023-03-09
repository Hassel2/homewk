import Text.Printf

teen :: Int -> Bool
teen x = x >= 13 && x <= 19

main = do
    printf "(teen? 14)\n  => "
    print (teen 14)
    printf "(teen? 12)\n  => "
    print (teen 12)
    printf "(teen? 20)\n  => "
    print (teen 20)
