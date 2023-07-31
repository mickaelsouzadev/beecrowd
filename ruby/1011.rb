pi = 3.14159

line = gets.split

r = line[0].to_f

volume = (4 / 3.0) * pi * (r ** 3)

message = sprintf("VOLUME = %.3f\n", volume)
puts message