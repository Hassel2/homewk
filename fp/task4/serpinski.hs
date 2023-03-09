import Text.Printf
 
serpinski :: Int -> [String]
serpinski x = [
            block, block, block,
            block, block, block,
            block, block, block
        ]
    where
        block = ser (x - 1)
        ser :: Int -> String
        ser x 
            | x == 1 = "█"
            | otherwise = "█"

main = do
    print (serpinski 1)

{-             ███
 -             █ █
 -             ███
 -
 -          █████████
 -          █ ██ ██ █
 -          █████████
 -          ███   ███
 -          █ █   █ █
 -          ███   ███
 -          █████████
 -          █ ██ ██ █
 -          █████████
 -
 - ███████████████████████████
 - █ ██ ██ ██ ██ ██ ██ ██ ██ █
 - ███████████████████████████
 - ███   ██████   ██████   ███
 - █ █   █ ██ █   █ ██ █   █ █
 - ███   ██████   ██████   ███
 - ███████████████████████████
 - █ ██ ██ ██ ██ ██ ██ ██ ██ █
 - ███████████████████████████
 - █████████         █████████
 - █ ██ ██ █         █ ██ ██ █
 - █████████         █████████
 - ███   ███         ███   ███
 - █ █   █ █         █ █   █ █
 - ███   ███         ███   ███
 - █████████         █████████
 - █ ██ ██ █         █ ██ ██ █
 - █████████         █████████
 - ███████████████████████████
 - █ ██ ██ ██ ██ ██ ██ ██ ██ █
 - ███████████████████████████
 - ███   ██████   ██████   ███
 - █ █   █ ██ █   █ ██ █   █ █
 - ███   ██████   ██████   ███
 - ███████████████████████████
 - █ ██ ██ ██ ██ ██ ██ ██ ██ █
 - ███████████████████████████ -}