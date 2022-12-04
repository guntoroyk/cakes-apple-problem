# Problem 3

- Type: Closed problem

## Scenario

Ainun have 20 cakes and 25 apples. She want to bundle that cakes and apples into boxes and give them to her friends. How many boxes that Ainun can make? And how many cakes and apples every box have?

## Terms and conditions

- Solve this problem using language you know and confident
- Create function to count boxes Ainun can make
- Create function to count how many cake and apple in a box
- Create unit test for the function
- Apples and cakes divided evenly every box
- Don’t forget to make documentation

## Bonus point

- Implementation in Go.

## Solution

1. Find the greatest common factor (GCF) between cakes and apples to get the number of boxes
2. Divide cakes and apples with the GCF value to get the number of cakes and apples each box

Example:

cakes = 20
apples = 25

the factors of each arguments are:

`20 = 1, 2, 4, 5, 10, 20`

`25 = 1, 5, 25`

To get the GCF, we can loop through reading the numbers from the smaller argument (in this case is `20`) to 2. If the number is divisible by both arguments, then that number is the GCF. If until 2 there's no number found, it means the GCF is 1.

We find that 5 is the first number that is divisible by both arguments, so the GCF (total number of boxes) is <b>5</b>.

Then, we can get the cakes and apple for each box:

`cakes = 20 / 5 = 4`

`apples = 25 / 5 = 5`

## Test the code solution

Run all the unit test

```
go test ./... -v
```

Run example case

```
go run main.go
```
