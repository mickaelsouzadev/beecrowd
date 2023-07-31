line1 = gets.split
line2 = gets.split

code1 = line1[0].to_i
piece1 = line1[1].to_i
price1 = line1[2].to_f

code2 = line2[0].to_i
piece2 = line2[1].to_i
price2 = line2[2].to_f

total = (price1 * piece1) + (price2 * piece2)

message = sprintf("VALOR A PAGAR: R$ %.2f\n", total)
puts message