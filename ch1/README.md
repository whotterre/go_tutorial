# Chapter 1 - Tutorial

Gave samples of Golang code to show relevant elements of a Golang program.

## Exercises

1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.

1.2: Modify the echo program to print the index and value of each of its arguments, one per line

1.7: The function call io.Copy(dst, src) reads from src and writes to dst. Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.

1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.

1.9: Modify fetch to also print the HTTP status code, found in resp.Status.

1.10: Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much.
Do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.