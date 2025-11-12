# LeetCode Solutions in Go

This repository contains my personal collection of LeetCode problem solutions implemented in Go (Golang). The solutions are organized by problem categories and include comprehensive documentation, test cases, and multiple approaches where applicable.

## 📁 Repository Structure

The repository is organized into the following categories:

### 🔢 Arrays

- **[Best Time to Buy and Sell Stock](Arrays/best-time-to-buy-and-sell-stock/)** - Two-pointer technique for maximum profit optimization
- **[Contains Duplicate](Arrays/contains-duplicate/)** - Duplicate detection in arrays
- **[Contains Nearby Duplicate](Arrays/contains-nearby-duplicate/)** - Proximity-based duplicate detection
- **[Defuse the Bomb](Arrays/defuse-the-bomb/)** - Circular array manipulation (Sliding Window)
- **[Find Median of Two Sorted Arrays](Arrays/find-median-sorted-arrays/)** - Binary search on sorted arrays
- **[Group Anagrams](Arrays/group-anagrams/)** - String grouping and hash map techniques
- **[Longest Common Prefix](Arrays/longest-common-prefix/)** - String prefix matching
- **[Longest Consecutive Sequence](Arrays/longest-consecutive-sequence/)** - Sequence detection with O(n) complexity
- **[Longest Harmonious Subsequence](Arrays/longest-harmonious-subsequence/)** - Array analysis with sliding window
- **[Maximum Strong Pair XOR](Arrays/maximum-strong-pair-xor/)** - XOR operations and pair analysis (Sliding Window)
- **[Product of Array Except Self](Arrays/product-of-array-except-self/)** - Array manipulation without division
- **[Three Sum](Arrays/three-sum/)** - Two-pointer technique for triplet finding
- **[Top K Frequent Elements](Arrays/top-k-frequent-elements/)** - Frequency analysis and sorting
- **[Trapping Rain Water](Arrays/trapping-rain-water/)** - Two-pointer technique for water trapping optimization
- **[Two Sum II - Input Array Is Sorted](Arrays/two-sum-ii-input-array-is-sorted/)** - Two-pointer technique on sorted arrays
- **[Valid Anagram](Arrays/valid-anagram/)** - Character frequency comparison
- **[Valid Sudoku](Arrays/valid-sudoku/)** - 2D array validation with complex rules

### 🔍 Binary Search

- **[Binary Search](Binary-Search/binary-search/)** - Classic O(log n) search in sorted arrays
- **[Find Minimum in Rotated Sorted Array](Binary-Search/find-minimum-in-rotated-sorted-array/)** - Binary search in rotated sorted array
- **[Find the Distance Value Between Two Arrays](Binary-Search/find-the-distance-value-between-two-arrays/)** - Using binary search to count elements with minimum distance
- **[First Bad Version](Binary-Search/first-bad-version/)** - Finding first bad version using binary search
- **[Koko Eating Bananas](Binary-Search/koko-eating-bananas/)** - Binary search for optimization problems
- **[Median of Two Sorted Arrays](Binary-Search/median-of-two-sorted-arrays/)** - Binary search approach to find median efficiently
- **[Missing Number](Binary-Search/missing-number/)** - Finding missing number using binary search or mathematical approach
- **[Search in Rotated Sorted Array](Binary-Search/search-in-rotated-sorted-array/)** - Modified binary search for rotated arrays
- **[Search Insert Position](Binary-Search/search-insert-position/)** - Finding insertion position using binary search
- **[Sqrt(x)](Binary-Search/sqrtx/)** - Computing square root using binary search
- **[Time Based Key-Value Store](Binary-Search/time-based-key-value-store/)** - Binary search for time-based data retrieval

### 🗂️ Hash Table

- **[Substrings of Size Three with Distinct Characters](Hash-Table/substrings-of-size-three-with-distinct-characters/)** - Fixed-size window for unique character substrings
- **[Two Sum](Hash-Table/two-sum/)** - Hash map for O(1) lookups in pair sum problems

### 🔗 Linked List

- **[Add Two Numbers](Linked-List/add-two-number/)** - Linked list arithmetic operations
- **[Add Two Numbers](Linked-List/add-two-numbers/)** - Linked list arithmetic operations
- **[Copy List with Random Pointer](Linked-List/copy-list-with-random-pointer/)** - Deep copying of complex linked structure
- **[Linked List Cycle](Linked-List/linked-list-cycle/)** - Floyd's cycle detection algorithm
- **[LRU Cache](Linked-List/lru-cache/)** - Linked list and hashmap implementation for cache
- **[Merge Two Sorted Lists](Linked-List/merge-two-sorted-lists/)** - Iterative linked list merging
- **[Remove Nth Node From End of List](Linked-List/remove-nth-node-from-end-of-list/)** - Two-pointer technique for node removal
- **[Reorder List](Linked-List/reorder-list/)** - Multiple linked list operations combined
- **[Reverse Linked List](Linked-List/reverse-linked-list/)** - In-place linked list reversal

### 🧮 Math

- **[Convert Date to Binary](Math/convert-date-to-bin/)** - Date manipulation and binary conversion
- **[Greatest Common Divisor of Strings](Math/greatest-common-divisor-of-strings/)** - String pattern analysis

### 📊 Matrix

- **[Check Valid Matrix](Matrix/check-valid/)** - Matrix validation algorithms
- **[Search a 2D Matrix](Matrix/search-a-2d-matrix/)** - Binary search in 2D matrix

### 📚 Stack

- **[Car Fleet](Stack/car-fleet/)** - Stack-based simulation of cars on a road
- **[Daily Temperatures](Stack/daily-temperatures/)** - Monotonic stack for next greater element problems
- **[Evaluate Reverse Polish Notation](Stack/evaluate-reverse-polish-notation/)** - Stack-based expression evaluation with custom stack implementation
- **[Min Stack](Stack/min-stack/)** - Stack data structure with constant-time minimum retrieval

### 📉 Monotonic Stack

- **[Final Prices With a Special Discount in a Shop](Monotonic-Stack/final-prices-with-a-special-discount-in-a-shop/)** - Next smaller element pattern
- **[Largest Rectangle in Histogram](Monotonic-Stack/largest-rectangle-in-histogram/)** - Stack-based area calculation with monotonic properties
- **[Next Greater Element I](Monotonic-Stack/next-greater-element-i/)** - Monotonic decreasing stack for next greater element

### 🔤 String

- **[Count K-Constraint Substrings](String/count-k-constraint-substrings/)** - Sliding window technique for substring counting
- **[Generate Parentheses](String/generate-parentheses/)** - Backtracking algorithm for combinatorial generation
- **[Valid Palindrome](String/valid-palindrome/)** - Two-pointer technique with string preprocessing
- **[Valid Parentheses](String/valid-paretheses/)** - Stack-based bracket matching

### 🌳 Tree

- **[Balanced Binary Tree](Tree/balanced-binary-tree/)** - Tree balance validation
- **[Binary Tree Level Order Traversal](Tree/binary-tree-level-order-traversal/)** - BFS approach for level-by-level tree traversal
- **[Binary Tree Right Side View](Tree/binary-tree-right-side-view/)** - Modified BFS to view tree from right side
- **[Construct Binary Tree from Preorder and Inorder Traversal](Tree/construct-binary-tree-from-preorder-and-inorder-traversal/)** - Tree reconstruction from traversals
- **[Count Good Nodes in Binary Tree](Tree/count-good-nodes-in-binary-tree/)** - DFS with path maximum tracking
- **[Diameter of Binary Tree](Tree/diameter-of-binary-tree/)** - Maximum path length calculation
- **[Invert Binary Tree](Tree/invert-binary-tree/)** - Tree mirroring operation
- **[Kth Smallest Element in a BST](Tree/kth-smallest-element-in-a-bst/)** - Efficient BST traversal for selection
- **[Lowest Common Ancestor of a Binary Search Tree](Tree/lowest-common-ancestor-of-a-binary-search-tree/)** - BST property-based traversal
- **[Maximum Depth of Binary Tree](Tree/maximum-depth-of-binary-tree/)** - Recursive tree height calculation
- **[Same Tree](Tree/same-tree/)** - Tree comparison and structural equality
- **[Subtree of Another Tree](Tree/subtree-of-another-tree/)** - Subtree matching algorithm
- **[Validate Binary Search Tree](Tree/validate-binary-search-tree/)** - BST property validation using inorder traversal

### 🔍 Sliding Window

- **[Find the K-Beauty of a Number](Sliding-Window/find-the-k-beauty-of-a-number/)** - Fixed-size sliding window over digits
- **[Longest Nice Substring](Sliding-Window/longest-nice-substring/)** - Recursive sliding window for substring property validation
- **[Longest Repeating Character Replacement](Sliding-Window/longest-repeating-character-replacement/)** - Sliding window with character frequency counting
- **[Longest Substring Without Repeating Characters](Sliding-Window/longest-substring-without-repeating-characters/)** - Dynamic sliding window technique
- **[Maximum Average Subarray I](Sliding-Window/maximum-average-subarray-i/)** - Fixed-size sliding window with running sum
- **[Minimum Difference Between Highest and Lowest of K Scores](Sliding-Window/minimum-difference-between-highest-and-lowest-of-k-scores/)** - Fixed-size window after sorting
- **[Minimum Window Substring](Sliding-Window/minimum-window-substring/)** - Variable-size sliding window optimization
- **[Permutation in String](Sliding-Window/permutation-in-string/)** - Fixed-size sliding window with pattern matching

### ↔️ Two Pointers

- **[Assign Cookies](Two-Pointers/assign-cookies/)** - Greedy assignment using two pointers
- **[Check If N and Its Double Exist](Two-Pointers/check-if-n-and-its-double-exist/)** - Finding pairs where one number is double of another using two pointers
- **[Container With Most Water](Two-Pointers/container-with-most-water/)** - Two-pointer technique for maximum area between lines
- **[Count Binary Substrings](Two-Pointers/count-binary-substrings/)** - Counting consecutive groups with two pointers
- **[DI String Match](Two-Pointers/di-string-match/)** - Constructing permutation based on increase/decrease pattern using two pointers
- **[Duplicate Zeros](Two-Pointers/duplicate-zeros/)** - In-place array modification duplicating zeros using two pointers
- **[Find the Duplicate Number](Two-Pointers/find-the-duplicate-number/)** - Floyd's cycle detection in array form
- **[Find the Index of the First Occurrence in a String](Two-Pointers/find-the-index-of-the-first-occurrence-in-a-string/)** - String matching with two pointers
- **[Happy Number](Two-Pointers/happy-number/)** - Floyd's cycle detection for number sequences
- **[Intersection of Two Arrays](Two-Pointers/intersection-of-two-arrays/)** - Finding common elements between arrays
- **[Intersection of Two Arrays II](Two-Pointers/intersection-of-two-arrays-ii/)** - Finding all common elements with duplicates
- **[Is Subsequence](Two-Pointers/is-subsequence/)** - Efficient subsequence validation with two pointers
- **[Long Pressed Name](Two-Pointers/long-pressed-name/)** - Validating typed name with possible long key presses using two pointers
- **[Move Zeroes](Two-Pointers/move-zeroes/)** - In-place array rearrangement with two pointers
- **[Remove Duplicates from Sorted Array](Two-Pointers/remove-duplicates-from-sorted-array/)** - In-place array modification with two pointers
- **[Remove Element](Two-Pointers/remove-element/)** - In-place removal with two-pointer approach
- **[Remove Palindromic Subsequences](Two-Pointers/remove-palindromic-subsequences/)** - Palindrome validation and removal analysis using two pointers
- **[Reverse String](Two-Pointers/reverse-string/)** - In-place string reversal with two pointers
- **[Reverse Only Letters](Two-Pointers/reverse-only-letters/)** - Reverse only letters while keeping non-letter characters in place
- **[Reverse String II](Two-Pointers/reverse-string-ii/)** - Partial string reversal with two pointers
- **[Reverse Vowels of a String](Two-Pointers/reverse-vowels-of-a-string/)** - Selective character reversal using two pointers
- **[Reverse Words in a String III](Two-Pointers/reverse-words-in-a-string-iii/)** - Reverse words within a string using two pointers
- **[Shortest Distance to a Character](Two-Pointers/shortest-distance-to-a-character/)** - Computing minimum distances with two-pass approach
- **[Sort Array by Parity](Two-Pointers/sort-array-by-parity/)** - In-place partitioning with two pointers
- **[Sort Array by Parity II](Two-Pointers/sort-array-by-parity-ii/)** - Arranging array with even indices for even numbers and odd indices for odd numbers
- **[Squares of a Sorted Array](Two-Pointers/squares-of-a-sorted-array/)** - Building sorted squares array using two pointers from both ends
- **[Valid Palindrome II](Two-Pointers/valid-palindrome-ii/)** - Palindrome validation with one deletion allowed

## 🛠️ Technical Features

### Programming Language

- **Go (Golang)** - All solutions implemented with idiomatic Go code
- Comprehensive use of Go's built-in data structures and standard library

### Code Quality

- **Clean Code Practices**: Well-documented functions with clear variable names
- **Multiple Approaches**: Many problems include both brute force and optimized solutions
- **Time & Space Complexity**: Comments indicating algorithmic complexity
- **Test Cases**: Comprehensive test suites using Go's testing framework

### Algorithmic Techniques

- **Sliding Window**: Multiple implementations showcasing different sliding window patterns
- **Two Pointers**: Efficient array traversal techniques
- **Hash Maps**: Fast lookup and frequency counting
- **Dynamic Programming**: Optimal substructure solutions
- **Binary Search**: Logarithmic search algorithms
- **Stack & Queue**: LIFO and FIFO data structure utilization
- **Tree Traversal**: DFS and BFS implementations
- **Sorting Algorithms**: Custom and built-in sorting techniques

## 📈 Difficulty Levels

The repository includes problems across all difficulty levels:

- **Easy**: Fundamental algorithmic concepts
- **Medium**: Intermediate problem-solving techniques
- **Hard**: Advanced algorithms and optimizations

## 🔍 Special Features

### Sliding Window Problems

This repository contains excellent examples of sliding window techniques:

- **Fixed Size Window**: Count Good Substrings
- **Variable Size Window**: Count K-Constraint Substrings
- **Multiple Window Approaches**: Maximum Strong Pair XOR

### Optimization Examples

- **O(n) Solutions**: Longest Consecutive Sequence
- **Space Optimization**: Product of Array Except Self
- **Early Termination**: Various sliding window implementations

### Testing

- **Unit Tests**: Comprehensive test coverage with table-driven tests
- **Edge Cases**: Boundary condition testing
- **Performance Testing**: Benchmarking for optimization validation

## 🚀 Getting Started

### Prerequisites

- Go 1.19+ installed on your system
- Basic understanding of Go syntax and data structures

### Running Solutions

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd leetcode
   ```

2. Navigate to any problem directory:

   ```bash
   cd Arrays/two-sum
   ```

3. Run the solution:

   ```bash
   go run main.go
   ```

4. Run tests (where available):
   ```bash
   go test
   ```

### Building Solutions

Each problem can be built as a standalone executable:

```bash
go build main.go
./main
```

## 📊 Problem Categories by Data Structure

| Data Structure  | Count | Examples                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| --------------- | ----- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Arrays          | 17    | Best Time to Buy and Sell Stock, Contains Duplicate, Contains Nearby Duplicate, Defuse the Bomb, Find Median of Two Sorted Arrays, Group Anagrams, Longest Common Prefix, Longest Consecutive Sequence, Longest Harmonious Subsequence, Maximum Strong Pair XOR, Product of Array Except Self, Three Sum, Top K Frequent Elements, Trapping Rain Water, Two Sum II - Input Array Is Sorted, Valid Anagram, Valid Sudoku                                                                                                                                                                                                                                                                  |
| Binary Search   | 11    | Binary Search, Find Minimum in Rotated Sorted Array, Find the Distance Value Between Two Arrays, First Bad Version, Koko Eating Bananas, Median of Two Sorted Arrays, Missing Number, Search in Rotated Sorted Array, Search Insert Position, Sqrt(x), Time Based Key-Value Store                                                                                                                                                                                                                                                                                                                                                                                                        |
| Hash Tables     | 2     | Substrings of Size Three with Distinct Characters, Two Sum                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| Linked Lists    | 9     | Add Two Numbers (2 versions), Copy List with Random Pointer, Linked List Cycle, LRU Cache, Merge Two Sorted Lists, Remove Nth Node From End of List, Reorder List, Reverse Linked List                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| Math            | 2     | Convert Date to Binary, Greatest Common Divisor of Strings                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| Matrix          | 2     | Check Valid Matrix, Search a 2D Matrix                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| Monotonic Stack | 3     | Final Prices With a Special Discount in a Shop, Largest Rectangle in Histogram, Next Greater Element I                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| Sliding Window  | 8     | Find the K-Beauty of a Number, Longest Nice Substring, Longest Repeating Character Replacement, Longest Substring Without Repeating Characters, Maximum Average Subarray I, Minimum Difference Between Highest and Lowest of K Scores, Minimum Window Substring, Permutation in String                                                                                                                                                                                                                                                                                                                                                                                                   |
| Stacks          | 4     | Car Fleet, Daily Temperatures, Evaluate Reverse Polish Notation, Min Stack                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| Strings         | 4     | Count K-Constraint Substrings, Generate Parentheses, Valid Palindrome, Valid Parentheses                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| Trees           | 13    | Balanced Binary Tree, Binary Tree Level Order Traversal, Binary Tree Right Side View, Construct Binary Tree from Preorder and Inorder Traversal, Count Good Nodes in Binary Tree, Diameter of Binary Tree, Invert Binary Tree, Kth Smallest Element in a BST, Lowest Common Ancestor of a Binary Search Tree, Maximum Depth of Binary Tree, Same Tree, Subtree of Another Tree, Validate Binary Search Tree                                                                                                                                                                                                                                                                              |
| Two Pointers    | 27    | Assign Cookies, Check If N and Its Double Exist, Container With Most Water, Count Binary Substrings, DI String Match, Duplicate Zeros, Find the Duplicate Number, Find the Index of the First Occurrence in a String, Happy Number, Intersection of Two Arrays, Intersection of Two Arrays II, Is Subsequence, Long Pressed Name, Move Zeroes, Remove Duplicates from Sorted Array, Remove Element, Remove Palindromic Subsequences, Reverse Only Letters, Reverse String, Reverse String II, Reverse Vowels of a String, Reverse Words in a String III, Shortest Distance to a Character, Sort Array by Parity, Sort Array by Parity II, Squares of a Sorted Array, Valid Palindrome II |

## 🎯 Algorithm Patterns

### Most Common Patterns

1. **Hash Map Lookups** - Fast O(1) access patterns
2. **Two Pointers** - Efficient array traversal
3. **Sliding Window** - Substring and subarray problems
4. **Sorting + Binary Search** - Optimized search operations
5. **Stack Operations** - LIFO processing patterns

### Advanced Techniques

- **Bit Manipulation** - XOR operations for unique solutions
- **Mathematical Optimization** - GCD and modular arithmetic
- **Tree Algorithms** - Recursive depth calculations
- **Dynamic Programming** - Optimal subproblem solutions

## 🔧 Code Organization

Each problem directory typically contains:

- `main.go` - Primary solution implementation
- `README.md` - Problem description and approach explanation
- `*_test.go` - Unit tests (where applicable)
- `go.mod` - Module definition (for problems requiring external dependencies)

## 📚 Learning Resources

This repository serves as:

- **Algorithm Reference** - Practical implementations of common algorithms
- **Go Programming Examples** - Idiomatic Go code patterns
- **Interview Preparation** - Solutions to commonly asked coding questions
- **Performance Analysis** - Examples of optimization techniques

## 🤝 Contributing

While this is a personal practice repository, suggestions and improvements are welcome:

1. Fork the repository
2. Create a feature branch
3. Add your improvements
4. Submit a pull request

## 📄 License

This repository is for educational purposes. All LeetCode problems are property of LeetCode LLC.

## 📞 Contact

- **Repository Owner**: huypq02
- **Current Branch**: main
- **Language**: Go (Golang)

---

**Happy Coding! 🚀**

> "The best way to learn algorithms is to implement them yourself." - Practice makes perfect!
