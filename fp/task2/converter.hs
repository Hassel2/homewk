celToFar :: Float -> Float 
celToFar celsius = (celsius * 9/5) + 32 

farToCel :: Float -> Float
farToCel fahrenheit = (fahrenheit - 32) * 5/9

main = do
    print "Write number of celsius you want to convert in fahrenheit"
    cel <- getLine 
    print(celToFar (read cel :: Float))
    print "Write number of fahrenheit you want to convert in celsius"
    far <- getLine
    print(farToCel (read far :: Float))
