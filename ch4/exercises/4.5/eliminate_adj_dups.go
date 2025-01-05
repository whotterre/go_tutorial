package main

func adjacentDupsRemover(slice []string) []string {
    if len(slice) == 0 {
        return slice
    }

    // writeIdx points to the next position to write a unique element
    writeIdx := 1

    for i := 1; i < len(slice); i++ {
        if slice[i] != slice[i-1] {
			// Move unique element to the correct position
            slice[writeIdx] = slice[i] 
            writeIdx++
        }
    }

    return slice[:writeIdx]
}