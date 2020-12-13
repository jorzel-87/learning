#!/usr/bin/python3

import random

secretNumber = random.randint(1,10)

for i in range(1,6):
    print("Guess!:")
    guess = int(input())
    if guess < secretNumber:
        print('too low..')
    elif guess > secretNumber:
        print('too high..')
    else:
        break
if guess == secretNumber:
    print("Good, good! You guessed my number in " + str(i) + " guesses!")
else:
    print("Noooooo!")