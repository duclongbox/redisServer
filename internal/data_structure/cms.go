package data_structure

import (
	"github.com/spaolacci/murmur3"
	"math"
)

const Log10PoitFive = -0.3010299956639 // computed value of log10(0.5)

type CMS struct{
	width uint32
	depth uint32
	// 2D slice to hold the counts
	counter [][]uint32
}


func CreateCMS(w uint32, d uint32) *CMS {
	cms := &CMS{
		width:   w,
		depth:   d,
	}

	// Initialize the 2D slice with  d elements
	cms.counter = make([][]uint32, d)
	// loop through initialize each row with w elements
	for i := range cms.counter {
		cms.counter[i] = make([]uint32, w)
	}	
	return cms
}

//calulate dimensions of width and rows based on error rate and probability
//errRate: excepted error rate for frequency estimation
//errProb: accepteed probability that error exceeds the error rate 
func CalcCMSDim(errRate float64, errProb float64) (uint32, uint32) {
	w := uint32(math.Ceil(2.0 / errRate)) 

	d := uint32(math.Ceil(math.Log10(errProb) / Log10PoitFive)) 
	return w, d
}

func (c *CMS) calcHash(item string, seed uint32) uint32 {
	hasher := murmur3.New32WithSeed(seed)
	hasher.Write([]byte(item))
	return hasher.Sum32()
}

// Increments the count for an item by a specific value
// return the estimated count for the item after increment
func (c *CMS) IncrBy(item string, value uint32) uint32{
	var minCount uint32 = math.MaxUint32

	// loop through each row of 2D slice
	for i := uint32(0); i < c.depth; i++{
		//calculate each hash value with different seed
		hash := c.calcHash(item, i)
		// get column index in the row
		j := hash % c.width
		
		// safely add the value to preven overflow
		if math.MaxUint32 - c.counter[i][j] < value {
			c.counter[i][j] = math.MaxUint32
		} else {
			c.counter[i][j] += value
		}

		// update minCount if the current count is smaller
		if c.counter[i][j] < minCount{
			minCount = c.counter[i][j]
		}
		
	}
	return minCount
}

// returns the estimated count for an item
// retrieves the miniim count across all hash fucntion rows, -> provide the most accurate
func (c *CMS) Count(item string) uint32{
	var minCount uint32 = math.MaxUint32

	// loop through each row of 2D slice
	for i := uint32(0); i < c.depth; i++{
		//calculate each hash value with different seed
		hash := c.calcHash(item, i)
		// get column index in the row
		j := hash % c.width	

		if c.counter[i][j] < minCount{
			minCount = c.counter[i][j]
		}

	}
	return minCount
}
