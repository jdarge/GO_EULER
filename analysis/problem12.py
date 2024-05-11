import matplotlib.pyplot as plt

# Read numbers from output.txt
with open('output12.txt', 'r') as file:
    numbers = [int(line.strip()) for line in file]

# Create x-values (line numbers) from 1 to the number of lines in the file
x_values = list(range(1, len(numbers) + 1))

# Plotting
plt.plot(x_values, numbers)
plt.xlabel('Line Number')
plt.ylabel('Number')
plt.title('Graph of Numbers from output.txt')
plt.show()
