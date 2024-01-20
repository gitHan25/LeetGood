
func areaOfMaxDiagonal(dimensions [][]int) int {
    maxVal:=0
    maxDiag:=0
    for _,d := range dimensions{
        length,width := d[0],d[1]
        
        total := (length*length)+(width*width)
        if total > maxDiag|| total == maxDiag && d[0]*d[1]>maxVal {
            maxDiag=total
            maxVal = d[0]*d[1]
            
            
        }
    }
    return maxVal
}