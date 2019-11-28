package hello

import (
   "rsc.io/quote"
   quoteV3 "rsc.io/quote/v3"   
)

func Proverb() string {
    return quoteV3.Concurrency()
}

func Hello() string{
	return quote.Hello()
}




