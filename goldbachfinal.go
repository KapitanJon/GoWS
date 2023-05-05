package main
//Necessary Packages To Run
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	)

//This function accepts an integer n
//It sees if the n is prime and returns true if it is and false if it isnt.
func isPrime(n int) bool{
	if n < 2{
		return false
		}
	for i:=2; i*i <= n; i++{
		if n%i == 0{
			return false
			}
		}
	return true
	}	
		
//This function gets all of the prime numbers up until the
//input n
//returns an array of those integers
func sieveOfErat(n int) []int{
	primes := make([]int, 0)
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++{
		if !sieve[i]{
			primes = append(primes,i)
			for j := i * i; j <= n; j += i {
               			sieve[j] = true
            }
        }
    }
    return primes
}
//Does the goldbach conjecture
//accepts an input returns a set of arrays in an array called results
func Goldbach(n int)[][]int{
	num := n
	primes := sieveOfErat(num)
	var results [][]int
	for _, p := range primes{
		if isPrime(num - p){
			stop := false
			//NOTE:This removes all repitions from the output
			for _, pair := range results{
				if pair[0] == p || pair[1] == p || pair[0] == n-p || pair[1] == n-p{
				stop = true
				break
				}
			}
			if ! stop{
			pair := []int{p, num -p}
			results = append(results, pair)
			}
			}
		}
		return results
	}
//Main function contains alot of boilerplate
//calls necessary to read from and
//write to a file
//Main calls Goldbach and rights the outpu 'results' to a file
//It works with a file of any length
	func main(){
	file,err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
        }
	defer file.Close()
	outputFile, err := os.Create("myResults.txt")
	        if err != nil {
            fmt.Println(err)
            return
        }
        defer outputFile.Close()
        writein := bufio.NewWriter(outputFile)
	scan := bufio.NewScanner(file)
	for scan.Scan(){
		num,err := strconv.Atoi(scan.Text())
	if err != nil {
		fmt.Println(err)
		return
        }
	
	pairs := Goldbach(num)
	foundpairs := len(pairs)
	fmt.Fprintf(writein,"We found %d pairs for %d.\n", foundpairs,num)
	for _, pair := range pairs{
		fmt.Fprintf(writein,"%d = %d + %d\n",num,pair[0],pair[1])
	}
	fmt.Fprintln(writein)
	}
}
