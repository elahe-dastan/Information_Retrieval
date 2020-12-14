# Information_Retrieval

This is my information retrieval course project probably you don't need the whole project (if you don't have the course
in amirkabir next semesters :wink: but the ideas and algorithms can be used in different situations and I try to explain
them here.

## Memory Shortage
This problem is the most important thing I'd like to solve in this project cause if we have as much memory as we like indexing
a document collection will be a piece of cake so in order to simulate the situation for huge document collection I assume
I have only 200 bytes of memory 

### Implementation point
In the previous part I said I wanted to assume I have a 200 byte memory but this assumption will cause a headache in future
cause I can't guess how to read the file in order to both read the file as much as possible and also read the tokens completely,
let's simply use a trick I assume my memory can place 6 words :stuck_out_tongue_winking_eye: .

### BSBI
I want to read documents term by term and keep term id pair in a sorted way but I can't read all the documents at once 
cause I don't have enough memory so I use BSBI algorithm
 
# Make invertible index
####  prepare documents

# the memory is 160 bytes
