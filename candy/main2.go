package main
import "fmt"


func candy(ratings []int) int {
	n:= len(ratings)
	candies:=make([]int,n)

	for i:=0 ; i<n ; i++ {
		candies[i]=1
	}
}





func main() {
	rating:=[]int{1,0,2}
	fmt.Println(candy(rating))

}