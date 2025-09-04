package hamming

// Compute calculates the Hamming distance between two DNA strands.

// Logic is as follows:
// 1. If the lengths of the two strands are not equal, return an error.
// 2. Initialize a counter to zero.
// 3. Iterate through the characters of both strands simultaneously.
// 4. For each position, if the characters differ, increment the counter.
// 5. Return the counter as the Hamming distance.
