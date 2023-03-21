main = do
    let names = ["Elena", "Vasiliy"]
    let surnames = ["Tumanova", "Vlasov"]
    let patronymics = ["Ivanovna", "Petrovich"]
    let snp = [x ++ " " ++ y ++ " " ++ z | (x, y, z) <- zip3 surnames names patronymics]
    let nsp = [y ++ " " ++ x ++ " " ++ z | (x, y, z) <- zip3 surnames names patronymics]
    print snp
    print nsp
    print (reverse nsp)
    
