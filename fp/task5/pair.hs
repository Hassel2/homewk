import Text.Printf

makePolarPair :: Float -> Float -> (Float, Float)
makePolarPair rad alph= (rad, alph)

getPolarPairRadius :: (Float, Float) -> Float
getPolarPairRadius = fst

getPolarPairDegree :: (Float, Float) -> Float
getPolarPairDegree = snd

main = do
    printf "(make-polar-pair 3 60)\n  => "
    let pp = makePolarPair 3 60 
    print pp
    printf "(get-polar-pair-radius pp)\n  => " 
    print (getPolarPairRadius pp)
    printf "(get-polar-pair-degree pp)\n  => " 
    print (getPolarPairDegree pp)
