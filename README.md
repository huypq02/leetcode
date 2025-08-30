# LeetCode Solutions in Go

This repository contains my personal collection of LeetCode problem solutions implemented in Go (Golang). The solutions are organized by problem categories and include comprehensive documentation, test cases, and multiple approaches where applicable.

## 📁 Repository Structure

The repository is organized into the following categories:

### 🔢 Arrays

- **[Best Time to Buy and Sell Stock](Arrays/best-time-to-buy-and-sell-stock/)** - Stock trading optimization
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
- **[Two Sum](Arrays/two-sum/)** - Hash map solution for pair sum problems
- **[Two Sum II - Input Array Is Sorted](Arrays/two-sum-ii-input-array-is-sorted/)** - Two-pointer technique on sorted arrays
- **[Valid Anagram](Arrays/valid-anagram/)** - Character frequency comparison
- **[Valid Sudoku](Arrays/valid-sudoku/)** - 2D array validation with complex rules

### 🗂️ Hash Table

- **[Count Good Substrings](Hash-Table/count-good-substring/)** - Substring analysis with distinct characters (Sliding Window)

### 🔗 Linked List

- **[Add Two Numbers](Linked-List/add-two-number/)** - Linked list arithmetic operations

### 🧮 Math

- **[Convert Date to Binary](Math/convert-date-to-bin/)** - Date manipulation and binary conversion
- **[Greatest Common Divisor of Strings](Math/greatest-common-divisor-of-strings/)** - String pattern analysis

### 📊 Matrix

- **[Check Valid Matrix](Matrix/check-valid/)** - Matrix validation algorithms

### 📚 Stack

- **[Daily Temperatures](Stack/daily-temperatures/)** - Monotonic stack for next greater element problems
- **[Evaluate Reverse Polish Notation](Stack/evaluate-reverse-polish-notation/)** - Stack-based expression evaluation with custom stack implementation
- **[Min Stack](Stack/min-stack/)** - Stack data structure with constant-time minimum retrieval

### 🔤 String

- **[Count K-Constraint Substrings](String/count-k-constraint-substrings/)** - Sliding window technique for substring counting
- **[Generate Parentheses](String/generate-parentheses/)** - Backtracking algorithm for combinatorial generation
- **[Valid Palindrome](String/valid-palindrome/)** - Two-pointer technique with string preprocessing
- **[Valid Parentheses](String/valid-paretheses/)** - Stack-based bracket matching

### 🌳 Tree

- **[Find Balanced Binary Tree](Tree/find-balanced-btree/)** - Tree traversal and balance checking

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

| Data Structure | Count | Examples                                                  |
| -------------- | ----- | --------------------------------------------------------- |
| Arrays         | 17    | Two Sum, Two Sum II, Three Sum, Valid Sudoku              |
| Hash Tables    | 1     | Count Good Substrings                                     |
| Linked Lists   | 1     | Add Two Numbers                                           |
| Stacks         | 3     | Daily Temperatures, Evaluate RPN, Min Stack               |
| Trees          | 1     | Balanced Binary Tree                                      |
| Strings        | 4     | Generate Parentheses, Valid Palindrome, Valid Parentheses |
| Math           | 2     | GCD of Strings, Date to Binary                            |
| Matrix         | 1     | Valid Matrix Check                                        |

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
