# Number to LCD
Goal: write a program that displays LCD style numbers.

## Part 1
Write a program that given a number (with arbitrary number of digits), converts it into LCD style numbers using the following format:

```
 _     _  _     _  _  _  _  _  
| |  | _| _||_||_ |_   ||_||_|  
|_|  ||_  _|  | _||_|  ||_| _|
```  
(each digit is 3 lines high)

Note: Please do NOT read the second part before completing the first. Part of the purpose of this kata is to make you practice refactoring and adapting to changing requirements. 

## Part 2
Change your program to support variable width or height of the digits. For example for width = 3 and height = 2 the digit 2 will be:

```
 ___
    |
 ___|
|
|___ 
```

w=3, h=3:

```
 ___ 
    |
    |
 ___|
|    
|    
|___ 
```