#Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

##Algorithm

One-pass Hash Table. While we iterate and inserting elements into the table, we also look back to check if current element's complement already exists in the table. If it exists, we have found a solution and return immediately.

Time complexity O(n). We traverse the list containing n elements only once.
Space complexity O(n). The extra space required depends on the number of items stored in the hash table, which stores at most n elements
