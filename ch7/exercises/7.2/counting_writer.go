package main

import "io"

// countingWriter wraps an io.Writer and tracks the number of bytes written.
type countingWriter struct {
	writer io.Writer // The underlying writer to delegate writes to
	count  *int64    // Pointer to the byte counter
}

// Write writes data to the underlying writer and updates the byte counter.
func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p) // Delegate the write
	*cw.count += int64(n)        // Update the byte counter
	return n, err                // Return the result of the write
}

// CountingWriter returns a Writer that wraps the original writer and counts bytes written.
// It also returns a pointer to the byte counter.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64                          // Initialize the counter
	cw := &countingWriter{writer: w, count: &count} // Create the wrapped writer
	return cw, &count                        // Return the wrapped writer and counter
}
