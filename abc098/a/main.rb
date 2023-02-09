line = gets.split(' ').map(&:to_i)

a = line[0]
b = line[1]

p [a+b, a-b, a*b].max()
