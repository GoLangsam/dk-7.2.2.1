# dk-7.2.2.1

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangsam/dk-7.2.2.1)](https://goreportcard.com/report/github.com/GoLangsam/dk-7.2.2.1)
[![Build Status](https://travis-ci.org/GoLangsam/dk-7.2.2.1.svg?branch=master)](https://travis-ci.org/GoLangsam/dk-7.2.2.1)
[![GoDoc](https://godoc.org/github.com/GoLangsam/dk-7.2.2.1?status.svg)](https://godoc.org/github.com/GoLangsam/dk-7.2.2.1)

[dk](https://en.wikipedia.org/wiki/Donald_Knuth)-[7.2.2.1](http://www.cs.stanford.edu/~knuth/fasc5c.ps.gz) provides package [dl](https://en.wikipedia.org/wiki/Dancing_Links) - what else ;-)

---
Note: ['Don'](https://en.wikipedia.org/wiki/Donald_Knuth) intends to "save time and space by allocationg them sequential."
[Page 4 first paragraph in his pre-print](http://www.cs.stanford.edu/~knuth/fasc5c.ps.gz).

The implementation in this repo takes his approach very literally.

And *Yes*, the data is very compact and its locality may be very friendly to CPU caches.
Further, it is easy to clone the data as is needed by any concurrent approach to subproblems.

Just - due to the very many index operations involved in this sequential allocation, performance time suffers substantially.

So this implementation is useful as a counter-example and for comparsions with other implementations
such as a hybrid approach (with pointers and indices) [here](https://github.com/GoLangsam/taocp-7.2.2.1)
or a pointers-only approach [here](https://github.com/GoLangsam/tees).

---
Your suggestions, remarks, questions and/or contributions are welcome ;-)

---
## Think deep - code happy - be simple - see clear :-)

---
## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/GoLangsam/dk-7.2.2.1/badge.svg?style=beer-square)](https://beerpay.io/GoLangsam/dk-7.2.2.1)  [![Beerpay](https://beerpay.io/GoLangsam/dk-7.2.2.1/make-wish.svg?style=flat-square)](https://beerpay.io/GoLangsam/dk-7.2.2.1?focus=wish)