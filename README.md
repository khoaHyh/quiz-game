# README

A quiz game app built to learn golang as a part of Gophercises.
Link to repo [gophercises/quiz](https://github.com/gophercises/quiz).

## TODOS

### Part 1

Create a program that will read in a quiz provided via a CSV file and will then give 
the quiz to a user keeping track of how many questions they get right and how many they
get incorrect.

#### Details
  * regardless of whether the answer is correct or wrong the next question should be asked 
    immediately afterwards
  * user should be able to customize the filename via a flag
  * csv file will be in a format like below:
    ```csv
    5+5,10
    7+3,10
    1+1,2
    8+3,11
    1+2,3
    8+6,14
    3+1,4
    1+4,5
    5+1,6
    2+3,5
    3+3,6
    2+4,6
    5+2,7
    ```
    * where the 1st column is a question and the second column in the same row is the answer to that question
  * quizzes will be relatively short (<100 questions) with single word/number answers
  * the end of the quiz should output the total number of questions correct and how many 
    questions there were in total

### Part 2

The program from part 1 will be adapted with the addition of a timer.

#### Details
  * default time limit should be 30 seconds
    * should be customizable via a flag
  * quiz should stop as soon as the time limit exceeded
    * shuld stop the quiz entirely even if the user is currently answering a question
  * users should be asked to press enter (or some other key) before the timer starts
  * questions should be printed out to the screen one at a time until the user provides an answer
  * regardless of whether the answer is correct or wrong the next question should be asked
  * at the end of the quiz the program should still output the total number of questions correct
    and how many questions there were in total

### Bonus

1. add strin  trimming and cleanup to help ensure that correct answers with whitespace, capitalizaiton,
   etc are not considered incorrect.
2. Add an option (a new flag) to shuffle the quiz order each time it is run
