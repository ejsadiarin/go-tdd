# Learning Go with Tests (focuses on TDD)
- ejsadiarin: started at 2024-01-01 _ for january (ejsadiarin month) 

## Learnings
- named return value for functions are good for `private` functions that is used for the `Public` functions
- TDD Philosophy:
  - Write tests first
  - Write minimal function code to make test fail (run test) - to see if err message is good and "descriptive"
  - Write the minimum amount of code to make the tests pass
  - Refactor the code

## Discipline
Let's go over the cycle again:
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor
 
On the face of it this may seem tedious but sticking to the feedback loop is important.
 
Not only does it ensure that you have relevant tests, it helps ensure you design good software by refactoring with the safety of tests.
