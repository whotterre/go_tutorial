package main

// rotate rotates a slice s left by n positions in a single pass.
func rotate(s []int, n int) {
    n = n % len(s)
    if n == 0 {
        return 
    }

    // Use a temporary copy to rearrange elements
    tmp := append(s[n:], s[:n]...)
    copy(s, tmp)
}


